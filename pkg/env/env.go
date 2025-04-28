package env

import (
	"os"

	"github.com/joho/godotenv"
)

// Loader defines an interface for environment variable management.
type Loader interface {
	Load(filenames ...string) error
	Get(key string) string
	MustGet(key string) string
	IsProduction() bool
}

type loader struct{}

// NewLoader returns a new instance of the env Loader.
func NewLoader() Loader {
	return &loader{}
}

// Load loads environment variables from specified .env files.
// If no filenames are provided, it loads ".env" by default.
func (l *loader) Load(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}
	return godotenv.Load(filenames...)
}

// Get returns the value of the environment variable, or an empty string if not set.
func (l *loader) Get(key string) string {
	return os.Getenv(key)
}

// MustGet returns the value of the environment variable, or panics if not set.
func (l *loader) MustGet(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("environment variable not set: " + key)
	}
	return value
}

// IsProduction checks if ENVIRONMENT variable is set to "production".
func (l *loader) IsProduction() bool {
	return l.Get("ENVIRONMENT") == "production"
}
