package checker

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/mjarkk/framework-microwave/pkg/gui"
)

// Env checks if all the env variables are set
func Env() error {
	err := godotenv.Load()
	if err != nil {
		gui.CritErr("No env file found")
	}

	checks := []string{
		"MONGODB_URL",
		"MONGODB_DATABASE",
		"DEBUG",
	}
	for _, check := range checks {
		if len(os.Getenv(check)) == 0 {
			return errors.New(check)
		}
	}
	return nil
}
