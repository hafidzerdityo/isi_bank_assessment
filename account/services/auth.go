package services

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/pkg/utils"
)


func (s *ServiceSetup)CheckAccountAndPin(reqPayload dao.CheckAccountAndPinReq) (isExist bool, remark string, err error) {
	reqPayloadForLog := reqPayload
	reqPayloadForLog.Pin = "*REDACTED*"
	s.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayloadForLog)}, nil, "START: CheckAccountAndPin Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		remark = "Error When Initializing DB"
		return
	}

	// check if account exist
	accountData, err := s.Datastore.GetAccount(s.Db, reqPayloadForLog.NoRekening)
	if err != nil {
		tx.Rollback()
		remark = "Data Get Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	if accountData.ID == 0{
		tx.Rollback()
		err = fmt.Errorf("account not exist error")
		remark = "no_rekening belum terdaftar" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	// check if pin correct
	err = utils.VerifyPassword(reqPayload.Pin, accountData.HashedPin)
	if err != nil{
		tx.Rollback()
		err = fmt.Errorf("wrong pin error")
		remark = "pin salah" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	tx.Commit()
	isExist = true

	remark = "END: CheckAccountAndPin Service"
	s.Logger.Info(
		logrus.Fields{"response": "Account Exist and Pin is correct"}, nil, remark,
	)
	return
}

