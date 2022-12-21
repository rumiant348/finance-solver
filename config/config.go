package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetConfig() Config {
	if checkIfProdFromEnv() {
		port, err := strconv.Atoi(os.Getenv("port"))
		if err != nil {
			panic(err)
		}
		fmt.Println("Running with prod env")
		return Config{
			Env:            "prod",
			Pepper:         os.Getenv("pepper"),
			HMACKey:        os.Getenv("hmac_key"),
			Port:           port,
			connectionInfo: getConnectionInfoFromEnv(),
		}
	}
	fmt.Println("Running with dev env")
	return Config{
		Env:            "dev",
		Pepper:         "secret-secret-secret",
		HMACKey:        "secret-hmac-hmacSecretKey",
		Port:           3000,
		connectionInfo: getConnectionInfoDev(),
	}
}

type Config struct {
	Env            string
	Pepper         string
	HMACKey        string
	Port           int
	connectionInfo string
}

func checkIfProdFromEnv() bool {
	return os.Getenv("env") == "prod"
}

func (c *Config) IsProd() bool {
	return c.Env == "prod"
}

func (c *Config) Dialect() string {
	return "postgres"
}

func (c *Config) ConnectionInfo() string {
	return c.connectionInfo
}

func getConnectionInfoFromEnv() string {
	connectionInfo := os.Getenv("DATABASE_URL")
	if connectionInfo == "" {
		panic("Empty env var DATABASE_URL")
	}
	return fmt.Sprintf("%s?sslmode=disable", connectionInfo)
}

func getConnectionInfoDev() string {
	return "postgresql://localhost:5432/finance_solver_dev?sslmode=disable"
}
