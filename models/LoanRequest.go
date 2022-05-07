package models

import (
	"time"

	"gorm.io/gorm"
)

type LoanRequest struct {
	gorm.Model
	UID          string    `gorm:"column:uid type:text"`
	Amount       float64   `gorm:"column:amount" json:"amount"`
	StartDate    time.Time `gorm:"column:startdate" json:"startdate"`
	EndDate      time.Time `gorm:"column:enddate" json:"enddate"`
	InterestRate float64   `gorm:"column:interestrate" json:"interestrate"`
	ClientID     string    `gorm:"type:text" json:"clientid"`
}
