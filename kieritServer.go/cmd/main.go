package main

import (
	"example.com/m/modules/engine"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := engine.New()
	app.Run()
}
