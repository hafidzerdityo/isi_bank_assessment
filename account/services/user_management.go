package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/pkg/utils"
)


func (s *ServiceSetup)CreateUser(reqPayload dao.CreateCustReq) (appResponse dao.CreateCustRes, remark string, err error) {
	reqPayloadForLog := reqPayload
	reqPayloadForLog.Pin = "*REDACTED*"
	s.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayloadForLog)}, nil, "START: CreateUser Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		remark = "Error When Initializing DB"
		return
	}

	if len(reqPayload.Pin) != 6{
		tx.Rollback()
		err = fmt.Errorf("pin validation error")
		remark = "Pin Length Must be 6" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	if !utils.IsDigit(reqPayload.Pin){
		tx.Rollback()
		err = fmt.Errorf("pin validation error")
		remark = "Pin Must be a String of Digit" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	convertedNoHp, match := utils.ValidatePhoneNumber(reqPayload.NoHp)
	if !match{
		tx.Rollback()
		err = fmt.Errorf("no_hp validation error")
		remark = "Format No Hp Tidak Sesuai" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	reqPayload.NoHp = convertedNoHp

	var customerInsertParam dao.Customer

	customerInsertParam.Nama = strings.ToUpper(reqPayload.Nama)
	customerInsertParam.Nik = reqPayload.Nik
	customerInsertParam.NoHp = reqPayload.NoHp
	isValidEmail := utils.ValidateEmail(reqPayload.Email)
	if !isValidEmail{
		tx.Rollback()
		err = fmt.Errorf("email validation error")
		remark = "Format E-Mail Tidak Sesuai" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}
	customerInsertParam.Email = reqPayload.Email
	customerInsertParam.CreatedAt = time.Now()
	customerInsertParam.UpdatedAt = nil
	customerInsertParam.IsDeleted = false

	// check if customer exist
	customerData, err := s.Datastore.GetCustomer(s.Db, customerInsertParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Get Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	if customerData.ID != 0{
		tx.Rollback()
		err = fmt.Errorf("customer exist error")
		remark = "No Hp atau Nik sudah terdaftar" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}
	
	// insert customer data
	err = s.Datastore.InsertCustomer(tx, customerInsertParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Insertion Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	// get ID from inserted customer data
	customerData, err = s.Datastore.GetCustomer(tx, customerInsertParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Get Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	var accountInsertParam dao.Account
	accountInsertParam.IdNasabah = customerData.ID
	getHashedPin, err := utils.HashPassword(reqPayload.Pin)
	if err != nil {
		tx.Rollback()
		remark = "Error when Hashing PIN"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	accountInsertParam.HashedPin = getHashedPin
	getHashedPassword, err := utils.HashPassword(reqPayload.Password)
	if err != nil {
		tx.Rollback()
		remark = "Error when Hashing Password"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	accountInsertParam.HashedPassword = getHashedPassword
	noRekening := utils.GenerateNumericUUID(9)
	accountInsertParam.NoRekening = noRekening
	accountInsertParam.Saldo = 0
	accountInsertParam.CreatedAt = time.Now()
	accountInsertParam.UpdatedAt = nil
	accountInsertParam.IsDeleted = false

	// insert account data
	err = s.Datastore.InsertAccount(tx, accountInsertParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Insertion Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	appResponse.NoRekening = &noRekening

	tx.Commit()

	remark = "END: CreateUser Service"
	s.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", appResponse)}, nil, remark,
	)
	return
}

