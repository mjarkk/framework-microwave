package checker

import (
	"errors"
	"os"
)

// Env checks if all the env variables are set
func Env() error {
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
