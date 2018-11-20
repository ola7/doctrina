package main

// Primary implementation of the Rock-Scissor-Paper (RSP) application.

import (
	"log"

	"./dbclient"
	"./service"
)

var appname = "RSP-SERVICE"

func main() {
	log.Println("Starting application", appname)
	initializeBoltClient()
	service.StartHTTPService("8989")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.SeedFakeUsers(100)
}
