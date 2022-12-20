package main

import (
	// "log"
	"os"

	"github.com/zeimedee/slightlytechie/router"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading env file")
	// }

	port := os.Getenv("PORT")

	router := router.SetUpRouter()

	router.Run(port)
}
