package main

import (
	"context"
	"flag"
	"log"
	"microAPI/handlers"
	"microAPI/initializers"
	logging "microAPI/logger"
	"microAPI/middlewares"
	"microAPI/router"
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
		Handler:      middlewares.Recover(logger)(middlewares.Log(logger)(mux)),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}

	apiV1 := router.NewRouteGroup("/api/v1", mux, logger)

	// Создаем обработчики
	productsHandler := handlers.NewProductHandler(logger)

	// Регистрация маршрутов
	router.RegisterHandlers(apiV1, productsHandler)

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
