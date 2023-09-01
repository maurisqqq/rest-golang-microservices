package main

import (
	"jojonomic/app/config"
	"jojonomic/app/connections"
	router "jojonomic/modules/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("starting API server")

	connections.ConnectPostgres()
	db := connections.DB
	go connections.ConsumerConnection(db)

	logrus.Infof("🚀 Succeed Running On http://localhost ", config.Env("APP_PORT"))
	router.SetupRoutes(db)
}
