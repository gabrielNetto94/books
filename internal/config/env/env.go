package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetVariable(variable string) string {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return os.Getenv(variable)
}

func IsProduction() bool {
	return GetVariable("ENVIRONMENT") == "production"
}
