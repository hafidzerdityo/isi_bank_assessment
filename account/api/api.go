package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hafidzresttemplate.com/datastore"
	"hafidzresttemplate.com/services"
)

type ApiSetup struct {
    Logger *logrus.Logger
	Services *services.ServiceSetup
}

func NewApiSetup(loggerInit *logrus.Logger, db *gorm.DB)(apiSet ApiSetup) {
	apiSet = ApiSetup{
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

func InitApi(loggerInit *logrus.Logger, db *gorm.DB)(app *fiber.App) {
	app = fiber.New()
	app.Use(logger.New())

	apiSetup := NewApiSetup(loggerInit, db)
	apiSetup.Logger.Info("Setting up api routes...")

	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user_management")
	trx := v1.Group("/transaction")
	inq := v1.Group("/inquiry")
	trx.Use(apiSetup.TransactionMiddleware())
	inq.Use(apiSetup.TransactionMiddleware())
	user.Post("/daftar", apiSetup.CreateUser)
	trx.Post("/tabung", apiSetup.CreateTabung)
	trx.Post("/tarik", apiSetup.CreateTarik)
	trx.Post("/transfer", apiSetup.CreateTransfer)
	inq.Get("/saldo/:no_rekening", apiSetup.GetSaldo)
	inq.Get("/mutasi/:no_rekening", apiSetup.GetMutasi)
	
	return 
}