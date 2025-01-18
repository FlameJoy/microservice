package main

import (
	"context"
	"fmt"
	"math/rand"
	"microsvc/common/utils"
	"microsvc/order-service/proto"
)

type OrderServer struct {
	proto.UnimplementedOrderServiceServer
	logger *utils.CustomLogger
}

func (s *OrderServer) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateResp, error) {
	s.logger.Info("Validate order data")
	if req.Name == "" || req.Quantity <= 0 || req.Price <= 0 || req.ItemID == 0 {
		return nil, fmt.Errorf("invalid credentials")
	}

	s.logger.Info("Prepare response")
	resp := new(proto.CreateResp)
	resp.ID = int32(rand.Intn(100))
	resp.Name = "Random name"
	resp.TotalSum = req.Price * req.Quantity

	return resp, nil
}

func (s *OrderServer) Update(ctx context.Context, req *proto.UpdateReq) (*proto.UpdateResp, error) {
	if req.Name == "" || req.Quantity <= 0 || req.Price <= 0 || req.ID == 0 {
		return nil, fmt.Errorf("invalid credentials")
	}

	update := new(proto.UpdateResp)
	update.ID = req.ID
	update.Name = req.Name
	update.TotalSum = req.Price * req.Quantity

	return update, nil
}

func (s *OrderServer) Delete(ctx context.Context, req *proto.DeleteReq) (*proto.DeleteResp, error) {
	delete := new(proto.DeleteResp)
	if req.ID == 0 {
		delete.Deleted = false
		delete.Message = "id mismatch"
		return delete, fmt.Errorf("invalid credentials")
	}

	delete.Deleted = true
	delete.Message = "successfuly deleted"
	return delete, nil
}
