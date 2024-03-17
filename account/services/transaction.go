package services

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
)



func (s *ServiceSetup)CreateTabung(reqPayload dao.CreateTabungTarikUpdate) (appResponse dao.SaldoRes, remark string, err error) {
	s.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: CreateTabung Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		remark = "Error When Initializing DB"
		return
	}

	// check if account exist
	accountData, err := s.Datastore.GetAccount(s.Db, reqPayload.NoRekening)
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
		err = fmt.Errorf("account exist error")
		remark = "no_rekening belum terdaftar" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	// update account saldo (increase)
	saldo, err := s.Datastore.UpdateSaldo(tx, reqPayload)
	if err != nil {
		tx.Rollback()
		remark = "Data Update Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	// create catatan transfer
	var insertCatatanParam dao.Transaction
	insertCatatanParam.IdRekening = accountData.ID
	insertCatatanParam.JenisTransaksi = "C"
	insertCatatanParam.Nominal = reqPayload.Nominal
	insertCatatanParam.Waktu = time.Now()
	insertCatatanParam.NomorRekeningTujuan = nil
	err = s.Datastore.InsertCatatan(tx, insertCatatanParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Insertion Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	tx.Commit()

	var kafkaMessageParam dao.KafkaProducer
	kafkaMessageParam.TanggalTransaksi = insertCatatanParam.Waktu
	kafkaMessageParam.NoRekeningKredit = accountData.NoRekening
	kafkaMessageParam.NoRekeningDebit = ""
	kafkaMessageParam.NominalKredit = insertCatatanParam.Nominal
	kafkaMessageParam.NominalDebit = 0

	// send message to kafka
	err = s.KafkaJournal.ProduceJournal(kafkaMessageParam)
	if err != nil {
		tx.Rollback()
		remark = "Failed to send message to kafka"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	appResponse.Saldo = &saldo

	remark = "END: CreateTabung Service"
	s.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", appResponse)}, nil, remark,
	)
	return
}


func (s *ServiceSetup)CreateTarik(reqPayload dao.CreateTabungTarikUpdate) (appResponse dao.SaldoRes, remark string, err error) {
	s.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: CreateTarik Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		remark = "Error When Initializing DB"
		return
	}

	// check if account exist
	accountData, err := s.Datastore.GetAccount(s.Db, reqPayload.NoRekening)
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
		err = fmt.Errorf("account exist error")
		remark = "no_rekening belum terdaftar" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	if accountData.Saldo < reqPayload.Nominal{
		tx.Rollback()
		err = fmt.Errorf("insufficient saldo error")
		remark = "maaf, saldo tidak cukup" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	// update account saldo (decrease)
	var decreaseSaldoParam dao.CreateTabungTarikUpdate
	decreaseSaldoParam.Nominal = -1 * reqPayload.Nominal
	decreaseSaldoParam.NoRekening = reqPayload.NoRekening
	saldo, err := s.Datastore.UpdateSaldo(tx, decreaseSaldoParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Update Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	// create catatan transfer
	var insertCatatanParam dao.Transaction
	insertCatatanParam.IdRekening = accountData.ID
	insertCatatanParam.JenisTransaksi = "D"
	insertCatatanParam.Nominal = reqPayload.Nominal
	insertCatatanParam.Waktu = time.Now()
	insertCatatanParam.NomorRekeningTujuan = nil
	err = s.Datastore.InsertCatatan(tx, insertCatatanParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Insertion Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}


	tx.Commit()

	var kafkaMessageParam dao.KafkaProducer
	kafkaMessageParam.TanggalTransaksi = insertCatatanParam.Waktu
	kafkaMessageParam.NoRekeningKredit = ""
	kafkaMessageParam.NoRekeningDebit = accountData.NoRekening
	kafkaMessageParam.NominalKredit = 0
	kafkaMessageParam.NominalDebit = insertCatatanParam.Nominal

	// send message to kafka
	err = s.KafkaJournal.ProduceJournal(kafkaMessageParam)
	if err != nil {
		tx.Rollback()
		remark = "Failed to send message to kafka"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	appResponse.Saldo = &saldo

	remark = "END: CreateTarik Service"
	s.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", appResponse)}, nil, remark,
	)
	return
}


func (s *ServiceSetup)CreateTransfer(reqPayload dao.CreateTransferUpdate) (appResponse dao.SaldoRes, remark string, err error) {
	s.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: CreateTransfer Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		remark = "Error When Initializing DB"
		return
	}
	// check if sender account exist
	accountSenderData, err := s.Datastore.GetAccount(s.Db, reqPayload.NoRekeningAsal)
	if err != nil {
		tx.Rollback()
		remark = "Data Get Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	if accountSenderData.ID == 0{
		tx.Rollback()
		err = fmt.Errorf("account exist error")
		remark = "no_rekening asal belum terdaftar" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	// check if benefciary account exist
	accountReceiverData, err := s.Datastore.GetAccount(s.Db, reqPayload.NoRekeningTujuan)
	if err != nil {
		tx.Rollback()
		remark = "Data Get Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	if accountReceiverData.ID == 0{
		tx.Rollback()
		err = fmt.Errorf("account exist error")
		remark = "no_rekening tujuan belum terdaftar" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	if accountSenderData.Saldo < reqPayload.Nominal{
		tx.Rollback()
		err = fmt.Errorf("insufficient saldo error")
		remark = "maaf, saldo tidak cukup" 
		s.Logger.Error(
			logrus.Fields{"validation_error": err.Error()}, nil, remark,
		)
		return
	}

	// update sender account saldo (decrease)
	var decreaseSaldoParam dao.CreateTabungTarikUpdate 
	decreaseSaldoParam.NoRekening = reqPayload.NoRekeningAsal
	decreaseSaldoParam.Nominal = -1 * reqPayload.Nominal

	saldo, err := s.Datastore.UpdateSaldo(tx, decreaseSaldoParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Update Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}
	
	// update account saldo (increase)
	var increaseSaldoParam dao.CreateTabungTarikUpdate 
	increaseSaldoParam.NoRekening = reqPayload.NoRekeningTujuan
	increaseSaldoParam.Nominal = reqPayload.Nominal
	_, err = s.Datastore.UpdateSaldo(tx, increaseSaldoParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Update Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	catatanTime := time.Now()

	// create catatan transfer sender
	var insertCatatanParam dao.Transaction
	insertCatatanParam.IdRekening = accountSenderData.ID
	insertCatatanParam.JenisTransaksi = "T"
	insertCatatanParam.Nominal = reqPayload.Nominal
	insertCatatanParam.Waktu = catatanTime
	insertCatatanParam.NomorRekeningTujuan = &accountReceiverData.NoRekening
	err = s.Datastore.InsertCatatan(tx, insertCatatanParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Insertion Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	// create catatan transfer receiver
	insertCatatanParam.IdRekening = accountReceiverData.ID
	insertCatatanParam.JenisTransaksi = "T"
	insertCatatanParam.Nominal = reqPayload.Nominal
	insertCatatanParam.Waktu = catatanTime
	insertCatatanParam.NomorRekeningTujuan = nil
	err = s.Datastore.InsertCatatan(tx, insertCatatanParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Insertion Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	tx.Commit()

	var kafkaMessageParam dao.KafkaProducer
	kafkaMessageParam.TanggalTransaksi = insertCatatanParam.Waktu
	kafkaMessageParam.NoRekeningKredit = accountSenderData.NoRekening
	kafkaMessageParam.NoRekeningDebit = accountReceiverData.NoRekening
	kafkaMessageParam.NominalKredit = insertCatatanParam.Nominal
	kafkaMessageParam.NominalDebit = insertCatatanParam.Nominal

	// send message to kafka
	err = s.KafkaJournal.ProduceJournal(kafkaMessageParam)
	if err != nil {
		tx.Rollback()
		remark = "Failed to send message to kafka"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}


	appResponse.Saldo = &saldo

	remark = "END: CreateTransfer Service"
	s.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", appResponse)}, nil, remark,
	)
	return
}

