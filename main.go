package main

import (
	"jojonomic/app/config"
	"jojonomic/app/connections"
	router "jojonomic/modules/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting API server")

	connections.ConnectPostgres()
	db := connections.DB
	go connections.ConsumerConnection(db)

	logrus.Infof("🚀 Succeed Running On http://localhost %v", config.Env("APP_PORT"))
	router.SetupRoutes(db)
}
