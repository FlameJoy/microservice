package main

import (
	"context"
	"errors"
	"fmt"
	"microsvc/common/utils"
	"microsvc/order-service/proto"
	"microsvc/storage"
	"time"
)

type OrderServer struct {
	proto.UnimplementedOrderServiceServer
	logger  *utils.CustomLogger
	storage storage.Storage
}

func (s *OrderServer) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateResp, error) {
	s.logger.Info("order svc: starts gRPC server Create func")

	tx, err := s.storage.DB().Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	fmt.Println("ERROR")

	var price, stock int64
	selectQuery := "SELECT price, stock FROM products WHERE id=$1 FOR UPDATE"
	err = tx.QueryRow(selectQuery, req.ProductId).Scan(&price, &stock)
	if err != nil {
		return nil, err
	}

	if stock < req.Quantity {
		return nil, errors.New("not enough stock")
	}

	updateQuery := "UPDATE products SET stock = stock - $1 WHERE id = $2"
	_, err = tx.Exec(updateQuery, req.Quantity, req.ProductId)
	if err != nil {
		return nil, err
	}

	insertQuery := `
		INSERT INTO orders (user_id, product_id, quantity, price, total_sum, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var orderId int64
	err = tx.QueryRow(insertQuery, req.UserId, req.ProductId, req.Quantity, price, req.Quantity*price, "pending", time.Now(), time.Now()).Scan(&orderId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &proto.CreateResp{
		Id: orderId,
	}, nil
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
