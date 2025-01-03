package main

import (
	"context"
	"flag"
	"log"
	"microsvc/common/utils"
	"microsvc/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	debug = flag.Bool("debug", false, "debugging code")
)

func main() {
	flag.Parse()

	logger := utils.NewLogger(utils.INFO, log.New(os.Stdout, "auth-service ", log.LstdFlags), false)
	if *debug {
		logger.SetLevel(utils.DEBUG)
	}

	mux := http.NewServeMux()

	err := utils.LoadEnv("./.env")
	if err != nil {
		logger.Fatal("LoadEnv error: %s", err)
	}

	port := os.Getenv("AUTH_SVC_PORT")
	domain := os.Getenv("DOMAIN_NAME")

	server := http.Server{
		Addr:         port,
		Handler:      middleware.RecoverMW(logger)(middleware.LoggerMW(logger)(mux)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		logger.Info("API gateway starts in %s%s", domain, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Could not listen server on %s: %v\n", port, err)
		}
	}()

	<-done
	logger.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown failed: %v\n", err)
	}
	logger.Info("Server exited properly")
}
