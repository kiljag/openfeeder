package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDb *gorm.DB

func createPgPool() *gorm.DB {

	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=staging port=%s sslmode=disable", pgHost, pgUser, pgPass, pgPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDb.SetMaxIdleConns(2)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return db
}

func GetPGConn() *gorm.DB {
	if gormDb == nil {
		gormDb = createPgPool()
	}
	return gormDb
}
