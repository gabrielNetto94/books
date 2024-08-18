package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetVariable(variable string) string {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(variable)
}
