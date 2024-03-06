package services

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
)


func (s *ServiceSetup)CreateJournal(reqPayload dao.KafkaConsumer) (appResponse dao.CreateJournalRes, remark string, err error) {
	s.Logger.Info(
		logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: CreateJournal Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		remark = "Error When Initializing DB"
		return
	}

	// insert journal data
	var insertJournalParam dao.Journal
	insertJournalParam.TanggalTransaksi = reqPayload.TanggalTransaksi
	insertJournalParam.NoRekeningKredit = reqPayload.NoRekeningKredit
	insertJournalParam.NoRekeningDebit = reqPayload.NoRekeningDebit
	insertJournalParam.NominalKredit = reqPayload.NominalKredit
	insertJournalParam.NominalDebit = reqPayload.NominalDebit
	err = s.Datastore.InsertJournal(tx, insertJournalParam)
	if err != nil {
		tx.Rollback()
		remark = "Data Insertion Error"
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, remark,
		)
		return
	}

	appResponse.Success = true

	tx.Commit()

	remark = "END: CreateJournal Service"
	s.Logger.Info(
		logrus.Fields{"response": fmt.Sprintf("%+v", appResponse)}, nil, remark,
	)
	return
}

