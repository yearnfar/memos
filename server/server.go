package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	"github.com/yearnfar/memos/internal/config"
	"github.com/yearnfar/memos/server/interceptor"
)

type Server struct {
	echoServer *echo.Echo
	grpcServer *grpc.Server
}

func NewService(ctx context.Context) *Server {
	echoServer := echo.New()
	echoServer.Debug = true
	echoServer.HideBanner = true
	echoServer.HidePort = true
	echoServer.Use(middleware.Recover())

	grpcServer := grpc.NewServer(
		// Override the maximum receiving message size to 100M for uploading large resources.
		grpc.MaxRecvMsgSize(100*1024*1024),
		grpc.ChainUnaryInterceptor(
			interceptor.NewLoggerInterceptor().LoggerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			// interceptor.NewGRPCAuthInterceptor(nil, cfg.Secret).AuthenticationInterceptor,
		))

	s := &Server{
		echoServer: echoServer,
		grpcServer: grpcServer,
	}

	registerGRPC(s)
	registerGateway(ctx, s)
	return s
}

func (s *Server) Start(ctx context.Context) error {
	cfg := config.GetApp().Server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: grpcHandlerFunc(s.grpcServer, s.echoServer),
	}
	log.Printf("host: %s, port: %d", cfg.Host, cfg.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	return nil
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func (s *Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Shutdown echo server.
	if err := s.echoServer.Shutdown(ctx); err != nil {
		fmt.Printf("failed to shutdown server, error: %v\n", err)
	}

	fmt.Printf("memos stopped properly\n")
}
