package main

import (
	"flag"
	"log"
	"microsvc/common/utils"
	"os"
	"os/signal"
	"sync"
)

var (
	debug = flag.Bool("debug", false, "debugging code")
)

func main() {
	flag.Parse()

	logger := utils.NewLogger(utils.INFO, log.New(os.Stdout, "gateway-api ", log.LstdFlags), false)
	if *debug {
		logger.SetLevel(utils.DEBUG)
	}

	err := utils.LoadEnv("../.env")
	if err != nil {
		logger.Fatal("LoadEnv error: %s", err)
	}

	var wg sync.WaitGroup

	// Завершение по сигналу
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	// HTTP-сервер
	wg.Add(1)
	go func() {
		defer wg.Done()
		StartHTTPServer(done, logger)
	}()

	// gRPC-сервер
	wg.Add(1)
	go func() {
		defer wg.Done()
		StartGRPCServer(":50051", ":50052", ":50053", done, logger)
	}()

	<-done
	log.Println("Shutting down servers...")

	// Передаем сигнал всем горутинам для завершения
	close(done)

	wg.Wait()
	log.Println("Servers exited properly")
}
