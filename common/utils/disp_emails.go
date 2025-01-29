package utils

import (
	"bufio"
	"log"
	"os"
)

func DispEmailDomains() (dispEmailDomains []string) {
	file, err := os.Open("../disposable_email_blocklist.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dispEmailDomains = append(dispEmailDomains, scanner.Text())
	}
	return dispEmailDomains
}
