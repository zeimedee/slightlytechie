package main

import (
	// "log"

	"os"

	// "github.com/joho/godotenv"
	"github.com/zeimedee/slightlytechie/router"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading env file")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := router.SetUpRouter()

	router.Run(port)
}
