package main

import (
	"microsvc/common/utils"
	"microsvc/order-service/proto"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func StartGRPCServer(address string, logger *utils.CustomLogger) {
	server := grpc.NewServer()

	proto.RegisterOrderServiceServer(server, &OrderServer{logger: logger})

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		logger.Info("gRPC order svc server is running on %s", address)
		if err := server.Serve(listener); err != nil {
			logger.Fatal("Failed to serve gRPC server: %v", err)
		}
	}()

	<-done

	logger.Info("grpcServer is shutting down...")
	server.GracefulStop()
	logger.Info("grpcServer exited properly")
}
