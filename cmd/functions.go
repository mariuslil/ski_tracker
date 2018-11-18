package main

import (
	"os"
)

func getPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	return ":" + port
}