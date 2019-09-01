package main

import "os"

func setupEnvironment() {
	if os.Getenv("DB_LOG_VERBOSE") == "" {
		_ = os.Setenv("DB_LOG_VERBOSE", "true")
	}
}
