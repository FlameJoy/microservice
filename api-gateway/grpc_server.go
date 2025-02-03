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
	pbProduct "microsvc/product-service/proto"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GatewayServer *grpcServer
)

type grpcServer struct {
	pbGateway.GatewayServiceServer
	authClient    pbAuth.AuthServiceClient
	productClient pbProduct.ProductServiceClient
	orderClient   pbOrder.OrderServiceClient
	logger        *utils.CustomLogger
}

func NewGRPCServer(authSvcAddr, orderSvcAddr, productSvcAddr string, logger *utils.CustomLogger) (*grpcServer, error) {
	// Используем NewClient для подключения к Auth-сервису
	authConn, err := grpc.NewClient(authSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	authClient := pbAuth.NewAuthServiceClient(authConn)

	productConn, err := grpc.NewClient(productSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	productClient := pbProduct.NewProductServiceClient(productConn)

	orderConn, err := grpc.NewClient(orderSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	orderClient := pbOrder.NewOrderServiceClient(orderConn)

	return &grpcServer{
		authClient:    authClient,
		productClient: productClient,
		orderClient:   orderClient,
		logger:        logger,
	}, nil
}

// Реализация метода Login через Auth-сервис
func (s *grpcServer) Login(ctx context.Context, req *pbGateway.GatewayLoginReq) (*pbGateway.GatewayLoginResp, error) {
	// Проксируем запрос в Auth-сервис
	authReq := &pbAuth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	authResp, err := s.authClient.Login(ctx, authReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayLoginResp{
		Token:   authResp.Token,
		Message: authResp.Message,
	}, nil
}

// Реализация метода Register через Auth-сервис
func (s *grpcServer) Register(ctx context.Context, req *pbGateway.GatewayRegisterReq) (*pbGateway.GatewayRegisterResp, error) {
	regReq := &pbAuth.RegRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	authResp, err := s.authClient.Register(ctx, regReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayRegisterResp{
		Message: authResp.Message,
	}, nil
}

func (s *grpcServer) CreateProduct(ctx context.Context, req *pbGateway.GatewayCreateProductReq) (*pbGateway.GatewayCreateProductResp, error) {
	createReq := &pbProduct.CreateReq{
		SKU:      req.SKU,
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
		UOM:      req.UOM,
		Brand:    req.Brand,
		Stock:    req.Stock,
	}

	s.logger.Info("api gateway: starts gRPC server CreateProduct func")

	resp, err := s.productClient.Create(ctx, createReq)
	if err != nil {
		return nil, err
	}

	fmt.Println("ERROR")

	return &pbGateway.GatewayCreateProductResp{
		Id:      resp.Id,
		Message: resp.Message,
	}, err
}

func (s *grpcServer) UpdateProduct(ctx context.Context, req *pbGateway.GatewayUpdateProductReq) (*pbGateway.GatewayUpdateProductResp, error) {
	// Проксируем обновления в ProductService
	updateReq := &pbProduct.UpdateReq{
		SqlQuery: req.SqlQuery,
		Args:     req.Args,
	}

	updateResp, err := s.productClient.Update(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayUpdateProductResp{
		Success: updateResp.Success,
		Message: updateResp.Message,
	}, nil
}

func (s *grpcServer) DeleteProduct(ctx context.Context, req *pbGateway.GatewayDeleteProductReq) (*pbGateway.GatewayDeleteProductResp, error) {
	deleteReq := &pbProduct.DeleteReq{
		Id: req.Id,
	}

	resp, err := s.productClient.Delete(ctx, deleteReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayDeleteProductResp{
		Message: resp.Message,
	}, nil
}

func (s *grpcServer) CreateOrder(ctx context.Context, req *pbGateway.GatewayOrderCreateReq) (*pbGateway.GatewayOrderCreateResp, error) {
	orderReq := &pbOrder.CreateReq{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}

	s.logger.Info("api gateway: starts gRPC server CreateOrder func")

	orderResp, err := s.orderClient.Create(ctx, orderReq)
	if err != nil {
		return nil, err
	}

	return &pbGateway.GatewayOrderCreateResp{
		Id:       orderResp.Id,
		TotalSum: orderResp.TotalSum,
		Status:   orderResp.Status,
		Message:  orderResp.Message,
	}, nil
}

func StartGRPCServer(address, authSvcAddr string, orderSvcAddr string, productSvcAddr string, done chan os.Signal, logger *utils.CustomLogger) {
	// Создаём Gateway gRPC-сервер
	var err error
	GatewayServer, err = NewGRPCServer(authSvcAddr, orderSvcAddr, productSvcAddr, logger)
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
