package main

import (
	"context"
	"microsvc/common/utils"
	"microsvc/middleware"
	"net/http"
	"os"
	"time"
)

func StartHTTPServer(done chan os.Signal, logger *utils.CustomLogger) {

	mux := http.NewServeMux()

	port := os.Getenv("GATEWAY_PORT")
	domain := os.Getenv("DOMAIN_NAME")

	server := http.Server{
		Addr:         port,
		Handler:      middleware.RecoverMW(logger)(middleware.LoggerMW(logger)(mux)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	h := NewHandler(logger)

	registerHandlers(h, mux)

	go func() {
		logger.Info("API gateway starts in %s%s", domain, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Could not listen server on %s: %v\n", port, err)
		}
	}()

	<-done
	logger.Info("http Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown failed: %v\n", err)
	}
	logger.Info("http Server exited properly")
}
