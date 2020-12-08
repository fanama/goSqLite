package tools

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//InitValues ...
func InitValues() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Name = os.Getenv("NAME")
}
