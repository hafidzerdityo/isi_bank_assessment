package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/api"
	"hafidzresttemplate.com/logs"
	"hafidzresttemplate.com/startup"
)


func main() {

	loggerInit := logs.InitLog()
	dbInit, err := startup.Startup(loggerInit)
	if err != nil{
		remark := "Start Up Failed"
		loggerInit.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	app := api.InitApi(loggerInit, dbInit)
	app.Listen(fmt.Sprintf("%v:%v",startup.HOST, startup.PORT))
	loggerInit.Info("Service Started!")
}
