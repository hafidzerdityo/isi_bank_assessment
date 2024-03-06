package datastore

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hafidzresttemplate.com/dao"
)


func(d *DatastoreSetup) InsertCustomer(tx *gorm.DB, reqPayload dao.Customer)(err error){
	d.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: InsertCustomer Datastore",
	)

	if err = tx.Create(&reqPayload).Error; err != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}
	remark := "END: InsertCustomer Datastore"
	d.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", map[string]interface{}{})}, nil, remark,
	)

	return
}

func(d *DatastoreSetup) InsertAccount(tx *gorm.DB, reqPayload dao.Account)(err error){
	reqPayloadForLog := reqPayload
	reqPayloadForLog.HashedPin = "*REDACTED*"
	d.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayloadForLog)}, nil, "START: InsertAccount Datastore",
	)

	if err = tx.Create(&reqPayload).Error; err != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	remark := "END: InsertAccount Datastore"
	d.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v",  map[string]interface{}{})}, nil, remark,
	)

	return
}

func (d *DatastoreSetup) GetCustomer(tx *gorm.DB, reqPayload dao.Customer) (datastoreResponse dao.Customer, err error) {
	d.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: GetCustomer Datastore",
	)

	err = tx.Where("no_hp = ? OR nik = ?", reqPayload.NoHp, reqPayload.Nik).Find(&datastoreResponse).Error
	if err != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	remark := "END: GetCustomer Datastore"
	d.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", datastoreResponse)}, nil, remark,
	)

	return
}
