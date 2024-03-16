package startup

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type envValue struct{
	Host         string
	Port         string
	DbUser      string
	DbName      string
	DbPassword  string
	DbHost      string
	DbPort      string
	BrokerKafka string
	TopicKafka  string
}


func Startup(loggerInit *logrus.Logger)(dbInit *gorm.DB, envInit envValue, kafkaJournalInit KafkaConfig,  err error){
	loggerInit.Info("Startup...")
	godotenv.Load("service.env")
	godotenv.Load("config.env")
	godotenv.Load("kafka_config.env")

	envInit = envValue{
		Host		: os.Getenv("SERVICE_HOST"),
		Port        : os.Getenv("SERVICE_PORT"),
		DbUser      : os.Getenv("POSTGRES_USER"),
		DbName      : os.Getenv("POSTGRES_DB"),
		DbPassword  : os.Getenv("POSTGRES_PASSWORD"),
		DbHost      : os.Getenv("POSTGRES_HOST"),
		DbPort      : os.Getenv("POSTGRES_PORT"),
		BrokerKafka : os.Getenv("BROKER"),
		TopicKafka  : os.Getenv("TOPIC"),
	}


	loggerInit.Info("Initialize Database...")
	dbInit, err = InitializeDB(
		envInit.DbUser,
		envInit.DbName,
		envInit.DbPassword,
		envInit.DbHost,
		envInit.DbPort,
	)
	if err != nil{
		remark := "Error when initializing Database"
		loggerInit.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	loggerInit.Info("Initialize Kafka Topic...")
	err = CreateKafkaTopic(
		loggerInit,
		envInit.BrokerKafka,
		envInit.TopicKafka,
	)
	if err != nil{
		remark := "Error when initializing Kafka Topic"
		loggerInit.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	kafkaJournalInit = KafkaConfig{
		Logger: loggerInit,
		Broker: envInit.BrokerKafka,
		Topic: envInit.TopicKafka,
	}

	return
}

