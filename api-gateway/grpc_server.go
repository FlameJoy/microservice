package main

import (
	"context"
	"log"
	"os"

	// pb "microsvc/api-gateway/proto"

	pbGateway "microsvc/api-gateway/proto"
	pbAuth "microsvc/auth-service/proto"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GatewayServer *grpcServer
)

type grpcServer struct {
	pbGateway.GatewayServiceServer
	authClient pbAuth.AuthServiceClient
}

func NewGRPCServer(authSvcAddr string) (*grpcServer, error) {
	// Используем NewClient для подключения к Auth-сервису
	conn, err := grpc.NewClient(authSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	authClient := pbAuth.NewAuthServiceClient(conn)

	return &grpcServer{
		authClient: authClient,
	}, nil
}

// Реализация метода Login через Auth-сервис
func (s *grpcServer) Login(ctx context.Context, req *pbGateway.GatewayLoginRequest) (*pbGateway.GatewayLoginResponse, error) {
	// Проксируем запрос в Auth-сервис
	authReq := &pbAuth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	authResp, err := s.authClient.Login(ctx, authReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayLoginResponse{
		Token: authResp.Token,
	}, nil
}

// Реализация метода Register через Auth-сервис
func (s *grpcServer) Register(ctx context.Context, req *pbGateway.GatewayRegisterRequest) (*pbGateway.GatewayRegisterResponse, error) {
	authReq := &pbAuth.RegRequest{
		Username: req.Username,
		Password: req.Password,
	}

	logger.Info("api gateway: starts gRPC server Register func")

	authResp, err := s.authClient.Register(ctx, authReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayRegisterResponse{
		Message: authResp.Message,
	}, nil
}

func StartGRPCServer(address, authSvcAddr string, done chan os.Signal) {
	// Создаём Gateway gRPC-сервер
	var err error
	GatewayServer, err = NewGRPCServer(authSvcAddr)
	if err != nil {
		log.Fatalf("Failed to create gRPC server: %v", err)
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()

	// Регистрируем GatewayService
	pbGateway.RegisterGatewayServiceServer(grpcServer, GatewayServer)

	go func() {
		logger.Info("gRPC Gateway server is running on %s", address)
		// log.Printf("gRPC Gateway server is running on %s", address)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	<-done
	logger.Info("grpcServer is shutting down...")

	grpcServer.GracefulStop()
	logger.Info("grpcServer exited properly")
}
