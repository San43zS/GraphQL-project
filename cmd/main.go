package main

import (
	"GraphQL-project/internal/app"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
