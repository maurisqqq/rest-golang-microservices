package producer

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func (p *KafkaProducer) SendData(topic string, msg []byte) error {

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}

	_, _, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		return err
	}
	logrus.Infof("Send data success, topic : %v", topic)
	return nil
}
