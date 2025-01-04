package main

import (
	"encoding/json"
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

func (h *handler) UserValidate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := data.User{}

		u.FromJSON(r.Body)

		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(u); err != nil {
			utils.HttpRespErrRFC9457("UserValidate", "Validation error", err, http.StatusBadRequest, w, r, h.logger)
		}

		r.Context().Value()
	})
}

func (h *handler) ProxyRegReq(w http.ResponseWriter, r *http.Request) {
	u := data.User{}

	err := u.FromJSON(r.Body)
	if err != nil {
		utils.HttpRespErrRFC9457("ProxyRegReq", "Decode error", err, http.StatusBadRequest, w, r, h.logger)
	}

	ctx := r.Context()
	req := proto.GatewayRegisterRequest{
		Username: u.Name,
		Password: u.Pswd,
	}

	h.logger.Info("Received registration data, redirect to gRPC server")

	resp, err := GatewayServer.Register(ctx, &req)
	if err != nil {
		utils.HttpRespErrRFC9457("ProxyRegReq", "GatewayServer.Register error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}

	h.logger.Info("Received gRPC server response, send to client")

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		utils.HttpRespErrRFC9457("ProxyRegReq", "Encode error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}
}

func (h *handler) ProxyAuthReq(w http.ResponseWriter, r *http.Request) {

}
