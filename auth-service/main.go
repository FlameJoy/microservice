package main

import (
	"context"
	"fmt"
	"log"
	"microsvc/auth-service/proto" // Путь к сгенерированным файлам gRPC
	"microsvc/common/utils"
	"net"
	"os"

	"google.golang.org/grpc"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
}

// Метод Login
func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// Здесь ваша логика для проверки логина пользователя
	if req.Username == "admin" && req.Password == "password" {
		return &proto.LoginResponse{Token: "some_token"}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}

// Метод Register
func (s *AuthServer) Register(ctx context.Context, req *proto.RegRequest) (*proto.RegResponse, error) {

	fmt.Println("auth svc: starts gRPC server Register func")
	fmt.Println(req.Username)
	fmt.Println(req.Password)

	// Здесь ваша логика для регистрации нового пользователя
	if req.Username != "" && req.Password != "" {
		// Пример: добавление пользователя в базу данных
		log.Printf("User %s registered successfully", req.Username)
		return &proto.RegResponse{Message: "User registered successfully"}, nil
	}
	return nil, fmt.Errorf("invalid input")
}

// Запуск gRPC-сервера
func StartGRPCServer(address string) error {
	// Инициализация gRPC сервера
	server := grpc.NewServer()

	// Регистрация AuthService
	proto.RegisterAuthServiceServer(server, &AuthServer{})

	// Прослушивание на порту
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Printf("Auth service is listening on %s", address)

	// Запуск gRPC-сервера
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %v", err)
	}

	return nil
}

func main() {
	// Логирование и старт сервера
	logger := utils.NewLogger(utils.INFO, log.New(os.Stdout, "auth-service ", log.LstdFlags), false)

	// Запуск gRPC сервера на порту 50051
	if err := StartGRPCServer(":50052"); err != nil {
		logger.Fatal("Failed to start gRPC server: %v", err)
	}
}
