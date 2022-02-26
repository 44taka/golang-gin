package infrastructure

import (
	"os"
)

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
		Port     string
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	// DB設定
	c.DB.Host = os.Getenv("POSTGRES_HOST")
	c.DB.Username = os.Getenv("POSTGRES_USER")
	c.DB.Password = os.Getenv("POSTGRES_PASSWORD")
	c.DB.DBName = os.Getenv("POSTGRES_DB")
	c.DB.Port = os.Getenv("POSTGRES_PORT")

	// ポート番号
	c.Routing.Port = os.Getenv("PORT")
	if c.Routing.Port == "" {
		c.Routing.Port = "8080"
	}

	return c
}

func NewTestConfig() *Config {
	c := new(Config)

	// DB設定
	c.DB.Host = os.Getenv("POSTGRES_HOST")
	c.DB.Username = os.Getenv("POSTGRES_HOST")
	c.DB.Password = os.Getenv("POSTGRES_HOST")
	c.DB.DBName = os.Getenv("POSTGRES_HOST")
	c.DB.Port = os.Getenv("POSTGRES_PORT")

	// ポート番号
	c.Routing.Port = os.Getenv("PORT")
	if c.Routing.Port == "" {
		c.Routing.Port = "8080"
	}

	return c
}
