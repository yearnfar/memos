package server

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/yearnfar/memos/internal/api/v1"
	"github.com/yearnfar/memos/internal/config"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func registerGRPC(s *Server) {
	v1pb.RegisterAuthServiceServer(s.grpcServer, &v1.AuthService{})
	v1pb.RegisterWorkspaceServiceServer(s.grpcServer, &v1.WorkspaceService{})
	v1pb.RegisterWorkspaceSettingServiceServer(s.grpcServer, &v1.WorkspaceSettingService{})
	v1pb.RegisterUserServiceServer(s.grpcServer, &v1.UserService{})
	v1pb.RegisterInboxServiceServer(s.grpcServer, &v1.InboxService{})
	v1pb.RegisterMemoServiceServer(s.grpcServer, &v1.MemoService{})
	v1pb.RegisterResourceServiceServer(s.grpcServer, &v1.ResourceService{})
	reflection.Register(s.grpcServer)
}

func registerGateway(ctx context.Context, s *Server) error {
	cfg := config.GetApp().Server
	conn, _ := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024)),
	)

	gwMux := runtime.NewServeMux()
	if err := v1pb.RegisterAuthServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}
	if err := v1pb.RegisterWorkspaceServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}
	if err := v1pb.RegisterWorkspaceSettingServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}
	if err := v1pb.RegisterUserServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}
	if err := v1pb.RegisterInboxServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}
	if err := v1pb.RegisterMemoServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}
	if err := v1pb.RegisterResourceServiceHandler(ctx, gwMux, conn); err != nil {
		return err
	}

	gwGroup := s.echoServer.Group("")
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
	s.echoServer.Any("/memos.api.v1.*", echo.WrapHandler(wrappedGrpc))
	return nil
}
