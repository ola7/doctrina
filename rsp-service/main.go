package main

// Primary implementation of the Rock-Scissor-Paper (RSP) application.

import (
	"flag"

	"./config"
	"./dbclient"
	"./service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appname = "rsp-service"

func init() {

	// read command line flags
	profile := flag.String("profile", "test", "Environment profile like 'dev', 'test', etc")
	configServerUrl := flag.String("configServerUrl", "http://configserver:8888",
		"Address to config server (deployed within the swarm/paas)")
	configBranch := flag.String("configBranch", "master", "Git branch to fetch files from (typically 'master')")
	flag.Parse()

	// pass to viper to help read config/properties to be populated
	// in the main func
	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
}

func main() {

	logrus.Println("Starting application", appname)

	// load the config based on application flags
	// (populates viper)
	config.LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appname,
		viper.GetString("profile"),
		viper.GetString("configBranch"))

	initializeBoltClient()

	go config.StartListener(appname, viper.GetString("amqp_server_url"), viper.GetString("config_event_bus"))

	// port is extracted from config server via viper
	service.StartHTTPService(viper.GetString("server_port"))
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.SeedFakeUsers(100)
}
