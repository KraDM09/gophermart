package config

import (
	"flag"
	"os"
)

var (
	FlagRunAddr           string
	FlagLogLevel          string
	FlagDatabaseDsn       string
	FlagAccrualSystemAddr string

	PasswordSalt string
	JWTSecret    string
)

func ParseFlags() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&FlagLogLevel, "l", "info", "log level")
	flag.StringVar(&FlagDatabaseDsn, "d", "", "database dsn")
	flag.StringVar(&FlagAccrualSystemAddr, "r", ":8080", "accrual system address")

	flag.Parse()

	if serverAddress := os.Getenv("RUN_ADDRESS"); serverAddress != "" {
		FlagRunAddr = serverAddress
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		FlagLogLevel = envLogLevel
	}

	if envPasswordSalt := os.Getenv("PASSWORD_SALT"); envPasswordSalt != "" {
		PasswordSalt = envPasswordSalt
	}

	if envJWTSecret := os.Getenv("JWT_SECRET"); envJWTSecret != "" {
		JWTSecret = envJWTSecret
	}

	if envDatabaseDsn := os.Getenv("DATABASE_URI"); envDatabaseDsn != "" {
		FlagDatabaseDsn = envDatabaseDsn
	}

	if envAccrualSystemAddr := os.Getenv("ACCRUAL_SYSTEM_ADDR"); envAccrualSystemAddr != "" {
		FlagAccrualSystemAddr = envAccrualSystemAddr
	}
}
