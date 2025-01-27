package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// pb "microsvc/api-gateway/proto"

	pbGateway "microsvc/api-gateway/proto"
	pbAuth "microsvc/auth-service/proto"
	"microsvc/common/utils"
	pbOrder "microsvc/order-service/proto"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GatewayServer *grpcServer
)

type grpcServer struct {
	pbGateway.GatewayServiceServer
	authClient  pbAuth.AuthServiceClient
	orderClient pbOrder.OrderServiceClient
	logger      *utils.CustomLogger
}

func NewGRPCServer(authSvcAddr string, orderSvcAddr string, logger *utils.CustomLogger) (*grpcServer, error) {
	// Используем NewClient для подключения к Auth-сервису
	authConn, err := grpc.NewClient(authSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	authClient := pbAuth.NewAuthServiceClient(authConn)

	orderConn, err := grpc.NewClient(orderSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	orderClient := pbOrder.NewOrderServiceClient(orderConn)

	return &grpcServer{
		authClient:  authClient,
		orderClient: orderClient,
		logger:      logger,
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
		Email:    req.Email,
	}

	s.logger.Info("api gateway: starts gRPC server Register func")

	authResp, err := s.authClient.Register(ctx, authReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayRegisterResponse{
		Message: authResp.Message,
	}, nil
}

func (s *grpcServer) CreateOrder(ctx context.Context, req *pbGateway.GatewayOrderCreateReq) (*pbGateway.GatewayOrderCreateResp, error) {
	orderReq := &pbOrder.CreateReq{
		ItemID:   req.ItemID,
		Name:     req.Name,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	s.logger.Info("api gateway: starts gRPC server CreateOrder func")

	fmt.Println(orderReq)

	orderResp, err := s.orderClient.Create(ctx, orderReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayOrderCreateResp{
		ID:       orderResp.ID,
		Name:     orderResp.Name,
		TotalSum: orderResp.TotalSum,
	}, nil
}

func StartGRPCServer(address, authSvcAddr string, orderSvcAddr string, done chan os.Signal, logger *utils.CustomLogger) {
	// Создаём Gateway gRPC-сервер
	var err error
	GatewayServer, err = NewGRPCServer(authSvcAddr, orderSvcAddr, logger)
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
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	<-done
	logger.Info("grpcServer is shutting down...")

	grpcServer.GracefulStop()
	logger.Info("grpcServer exited properly")
}
