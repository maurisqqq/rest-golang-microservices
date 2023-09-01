package connections

import (
	"jojonomic/app/config"
	"jojonomic/app/kafka/consumer"
	"jojonomic/app/kafka/producer"
	"os"

	buybackstorage "jojonomic/modules/utilities/microservices/buyback_storage"
	hargaStorage "jojonomic/modules/utilities/microservices/input_harga_storage"
	topupStorage "jojonomic/modules/utilities/microservices/topup_storage"

	hargaRepository "jojonomic/modules/utilities/repositories/harga"
	rekeningRepository "jojonomic/modules/utilities/repositories/rekening"
	transaksiRepository "jojonomic/modules/utilities/repositories/transaksi"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func PublisherConnection() (*producer.KafkaProducer, error) {
	kafkaConfig := config.GetKafkaConfig()
	producers, err := sarama.NewSyncProducer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Success create kafka sync-producer")

	kafka := &producer.KafkaProducer{
		Producer: producers,
	}
	return kafka, nil
}

func ConsumerConnection(db *gorm.DB) {
	kafkaConfig := config.GetKafkaConfig()
	consumers, err := sarama.NewConsumer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Error create kakfa consumer got error %v", err)
	}
	defer func() {
		if err := consumers.Close(); err != nil {
			panic(err)
		}
	}()
	//Repository
	hargaRepository := hargaRepository.NewRepository(db)
	transaksiRepository := transaksiRepository.NewRepository(db)
	rekeningRepository := rekeningRepository.NewRepository(db)
	//Kafka consumer
	kafkaConsumer := &consumer.KafkaConsumer{
		Consumer:              consumers,
		HargaStorageService:   hargaStorage.NewService(hargaRepository),
		TopupStorageService:   topupStorage.NewService(hargaRepository, rekeningRepository, transaksiRepository),
		BuybackStorageService: buybackstorage.NewService(hargaRepository, rekeningRepository, transaksiRepository),
	}
	signals := make(chan os.Signal, 1)

	listTopics := []string{"input_harga", "topup", "buyback"}
	kafkaConsumer.Consume(listTopics, signals, db)
}
