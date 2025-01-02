package handlers

import (
	"context"
	"microAPI/data"
	logging "microAPI/logger"
	"microAPI/router"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	logger *logging.CustomLogger
}

func NewProductHandler(l *logging.CustomLogger) *ProductHandler {
	return &ProductHandler{
		logger: l,
	}
}

func (h *ProductHandler) RegisterRoutes(rg *router.RouteGroup) {
	products := rg.Group("/products")
	products.GET("", h.GetProducts)
	products.GET("/mw", h.ProductValidateMW(h.GetProducts))
}

type KeyProduct struct{}

func (h *ProductHandler) ProductValidateMW(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := data.Product{}

		err := data.FromJSON(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), KeyProduct{}, data)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	path := r.URL.Path

	parts := strings.Split(path, "/")

	idStr := parts[len(parts)-1]

	_, err := strconv.Atoi(idStr)

	// Получение query-параметров
	query := r.URL.Query()
	category := query.Get("category") // Например: ?category=electronics
	price := query.Get("price")       // Например: ?price=1000

	h.logger.Info("%s, %s", category, price)

	err = products.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
