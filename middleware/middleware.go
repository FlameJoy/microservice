package middleware

import (
	"context"
	"microsvc/common/utils"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	bearerPrefix = "Bearer "
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

type KeyUser struct{}

func AuthMW(logger *utils.CustomLogger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")
			tokenStr = tokenStr[len(bearerPrefix):]
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
				}
				secretKey := os.Getenv("SECRET")
				return []byte(secretKey), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "token cannot be parsed or isn't valid", http.StatusUnauthorized)
				return
				// return nil, errors.New("token cannot be parsed or isn't valid")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "jwt.Claims is nil, not jwt.MapClaims", http.StatusUnauthorized)
				return
			}
			userID, ok := claims["user_id"].(string)
			if !ok {
				http.Error(w, "can't retrive a user id from token", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(context.Background(), KeyUser{}, userID)
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}
