package main

import (
	"microsvc/auth-service/proto"
	"microsvc/common/utils"
	"microsvc/storage"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// Запуск gRPC-сервера
func StartGRPCServer(address string, logger *utils.CustomLogger, storage storage.Storage) {
	// Инициализация gRPC сервера
	server := grpc.NewServer()

	// Регистрация AuthService
	proto.RegisterAuthServiceServer(server, &AuthServer{logger: logger, storage: storage})

	// Прослушивание на порту
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		logger.Info("gRPC auth svc server is running on %s", address)
		if err := server.Serve(listener); err != nil {
			logger.Fatal("Failed to serve gRPC server: %v", err)
		}
	}()

	<-done

	logger.Info("Closing storage conn...")
	if err = storage.Close(); err != nil {
		logger.Error("Can't properly close storage conn: %v", err)
	}
	logger.Info("Conn closed")

	logger.Info("grpcServer is shutting down...")
	server.GracefulStop()
	logger.Info("grpcServer exited properly")
}
