package helper

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Connection() *gorm.DB {
	db,err := gorm.Open(postgres.Open("user=postgres password=root host=localhost port=5432 dbname=sekolah sslmode=disable TimeZone=Asia/Jakarta"))
	PanicIfError(err)
	set,err := db.DB()
	set.SetConnMaxIdleTime(5 * time.Minute)
	set.SetMaxOpenConns(100)
	set.SetConnMaxLifetime(1 * time.Hour)
	set.SetMaxIdleConns(10)
	return db
}
