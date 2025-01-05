package main

import (
	"context"
	"fmt"
	"microsvc/auth-service/proto"
	"microsvc/common/utils"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	logger *utils.CustomLogger
}

func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	if req.Username == "Ilya" && req.Password == "qwerty123" {
		return &proto.LoginResponse{Token: "some_token"}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegRequest) (*proto.RegResponse, error) {

	s.logger.Info("auth svc: starts gRPC server Register func")

	// Здесь ваша логика для регистрации нового пользователя
	if req.Username != "" && req.Password != "" {
		s.logger.Info("User %s registered successfully", req.Username)
		return &proto.RegResponse{Message: "User registered successfully"}, nil
	}
	return nil, fmt.Errorf("invalid input")
}
