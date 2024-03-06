package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
)

func (a *ConsumerSetup) CreateJournalLoop(ctx context.Context, topic string, broker string, group string) (err error) {

	a.Logger.Info("Jurnal Loop Consumer Started")
	var reqKafka dao.KafkaConsumer


	config := kafka.ReaderConfig{
		Brokers:         []string{broker},
		GroupID:         group,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {

		m, err := reader.ReadMessage(ctx)

		if err != nil {
			remark := "error when receiving message"
			a.Logger.Error(
				logrus.Fields{"error": err.Error()}, nil, remark,
			)
			continue
		}

		value := m.Value

		a.Logger.Info(logrus.Fields{
			"topic":     m.Topic,
			"partition": m.Partition,
			"offset":    m.Offset,
			"message":   string(value),
		}, nil, "message")

		err = json.Unmarshal(value, &reqKafka)

		if err != nil {
			a.Logger.Error(logrus.Fields{
				"error": err.Error(),
			}, nil, "error parsing message kafka")
		}

		a.Logger.Info(logrus.Fields{
			"message":reqKafka,
		},nil,"kafka message")
		
		_, remark, err := a.Services.CreateJournal(reqKafka)
		if err != nil {
			a.Logger.Error(
				logrus.Fields{"error": err.Error()}, nil, remark,
			)
			return err
		}
	
	}

}
