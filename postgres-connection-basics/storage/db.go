package storage

import (
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-connection-basics/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() *gorm.DB {
	var err error
	dsn := config.GetPostgresConnectionString()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect database: %v", err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	if DB == nil {
		log.Panic("Database connection is not initialized")
	}
	return DB
}
