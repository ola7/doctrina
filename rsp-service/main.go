package main

// Primary implementation of the Rock-Scissor-Paper (RSP) application.

import (
	"log"

	"./service"
)

var appname = "RSP-SERVICE"

func main() {
	log.Println("Starting application", appname)
	service.StartHTTPService("8080")
}
