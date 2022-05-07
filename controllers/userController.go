package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"github.com/mvrsss/bank-api/auth"
	"github.com/mvrsss/bank-api/database"
	"github.com/mvrsss/bank-api/models"
)

func ValidateUser(ctx *gin.Context, phoneNumber, password string) models.User {
	var user models.User
	record := database.DBInstance.Where("phonenumber = ?", phoneNumber).First(&user)
	if record.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, record.Error)
		// return
	}
	credentialError := user.CheckPassword(password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		// return
	}
	err := auth.ValidateToken(user.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		// return
	}
	return user
}

func RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//record := database.DBInstance.Where("phonenumber = ?", user.PhoneNumber).Where("name = ?", user.Name).Where("surname = ?", user.Surname).Where("iin = ?", user.IIN).Where("password = ?", user.Password).Where("address = ?", user.Address).First(&user)
	record := database.DBInstance.Where(&models.User{IIN: user.IIN, PhoneNumber: user.PhoneNumber}).First(&user)
	if record.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{"token": user.Token})
		return
	}
	if err := auth.ValidateToken(user.Token); err != nil {
		if err := user.HashPassword(user.Password); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		user.UID = shortuuid.New()
		record := database.DBInstance.Create(&user)
		if record.Error != nil {
			ctx.AbortWithError(http.StatusInternalServerError, record.Error)
			return
		}
		tokenString, err := auth.GenerateJWT(user.Name, user.Surname, user.IIN, user.PhoneNumber)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		user.Token = tokenString
		database.DBInstance.Save(&user)
	}
	ctx.JSON(http.StatusOK, gin.H{"token": user.Token})
}

type BalanceUpdateRequest struct {
	PhoneNumber string  `json:"phonenumber"`
	Password    string  `json:"password"`
	Amount      float64 `json:"amount"`
}

func UpdateBalance(ctx *gin.Context) {
	var request BalanceUpdateRequest
	var user models.User
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user = ValidateUser(ctx, request.PhoneNumber, request.Password)
	user.Balance += request.Amount
	database.DBInstance.Save(&user)
	ctx.JSON(http.StatusOK, gin.H{"message": "user balance was successfully updated", "uid": user.UID, "balance": user.Balance})
}
