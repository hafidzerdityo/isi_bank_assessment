package startup

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
)

type KafkaConfig struct {
	Logger *logrus.Logger
    Broker string
    Topic  string
}


func (k *KafkaConfig)ProduceJournal(message dao.KafkaProducer) (err error) {

    w := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{k.Broker},
        Topic:    k.Topic,
        Balancer: &kafka.LeastBytes{},
    })

    defer w.Close()

    msgJSON, err := json.Marshal(message)

    ctx := context.Background()

    if err != nil {
		remark := "Error When marshalling message" 
		k.Logger.Error(
			logrus.Fields{"kafka_message_error": err.Error()}, nil, remark,
		)
    }

    err = w.WriteMessages(ctx, kafka.Message{Value: msgJSON})

    if err != nil {
		remark := "Error When writing Message" 
		k.Logger.Error(
			logrus.Fields{"kafka_message_error": err.Error()}, nil, remark,
		)
		return
    }

	remark := "Kafka Message Sent Succesfully"
	k.Logger.Info(	
		logrus.Fields{"kafka_message": fmt.Sprintf("%+v", message)}, nil, remark,
	)

    return
}


func CreateKafkaTopic(loggerInit *logrus.Logger, broker string, topic string)(err error) {
	loggerInit.Info("Initiate Kafka Topic...")

    conn, err := kafka.Dial("tcp", broker)
    if err != nil {
		remark := "Gagal Dial Kafka"
        loggerInit.Error(
            logrus.Fields{"error": err.Error()}, nil, remark,
        )
    }

    defer conn.Close()

    controller, err := conn.Controller()
    if err != nil {
        panic(err.Error())
    }
    var controllerConn *kafka.Conn
    controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
    if err != nil {
		remark := "Error dial kafka controller"
        loggerInit.Error(
            logrus.Fields{"error": err.Error()}, nil, remark,
        )
    }
    defer controllerConn.Close()

    topicConfigs := []kafka.TopicConfig{
        {
            Topic:             topic,
            NumPartitions:     1,
            ReplicationFactor: 1,
        },
    }

    err = controllerConn.CreateTopics(topicConfigs...)
    if err != nil {
		remark := "Gagal controllerConn CreateTopics"
        loggerInit.Error(
            logrus.Fields{"error": err.Error()}, nil, remark,
        )
    }
	return
}
