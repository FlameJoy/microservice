package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"microsvc/api-gateway/proto"
	"microsvc/common/utils"
	"microsvc/data"
	"net/http"
	"strings"
	"time"
	"unicode"
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

	req := proto.GatewayRegisterReq{
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
	var u data.User
	if err := u.FromJSON(r.Body); err != nil {
		utils.HttpRespErrRFC9457("ProxyAuthReq", "FromJSON error", err, http.StatusBadRequest, w, r, h.logger)
		return
	}
	defer r.Body.Close()

	req := proto.GatewayLoginReq{
		Username: u.Username,
		Password: u.Pswd,
	}

	h.logger.Info("Received login data, redirect to auth-svc")

	resp, err := GatewayServer.Login(r.Context(), &req)
	if err != nil {
		utils.HttpRespErrRFC9457("ProxyAuthReq", "GatewayServer.Login error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}

	if resp.Token == "" {
		utils.HttpRespErrRFC9457("ProxyAuthReq", "Missing token in response", nil, http.StatusUnauthorized, w, r, h.logger)
		return
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "token",
	// 	Value:    resp.Token,
	// 	Path:     "/",
	// 	HttpOnly: true,  // XSS-protection
	// 	Secure:   false, // true, for HTTPS
	// 	MaxAge:   3600,
	// })

	h.logger.Info("Received auth-svc response, send to client")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		utils.HttpRespErrRFC9457("ProxyRegReq", "Encode error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}
}

type KeyProduct struct{}

func (h *handler) ProductValidate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := data.Product{}

		if err := p.FromJSON(r.Body); err != nil {
			utils.HttpRespErrRFC9457("ProductValidate", "FromJSON error", err, http.StatusBadRequest, w, r, h.logger)
			return
		}

		if err := p.Validate(); err != nil {
			utils.HttpRespErrRFC9457("ProductValidate", "Validation error", err, http.StatusBadRequest, w, r, h.logger)
			return
		}

		ctx := context.WithValue(context.Background(), KeyProduct{}, p)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func (h *handler) ProxyCreateProduct(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value(KeyProduct{})
	p, ok := value.(data.Product)
	if !ok {
		utils.HttpRespErrRFC9457("ProxyCreateProduct", "Interface conversion error", fmt.Errorf("%v is not data.Product", value), http.StatusInternalServerError, w, r, h.logger)
		return
	}

	req := proto.GatewayCreateProductReq{
		SKU:      utils.GenRandNums(10),
		Name:     p.Name,
		Price:    int64(p.Price),
		Category: p.Category,
		UOM:      p.UOM,
		Brand:    p.Brand,
		Stock:    int64(p.Stock),
	}

	resp, err := GatewayServer.CreateProduct(r.Context(), &req)
	if err != nil {
		utils.HttpRespErrRFC9457("ProxyCreateProduct", "GatewayServer.CreateProduct error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		utils.HttpRespErrRFC9457("ProxyRegReq", "Encode error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}
}

func (h *handler) ProxyUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var p data.Product
	if err := p.FromJSON(r.Body); err != nil {
		utils.HttpRespErrRFC9457("ProxyUpdateProduct", "FromJSON error", err, http.StatusBadRequest, w, r, h.logger)
		return
	}

	var setClauses []string
	var args []interface{}
	argIdx := 1

	if p.ID == 0 {
		utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("no product ID"), http.StatusBadRequest, w, r, h.logger)
		return
	}

	if p.Name != "" {
		if len(p.Name) > 50 {
			utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("product name too long"), http.StatusBadRequest, w, r, h.logger)
			return
		}

		for _, char := range p.Name {
			if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
				utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("product name must contain only letters and numbers"), http.StatusBadRequest, w, r, h.logger)
				return
			}
		}

		setClauses = append(setClauses, fmt.Sprintf("name=$%d", argIdx))
		args = append(args, p.Name)
		argIdx++
	}

	if p.Price > 0 {
		setClauses = append(setClauses, fmt.Sprintf("price=$%d", argIdx))
		args = append(args, p.Price)
		argIdx++
	}

	if p.Category != "" {
		if len(p.Category) > 50 {
			utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("category too long"), http.StatusBadRequest, w, r, h.logger)
			return
		}

		for _, char := range p.Name {
			if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
				utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("category must contain only letters and numbers"), http.StatusBadRequest, w, r, h.logger)
				return
			}
		}

		setClauses = append(setClauses, fmt.Sprintf("category=$%d", argIdx))
		args = append(args, p.Category)
		argIdx++
	}

	if p.UOM != "" {
		setClauses = append(setClauses, fmt.Sprintf("uom=$%d", argIdx))
		args = append(args, p.UOM)
		argIdx++
	}

	if p.Brand != "" {
		if len(p.Brand) > 50 {
			utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("brand too long"), http.StatusBadRequest, w, r, h.logger)
			return
		}

		for _, char := range p.Name {
			if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
				utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("brand must contain only letters and numbers"), http.StatusBadRequest, w, r, h.logger)
				return
			}
		}

		setClauses = append(setClauses, fmt.Sprintf("brand=$%d", argIdx))
		args = append(args, p.Brand)
		argIdx++
	}

	// Добавляем updated_at
	setClauses = append(setClauses, fmt.Sprintf("updated_at=$%d", argIdx))
	args = append(args, time.Now().Local().Format(time.RFC3339))
	argIdx++

	// Если нет обновляемых полей
	if len(setClauses) == 0 {
		utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Validation", errors.New("no fields to update"), http.StatusBadRequest, w, r, h.logger)
		return
	}

	// Собираем SQL-запрос
	update := fmt.Sprintf("UPDATE products SET %s WHERE id=$%d", strings.Join(setClauses, ", "), argIdx)
	args = append(args, p.ID) // Добавляем ID в конец

	req := proto.GatewayUpdateProductReq{
		SqlQuery: update,
		Args:     utils.ArgsToStringSlice(args), // Преобразуем аргументы в строки
	}

	resp, err := GatewayServer.UpdateProduct(context.Background(), &req)
	if err != nil {
		utils.HttpRespErrRFC9457("ProxyUpdateProduct", "GatewayServer.CreateProduct error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}

	if !resp.Success {
		utils.HttpRespErrRFC9457("ProxyUpdateProduct", "GatewayServer.CreateProduct error", errors.New(resp.Message), http.StatusInternalServerError, w, r, h.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		utils.HttpRespErrRFC9457("ProxyUpdateProduct", "Encode error", err, http.StatusInternalServerError, w, r, h.logger)
		return
	}
}

// type KeyOrder struct{}

// func (h *handler) OrderValidate(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		o := data.Order{}

// 		o.FromJSON(r.Body)

// 		validate := validator.New(validator.WithRequiredStructEnabled())

// 		if err := validate.Struct(o); err != nil {
// 			utils.HttpRespErrRFC9457("OrderValidate", "Validation error", err, http.StatusBadRequest, w, r, h.logger)
// 			return
// 		}

// 		ctx := context.WithValue(context.Background(), KeyOrder{}, o)
// 		req := r.WithContext(ctx)
// 		next.ServeHTTP(w, req)
// 	})
// }

// func (h *handler) ProxyOrderCreate(w http.ResponseWriter, r *http.Request) {
// 	value := r.Context().Value(KeyOrder{})
// 	o, ok := value.(data.Order)
// 	if !ok {
// 		utils.HttpRespErrRFC9457("ProxyOrderCreate", "Interface conversion error", fmt.Errorf("%v is nil, not data.Order", o), http.StatusInternalServerError, w, r, h.logger)
// 		return
// 	}

// 	req := proto.GatewayOrderCreateReq{
// 		ItemID:   int32(o.ItemID),
// 		Name:     o.Name,
// 		Quantity: int32(o.Quantity),
// 		Price:    int32(o.Price),
// 	}

// 	h.logger.Info("Received order data, redirect to gRPC server")

// 	resp, err := GatewayServer.CreateOrder(r.Context(), &req)
// 	if err != nil {
// 		utils.HttpRespErrRFC9457("ProxyOrderCreate", "GatewayServer.CreateOrder error", err, http.StatusInternalServerError, w, r, h.logger)
// 		return
// 	}

// 	h.logger.Info("Received gRPC server response, send to client")

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	if err = json.NewEncoder(w).Encode(&resp); err != nil {
// 		utils.HttpRespErrRFC9457("ProxyOrderCreate", "Encode error", err, http.StatusInternalServerError, w, r, h.logger)
// 		return
// 	}
// }
