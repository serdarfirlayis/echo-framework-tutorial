package storage

import (
	"log"

	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-hello-world/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error

	conString := config.GetPostgresConnectionString()
	log.Print(conString)

	dialector := postgres.Open(conString)

	DB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return DB
}
