package main

import (
	"context"
	"encoding/json"
	"fmt"
	"microsvc/api-gateway/data"
	"microsvc/api-gateway/proto"
	"microsvc/common/utils"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type handler struct {
	logger *utils.CustomLogger
}

func NewHandler(l *utils.CustomLogger) *handler {
	return &handler{
		logger: l,
	}
}

type KeyUser struct{}

func (h *handler) UserValidate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := data.User{}

		u.FromJSON(r.Body)

		// Username validation
		if err := u.ValidateUsername(); err != nil {
			utils.HttpRespErrRFC9457("UserValidate", "Validation error", err, http.StatusBadRequest, w, r, h.logger)
			return
		}

		// Email validation
		if err := u.ValidateEmail(); err != nil {
			utils.HttpRespErrRFC9457("UserValidate", "Validation error", err, http.StatusBadRequest, w, r, h.logger)
			return
		}

		// Password validation
		if err := u.ValidatePswd(); err != nil {
			utils.HttpRespErrRFC9457("UserValidate", "Validation error", err, http.StatusBadRequest, w, r, h.logger)
			return
		}

		ctx := context.WithValue(context.Background(), KeyUser{}, u)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func (h *handler) ProxyRegReq(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value(KeyUser{})
	u, ok := value.(data.User)
	if !ok {
		utils.HttpRespErrRFC9457("ProxyRegReq", "Interface conversion error", fmt.Errorf("%v is nil, not data.User", u), http.StatusInternalServerError, w, r, h.logger)
		return
	}

	req := proto.GatewayRegisterRequest{
		Username: u.Username,
		Password: u.Pswd,
		Email:    u.Email,
	}

	h.logger.Info("Received registration data, redirect to gRPC server")

	resp, err := GatewayServer.Register(r.Context(), &req)
	if err != nil {
		utils.HttpRespErrRFC9457("ProxyRegReq", "GatewayServer.Register error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}

	h.logger.Info("Received gRPC server response, send to client")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		utils.HttpRespErrRFC9457("ProxyRegReq", "Encode error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}
}

func (h *handler) ProxyAuthReq(w http.ResponseWriter, r *http.Request) {

}

type KeyOrder struct{}

func (h *handler) OrderValidate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := data.Order{}

		o.FromJSON(r.Body)

		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(o); err != nil {
			utils.HttpRespErrRFC9457("OrderValidate", "Validation error", err, http.StatusBadRequest, w, r, h.logger)
			return
		}

		ctx := context.WithValue(context.Background(), KeyOrder{}, o)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func (h *handler) ProxyOrderCreate(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value(KeyOrder{})
	o, ok := value.(data.Order)
	if !ok {
		utils.HttpRespErrRFC9457("ProxyOrderCreate", "Interface conversion error", fmt.Errorf("%v is nil, not data.Order", o), http.StatusInternalServerError, w, r, h.logger)
		return
	}

	req := proto.GatewayOrderCreateReq{
		ItemID:   int32(o.ItemID),
		Name:     o.Name,
		Quantity: int32(o.Quantity),
		Price:    int32(o.Price),
	}

	h.logger.Info("Received order data, redirect to gRPC server")

	resp, err := GatewayServer.CreateOrder(r.Context(), &req)
	if err != nil {
		utils.HttpRespErrRFC9457("ProxyOrderCreate", "GatewayServer.CreateOrder error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}

	h.logger.Info("Received gRPC server response, send to client")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		utils.HttpRespErrRFC9457("ProxyOrderCreate", "Encode error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}
}
