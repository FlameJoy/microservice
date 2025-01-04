package middleware

import (
	"microsvc/common/utils"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func LoggerMW(logger *utils.CustomLogger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("%s | %s | %s | %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
		})
	}
}

func RecoverMW(logger *utils.CustomLogger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Warn("Recovered from panic: %v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
