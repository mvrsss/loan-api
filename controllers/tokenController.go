package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvrsss/bank-api/models"
)

type TokenRequest struct {
	PhoneNumber string `json:"phonenumber"`
	Password    string `json:"password"`
}

func GetToken(ctx *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// check if email exists and password is correct
	user = ValidateUser(ctx, request.PhoneNumber, request.Password)
	ctx.JSON(http.StatusOK, gin.H{"token": user.Token, "uid": user.UID})
}
