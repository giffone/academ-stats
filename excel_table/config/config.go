package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	GRPCPort string
	DBAddr   string
	Debug    bool
}

func Configure() *Config {
	// own port
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("api port need check: is empty")
	}

	// database
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		log.Fatalf("postgres address: is empty")
	}

	// if need additional logging
	debug := false
	debugStr := strings.ToLower(os.Getenv("REQ_LOG"))
	if debugStr == "true" {
		debug = true
	}

	return &Config{
		GRPCPort:     fmt.Sprintf(":%s", port),
		DBAddr:      db,
		Debug:       debug,
	}
}
