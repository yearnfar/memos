package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/yearnfar/memos/internal/api/v1"
	"github.com/yearnfar/memos/internal/module/user"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() {
	user.Register(userSvc.Default())
}

func Register(grpcServer *grpc.Server) {
	v1pb.RegisterAuthServiceServer(grpcServer, &v1.AuthService{})
	reflection.Register(grpcServer)
}

func RegisterGateway(ctx context.Context, conn *grpc.ClientConn, echoServer *echo.Echo) error {
	gwMux := runtime.NewServeMux()
	if err := v1pb.RegisterWorkspaceServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}

	gwGroup := echoServer.Group("")
	gwGroup.Use(middleware.CORS())
	handler := echo.WrapHandler(gwMux)

	gwGroup.Any("/api/v1/*", handler)
	gwGroup.Any("/file/*", handler)

	// GRPC web proxy.
	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(_ string) bool {
			return true
		}),
	}
	wrappedGrpc := grpcweb.WrapServer(s.grpcServer, options...)
	echoServer.Any("/memos.api.v1.*", echo.WrapHandler(wrappedGrpc))
	return nil
}
