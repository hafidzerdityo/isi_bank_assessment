package startup

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

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
