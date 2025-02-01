package main

import (
	"context"
	"microsvc/common/utils"
	"microsvc/product-service/proto"
	"microsvc/storage"
	"time"
)

type ProductServer struct {
	proto.UnimplementedProductServiceServer
	logger  *utils.CustomLogger
	storage storage.Storage
}

func (s *ProductServer) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateResp, error) {
	s.logger.Info("product svc: starts gRPC server Create func")

	query := "INSERT INTO products (sku, name, price, category, uom, brand, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"

	id, err := s.storage.ExecuteQuery(query,
		req.SKU, req.Name, req.Price, req.Category, req.UOM, req.Brand, req.Stock, time.Now().Local().Format(time.RFC3339), time.Now().Local().Format(time.RFC3339),
	)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResp{Id: id, Message: "Product succesfully created"}, nil
}

func (s *ProductServer) Update(ctx context.Context, req *proto.UpdateReq) (*proto.UpdateResp, error) {
	s.logger.Info("product svc: starts gRPC server Update func")

	// Преобразуем аргументы обратно
	args := make([]interface{}, len(req.Args))
	for i, arg := range req.Args {
		args[i] = utils.ParseArg(arg)
	}

	// Открываем транзакцию
	tx, err := s.storage.DB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Блокируем строку, чтобы никто другой не мог её обновить, пока текущая транзакция не завершится
	_, err = tx.Exec("SELECT id FROM products WHERE id=$1 FOR UPDATE", args[len(args)-1])
	if err != nil {
		return nil, err
	}

	// Выполняем UPDATE
	_, err = tx.Exec(req.SqlQuery, args...)
	if err != nil {
		return nil, err
	}

	// Фиксируем транзакцию
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &proto.UpdateResp{
		Success: true,
		Message: "Product updated successfully",
	}, nil
}
