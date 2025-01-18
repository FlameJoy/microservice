package main

import (
	"log"
	"microsvc/common/utils"
	"os"
)

func main() {
	logger := utils.NewLogger(utils.INFO, log.New(os.Stdout, "Order svc ", log.LstdFlags), false)

	StartGRPCServer(":50053", logger)
}
