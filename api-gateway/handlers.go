package main

import (
	"context"
	"encoding/json"
	"microsvc/api-gateway/proto"
	"microsvc/common/utils"
	"net/http"
)

type handler struct {
	logger *utils.CustomLogger
}

func NewHandler(l *utils.CustomLogger) *handler {
	return &handler{
		logger: l,
	}
}

type UserReg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) ProxyRegReq(w http.ResponseWriter, r *http.Request) {
	u := UserReg{}

	_ = json.NewDecoder(r.Body).Decode(&u)

	ctx := context.Background()
	req := proto.GatewayRegisterRequest{
		Username: u.Username,
		Password: u.Password,
	}

	h.logger.Info("Received reg data, redirect to gRPC server")
	resp, err := GatewayServer.Register(ctx, &req)
	if err != nil {
		h.logger.Error("GatewayServer.Register error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)

	h.logger.Info("Received gRPC server response, send to client")
	json.NewEncoder(w).Encode(&resp)
}

func (h *handler) ProxyAuthReq(w http.ResponseWriter, r *http.Request) {

}
