package config

import (
	"flag"
	"os"
)

var (
	FlagRunAddr           string
	FlagLogLevel          = "info"
	FlagDatabaseDsn       string
	FlagAccrualSystemAddr string

	PasswordSalt string
	JWTSecret    string
)

func ParseFlags() {
	flag.StringVar(&FlagRunAddr, "a", "", "address and port to run server")
	flag.StringVar(&FlagDatabaseDsn, "d", "", "database dsn")
	flag.StringVar(&FlagAccrualSystemAddr, "r", "", "accrual system address")

	flag.Parse()

	if serverAddress := os.Getenv("RUN_ADDRESS"); serverAddress != "" {
		FlagRunAddr = serverAddress
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

	if envAccrualSystemAddr := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrualSystemAddr != "" {
		FlagAccrualSystemAddr = envAccrualSystemAddr
	}
}
