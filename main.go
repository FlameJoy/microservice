package main

import (
	"context"
	"flag"
	"log"
	"microAPI/initializers"
	logging "microAPI/logger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	debug = flag.Bool("debug", false, "debug level")
)

func main() {
	flag.Parse()
	logger := logging.NewCustomLogger(log.New(os.Stdout, "", log.LstdFlags), logging.INFO)
	if *debug {
		logger.SetLevel(logging.DEBUG)
	}

	if err := initializers.LoadEnv(".env"); err != nil {
		logger.Warn("warning: %v", err)
	}

	mux := http.NewServeMux()

	server := http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      recoveringMW(logger)(loggingMW(logger)(mux)),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}

	// handler := NewHandler(logger)

	registerHandler(mux, logger)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		logger.Info("Server starts in port%s", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal("Can't listen and serve: %v", err)
		}
	}()

	<-done
	logger.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Shutdown error %v", err)
	}
	logger.Info("Server exited properly")
}
