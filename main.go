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

	router := router.SetUpRouter()

	router.Run(":" + port)
}
