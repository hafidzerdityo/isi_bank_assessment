package dao

import "time"

type CreateCustReq struct {
	Nama string `json:"nama" validate:"required"`
	Nik  string `json:"nik" validate:"required"`
	NoHp string `json:"no_hp" validate:"required"`
	Pin  string `json:"pin" validate:"required"`
}
type CreateCustRes struct {
	NoRekening *string `json:"no_rekening"`
}

type SaldoRes struct {
	Saldo *float64 `json:"saldo"`
}

type CreateTabungTarikReq struct {
	NoRekening string  `json:"no_rekening" validate:"required"`
	Nominal    float64 `json:"nominal" validate:"required,gt=0"`
}

type CreateTransferReq struct {
	NoRekeningAsal   string  `json:"no_rekening_asal" validate:"required"`
	NoRekeningTujuan string  `json:"no_rekening_tujuan" validate:"required"`
	Nominal          float64 `json:"nominal" validate:"required,gt=0"`
}

type KafkaProducer struct {
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	NoRekeningKredit string    `json:"no_rekening_kredit"`
	NoRekeningDebit string    `json:"no_rekening_debit"`
	NominalKredit          float64   `json:"nominal_kredit"`
	NominalDebit          float64   `json:"nominal_debit"`
}

type NoRekeningReq struct {
	NoRekening string  `json:"no_rekening" validate:"required"`
}

type MutasiRes struct {
	Waktu time.Time `json:"waktu"`
	JenisTransaksi string    `json:"kode_transaksi"`
	Nominal float64    `json:"nominal"`
}

type CheckAccountAndPinReq struct {
	NoRekening string
	Pin    string
}