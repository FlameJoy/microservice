package main

import (
	logging "microAPI/logger"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func loggingMW(logger *logging.CustomLogger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("%s | %s | %s | %v ", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
		})
	}
}

func recoveringMW(logger *logging.CustomLogger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Warn("Recovering from panic: %v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
