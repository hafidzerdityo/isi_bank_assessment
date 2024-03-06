package consumer

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hafidzresttemplate.com/datastore"
	"hafidzresttemplate.com/services"
)

type ConsumerSetup struct {
    Logger *logrus.Logger
	Services *services.ServiceSetup
}

func NewConsumerSetup(loggerInit *logrus.Logger, db *gorm.DB)(consumerSetup ConsumerSetup) {
	consumerSetup = ConsumerSetup{
		Logger: loggerInit,
		Services: &services.ServiceSetup{
			Logger: loggerInit,
			Db: db,
			Datastore: &datastore.DatastoreSetup{
				Logger: loggerInit,
			},
		},
	}
    return 
}


func InitConsumer(loggerInit *logrus.Logger, db *gorm.DB, brokerKafka string, topicKafka string, groupKafka string)() {
	consumerSetup := NewConsumerSetup(loggerInit, db)
	ctx := context.Background()
	consumerSetup.CreateJournalLoop(ctx, topicKafka, brokerKafka, groupKafka)
	
}




