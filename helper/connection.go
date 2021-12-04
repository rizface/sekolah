package helper

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Connection() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)
	db,err := gorm.Open(postgres.Open("user=postgres password=root host=localhost port=5432 dbname=sekolah sslmode=disable TimeZone=Asia/Jakarta"),&gorm.Config{
		Logger: newLogger,
	})
	PanicIfError(err)
	set,err := db.DB()
	set.SetConnMaxIdleTime(5 * time.Minute)
	set.SetMaxOpenConns(100)
	set.SetConnMaxLifetime(1 * time.Hour)
	set.SetMaxIdleConns(10)
	return db
}
