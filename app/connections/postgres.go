package connections

import (
	"fmt"
	"jojonomic/app/config"
	modelHarga "jojonomic/modules/utilities/models/harga"
	modelRekening "jojonomic/modules/utilities/models/rekening"
	modelTransaski "jojonomic/modules/utilities/models/transaksi"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	var err error
	p := config.Env("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Env("DB_HOST"), port, config.Env("DB_USER"), config.Env("DB_PASSWORD"), config.Env("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect database")
	}
	logrus.Infof("Connected to database")

	// Migrate Database
	DB.AutoMigrate(modelHarga.Harga{})
	DB.AutoMigrate(modelRekening.Rekening{})
	DB.AutoMigrate(modelTransaski.Transaksi{})

}
