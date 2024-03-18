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


func (s *ServiceSetup)AccountLogin(reqPayload dao.AccountLoginReq) (appResponse dao.AccountLoginRes, remark string, err error) {
	reqPayloadForLog := reqPayload
	reqPayloadForLog.Pin = "*REDACTED*"
	reqPayloadForLog.Password = "*REDACTED*"
	s.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayloadForLog)}, nil, "START: AccountLogin Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		remark = "Error When Initializing DB"
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

	// check if email exist
	loginData, err := s.Datastore.CheckEmailAndGetHashedPassword(s.Db, reqPayload)
	if err != nil {
		tx.Rollback()
		remark = "Data Get Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	if loginData.ID == 0{
		tx.Rollback()
		err = fmt.Errorf("email not found error")
		remark = "Email tidak ditemukan" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	if reqPayload.Pin == "" && reqPayload.Password == ""{
		tx.Rollback()
		err = fmt.Errorf("password and pin empty error")
		remark = "silahkan input password atau pin" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}
	
	// check if pin correct
	if reqPayload.Pin != ""{
		err = utils.VerifyPassword(reqPayload.Pin, loginData.HashedPin)
		if err != nil{
			tx.Rollback()
			err = fmt.Errorf("wrong pin error")
			remark = "pin salah" 
			s.Logger.Error(
				logrus.Fields{"validation_error": err.Error()}, nil, remark,
			)
			return
		}
	}

	// check if password correct
	if reqPayload.Password != ""{
		err = utils.VerifyPassword(reqPayload.Password, loginData.HashedPassword)
		if err != nil{
			tx.Rollback()
			err = fmt.Errorf("wrong password error")
			remark = "password salah" 
			s.Logger.Error(
				logrus.Fields{"validation_error": err.Error()}, nil, remark,
			)
			return
		}
	}

	// generate token
	var JWTFieldParam dao.JWTField
	JWTFieldParam.Email = reqPayload.Email
	JWTFieldParam.NoRekening = loginData.NoRekening
	JWTFieldParam.NoHp = loginData.NoHp
	tokenString, err := utils.CreateJWTToken(JWTFieldParam)
	if err != nil{
		tx.Rollback()
		err = fmt.Errorf("wrong password error")
		remark = "password salah" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	appResponse.AccessToken = tokenString
	appResponse.TokenType = "bearer"

	tx.Commit()

	remark = "END: AccountLogin Service"
	s.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", appResponse)}, nil, remark,
	)
	return
}

