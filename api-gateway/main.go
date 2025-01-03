package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Завершение по сигналу
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	// HTTP-сервер
	wg.Add(1)
	go func() {
		defer wg.Done()
		StartHTTPServer(done)
	}()

	// gRPC-сервер
	wg.Add(1)
	go func() {
		defer wg.Done()
		StartGRPCServer(":50051", ":50052", done)
	}()

	<-done
	log.Println("Shutting down servers...")

	// Передаем сигнал всем горутинам для завершения
	close(done)

	wg.Wait()
	log.Println("Servers exited properly")
}
