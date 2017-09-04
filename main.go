package main

import (
"nvidia-gpu-monitoring/router"
"os"
)

func main() {
	router.RunApi(getPort())

}

func getPort() string {
	p := os.Getenv("PORT")

	if p != "" {
		return p
	} else {
		return "8080"
	}
}