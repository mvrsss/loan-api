package controllers

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mvrsss/bank-api/database"
	"github.com/mvrsss/bank-api/models"
)

type GetLoanRequest struct {
	PhoneNumber string `json:"phonenumber"`
	Password    string `json:"password"`
	ClientID    string `json:"clientid"`
}

type NextPaymentDetails struct {
	StartDate    time.Time
	EndDate      time.Time
	Amount       float64
	InterestRate float64
}

func GetUserLoan(ctx *gin.Context) {
	var user models.User
	var userLoanRequests []models.LoanRequest
	var request GetLoanRequest
	var EarlyRepaymentAmount float64

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	user = ValidateUser(ctx, request.PhoneNumber, request.Password)

	res := database.DBInstance.Table("loan_requests").Where(&models.LoanRequest{ClientID: request.ClientID}).Find(&userLoanRequests)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, res.Error)
		ctx.Abort()
		return
	}

	now := time.Now()
	paymentsLeftNum := 0
	paymentsAmount := 0.0
	nextPayment := &NextPaymentDetails{time.Now(), time.Now(), 0.0, 0.0}

	diffYear, diffMonth, diffDays := math.MaxInt64, math.MaxInt64, math.MaxInt64

	for _, i := range userLoanRequests {
		fmt.Println(i.EndDate.Year(), i.EndDate.Month(), i.EndDate.Day())
		if i.EndDate.Year() > now.Year() || (i.EndDate.Year() == now.Year() && (i.EndDate.Month() > now.Month() || (i.EndDate.Month() == now.Month() && i.EndDate.Day() > now.Day()))) {
			paymentsLeftNum += 1
			if i.EndDate.Year()-now.Year() < diffYear {
				nextPayment = &NextPaymentDetails{i.StartDate, i.EndDate, i.Amount, i.InterestRate}
			} else if i.EndDate.Year()-now.Year() == diffYear {
				if int(i.EndDate.Month())-int(now.Month()) < diffMonth {
					nextPayment = &NextPaymentDetails{i.StartDate, i.EndDate, i.Amount, i.InterestRate}
				} else if int(i.EndDate.Month())-int(now.Month()) == diffMonth {
					if i.EndDate.Day()-now.Day() < diffDays {
						nextPayment = &NextPaymentDetails{i.StartDate, i.EndDate, i.Amount, i.InterestRate}
					}
				}
			}
			diffYear = nextPayment.EndDate.Year() - now.Year()
			diffMonth = int(nextPayment.EndDate.Month()) - int(now.Month())
			diffDays = nextPayment.EndDate.Day() - now.Day()
		} else {
			paymentsAmount += i.Amount
		}
	}

	database.DBInstance.Model(&user).Update("balance", user.Balance-paymentsAmount)
	repayTodayPeriod := float64((nextPayment.EndDate.Year()-now.Year())*365 + (int(nextPayment.EndDate.Month())-int(now.Month()))*30 + (nextPayment.EndDate.Day() - now.Day()))
	LoanPeriod := float64((nextPayment.EndDate.Year()-nextPayment.StartDate.Year())*365 + (int(nextPayment.EndDate.Month())-int(nextPayment.StartDate.Month()))*30 + (nextPayment.EndDate.Day() - nextPayment.StartDate.Day()))
	EarlyRepaymentAmount = repayTodayPeriod / (LoanPeriod * ((nextPayment.InterestRate) / 100))

	ctx.JSON(http.StatusOK, gin.H{"Current balance": user.Balance, "Payments Number": paymentsLeftNum, "Next Payment Amount": nextPayment.Amount, "Next Payment date": nextPayment.EndDate, "Early Repayment Amount": EarlyRepaymentAmount})
	return
}
