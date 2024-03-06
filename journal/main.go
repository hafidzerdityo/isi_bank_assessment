package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/consumer"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/logs"
)


func main() {
	loggerInit := logs.InitLog()
	godotenv.Load("service.env")
	godotenv.Load("config.env")
	godotenv.Load("kafka_config.env")

	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbInit, err := dao.InitializeDB(
		dbUser,
		dbName,
		dbPassword,
		dbHost,
		dbPort,
	)
	if err != nil{
		remark := "Error when initializing Database"
		loggerInit.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}


	brokerKafka := os.Getenv("BROKER")
	topicKafka := os.Getenv("TOPIC")
	groupKafka := os.Getenv("GROUP")

	consumer.InitConsumer(loggerInit, dbInit, brokerKafka, topicKafka, groupKafka)
}
