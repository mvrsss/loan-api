package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"github.com/mvrsss/bank-api/database"
	"github.com/mvrsss/bank-api/models"
)

func NewLoanRequest(ctx *gin.Context) {
	var loan models.LoanRequest
	if err := ctx.BindJSON(&loan); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	loan.UID = shortuuid.New()

	res := database.DBInstance.Table("loan_requests").Create(&loan)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "record saved"})
	return
}
