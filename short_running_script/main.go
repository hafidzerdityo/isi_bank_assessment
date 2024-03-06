package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"hafidzresttemplate.com/database"
)


func main() {
	godotenv.Load("config.env")

	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbInit, err := database.InitializeDB(
		dbUser,
		dbName,
		dbPassword,
		dbHost,
		dbPort,
	)
	if err != nil{
		remark := "Error when initializing Database"
		panic(remark)
	}

	outputData, err := database.GetRekap(dbInit)
	if err != nil{
		remark := "Error when getting Data"
		panic(remark)
	}
	fmt.Printf("%+v", outputData)

}
