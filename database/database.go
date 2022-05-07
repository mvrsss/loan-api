package database

import (
	"log"

	"github.com/mvrsss/bank-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB
var err error

func Init(url string) {
	DBInstance, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DBInstance.AutoMigrate(&models.User{}, &models.LoanRequest{})
	return
}
