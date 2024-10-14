package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	APIPort     string
	SessionAddr string
	ExcelAddr   string
	DBAddr      string
	ZeroAddr    string
	ZeroToken   string
	Debug       bool
}

func (c Config) Print() {
	msg := `{
	"api port": "%s",
	"session addr": "%s",
	"excel addr": "%s",
	"db addr": "%s",
	"zero addr": "%s",
	"zero token": "%s",
	"debug": "%t",
}`
	log.Printf(msg, c.APIPort, c.SessionAddr, c.ExcelAddr, c.DBAddr, c.ZeroAddr, c.ZeroToken, c.Debug)
}

func Configure() *Config {
	// own port
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("api port need check: is empty")
	}

	// grpc - excel_table
	excelTableAddr := os.Getenv("E_SRV_ADDR")
	excelTablePort := os.Getenv("E_SRV_PORT")
	if excelTableAddr == "" || excelTablePort == "" {
		log.Fatalf("excel_table_grpc address need check: got \"%s\" and \"%s\"", excelTableAddr, excelTablePort)
	}

	// grpc - session_manager
	sessionTableAddr := os.Getenv("S_SRV_ADDR")
	sessionTablePort := os.Getenv("S_SRV_PORT")
	if excelTableAddr == "" || excelTablePort == "" {
		log.Fatalf("session_manager_grpc address need check: got \"%s\" and \"%s\"", sessionTableAddr, sessionTablePort)
	}

	// database
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		log.Fatalf("postgres address: is empty")
	}

	// graphql-server - zero one
	zeroAddr := os.Getenv("ZERO_DOMAIN")
	zeroToken := os.Getenv("ZERO_TOKEN")
	if zeroAddr == "" || zeroToken == "" {
		log.Fatalf("zero_one address or token need check: got \"%s\" and \"%s\"", zeroAddr, zeroToken)
	}

	// if need additional logging
	debug := false
	debugStr := strings.ToLower(os.Getenv("REQ_LOG"))
	if debugStr == "true" {
		debug = true
	}

	return &Config{
		APIPort:     fmt.Sprintf(":%s", port),
		SessionAddr: fmt.Sprintf("%s:%s", sessionTableAddr, sessionTablePort),
		ExcelAddr:   fmt.Sprintf("%s:%s", excelTableAddr, excelTablePort),
		DBAddr:      db,
		ZeroAddr:    zeroAddr,
		ZeroToken:   zeroToken,
		Debug:       debug,
	}
}
