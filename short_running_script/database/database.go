package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HasilQuery struct {
	Waktu   time.Time
	JenisTransaksi string
	Jumlah  float64
  }
  

func InitializeDB(user string, 
	dbname string, 
	password string, 
	host string,
	port string,) (*gorm.DB, error) {
	// Replace with your PostgreSQL connection details
	dsn := "user=" + user + " dbname=" + dbname + " password=" + password + " host=" + host + " port="+ port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// // AutoMigrate your models here
	// err = db.AutoMigrate(&Journal{})
	// if err != nil {
	// 	return nil, err
	// }
	
	return db, nil
}

func GetRekap(db *gorm.DB)(outputData []HasilQuery, err error){
	if err = db.Raw("select waktu::DATE, jenis_transaksi, sum(nominal) as jumlah from transaction group by 1,2").Scan(&outputData).Error; err != nil{
		return
	}
	return
}