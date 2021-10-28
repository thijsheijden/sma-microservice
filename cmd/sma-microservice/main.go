package main

import "sma-microservice/internal/api"

func main() {
	api := api.New()
	api.Start()
}
