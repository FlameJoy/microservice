package main

import (
	"log"
	"microsvc/auth-service/proto"
	"microsvc/common/utils"
	"net"

	"google.golang.org/grpc"
)

// Запуск gRPC-сервера
func StartGRPCServer(address string, logger *utils.CustomLogger) error {
	// Инициализация gRPC сервера
	server := grpc.NewServer()

	// Регистрация AuthService
	proto.RegisterAuthServiceServer(server, &AuthServer{logger: logger})

	// Прослушивание на порту
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("failed to listen: %v", err)
		return err
	}

	log.Printf("Auth service is listening on %s", address)

	// Запуск gRPC-сервера
	if err := server.Serve(listener); err != nil {
		logger.Error("failed to serve gRPC server: %v", err)
		return err
	}

	return nil
}
