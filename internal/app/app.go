package app

import (
	"GraphQL-project/internal/app/server"
	"github.com/joho/godotenv"
)

func Start() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := server.New().Start(); err != nil {
		return err
	}

	return nil
}
