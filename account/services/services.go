package services

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hafidzresttemplate.com/datastore"
	"hafidzresttemplate.com/startup"
)


type ServiceSetup struct{
	Logger *logrus.Logger
	Datastore *datastore.DatastoreSetup
	Db		*gorm.DB
	KafkaJournal startup.KafkaConfig
}



