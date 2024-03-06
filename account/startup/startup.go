package startup

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	HOST         string
	PORT         string
	DB_USER      string
	DB_NAME      string
	DB_PASSWORD  string
	DB_HOST      string
	DB_PORT      string
	BROKER_KAFKA string
	TOPIC_KAFKA  string
)

func Startup(loggerInit *logrus.Logger)(dbInit *gorm.DB, err error){
	loggerInit.Info("Startup...")
	godotenv.Load("service.env")
	godotenv.Load("config.env")
	godotenv.Load("kafka_config.env")
	HOST = os.Getenv("SERVICE_HOST")
	PORT = os.Getenv("SERVICE_PORT")
	DB_USER = os.Getenv("POSTGRES_USER")
	DB_NAME = os.Getenv("POSTGRES_DB")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	DB_HOST = os.Getenv("POSTGRES_HOST")
	DB_PORT = os.Getenv("POSTGRES_PORT")
	BROKER_KAFKA = os.Getenv("BROKER")
	TOPIC_KAFKA = os.Getenv("TOPIC")

	loggerInit.Info("Initialize Database...")
	dbInit, err = InitializeDB(
		DB_USER,
		DB_NAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
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
		BROKER_KAFKA,
		TOPIC_KAFKA,
	)
	if err != nil{
		remark := "Error when initializing Kafka Topic"
		loggerInit.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	return
}

