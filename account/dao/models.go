package dao

import (
	"time"
)

type Customer struct{
	ID          int       `gorm:"primaryKey"`
	Nama          string    `gorm:"size:255;not null"`
	Nik          string    `gorm:"size:255;not null;unique"`
	NoHp          string    `gorm:"size:255;not null;unique"`
	CreatedAt     time.Time 
    UpdatedAt     *time.Time `gorm:"autoUpdateTime:false"`
    IsDeleted     bool 
	Account        *Account     `gorm:"foreignKey:IdNasabah;references:ID"`
}

type Account struct{
	ID          int       `gorm:"primaryKey"`
	IdNasabah          int       `gorm:"not null"`
	NoRekening          string    `gorm:"size:255;not null;unique"`
	HashedPin          string    `gorm:"size:255;not null"`
	Saldo      float64     `gorm:"type:numeric(10,2);not null"`
	CreatedAt     time.Time 
    UpdatedAt     *time.Time `gorm:"autoUpdateTime:false"`
    IsDeleted     bool 
	Transaction        *Transaction     `gorm:"foreignKey:IdRekening;references:ID"`
}

type Transaction struct{
	ID          int       `gorm:"primaryKey"`
	IdRekening          int       `gorm:"not null"`
	JenisTransaksi          string    `gorm:"size:1;not null"`
	Nominal      float64     `gorm:"type:numeric(10,2);not null"`
	Waktu     time.Time 
	NomorRekeningTujuan *string    `gorm:"size:255"`
}


func (Customer) TableName() string {
	return "customer"
}

func (Account) TableName() string {
    return "account"
}

func (Transaction) TableName() string {
    return "transaction"
}

