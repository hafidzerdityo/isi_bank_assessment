package dao

import (
	"time"
)


type Journal struct{
	ID          int       `gorm:"primaryKey"`
	TanggalTransaksi     time.Time 
	NoRekeningKredit          string    `gorm:"size:255;"`
	NoRekeningDebit          string    `gorm:"size:255;"`
	NominalKredit      float64     `gorm:"type:numeric(10,2);not null"`
	NominalDebit      float64     `gorm:"type:numeric(10,2);not null"`
}


func (Journal) TableName() string {
    return "journal"
}


