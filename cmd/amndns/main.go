package main

import (
	"github.com/alidandach/AmnDNS/internal/dns"
	"github.com/alidandach/AmnDNS/internal/logger"
	"github.com/sirupsen/logrus"
)

// ANSI escape codes for coloring the logs
const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {
	// Initialize the logger
	log := logger.GetLogger()

	// Log starting message with color
	log.WithFields(logrus.Fields{"package": "main"}).Info(Green + "Starting AmnDNS..." + Reset)

	// Initialize the DNS server
	server := dns.NewServer()

	// Start the DNS server
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start AmnDNS: %v", err)
	}
}
