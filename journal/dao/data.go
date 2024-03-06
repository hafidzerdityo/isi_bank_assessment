package dao

import "time"

type KafkaConsumer struct {
	TanggalTransaksi time.Time `json:"no_rekening_asal"`
	NoRekeningKredit string    `json:"no_rekening_kredit"`
	NoRekeningDebit  string    `json:"no_rekening_debit"`
	NominalKredit    float64   `json:"nominal_kredit"`
	NominalDebit     float64   `json:"nominal_debit"`
}

type CreateJournalRes struct {
	Success bool `json:"success"`
}
