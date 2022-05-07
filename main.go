package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mvrsss/bank-api/config"
	"github.com/mvrsss/bank-api/controllers"
	"github.com/mvrsss/bank-api/database"
	"github.com/mvrsss/bank-api/middlewares"
)

func main() {
	configs, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	database.Init(configs.DBUrl)
	router := initRouter()
	router.Run(configs.Port)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/user")
	{
		api.POST("/register", controllers.RegisterUser)
		api.GET("/login", controllers.GetToken)
		authorized := api.Group("/authorized").Use(middlewares.Auth())
		{
			authorized.POST("/loanrequest", controllers.NewLoanRequest)
			authorized.POST("/updatebalance", controllers.UpdateBalance)
			authorized.GET("/getuserloan", controllers.GetUserLoan)
		}
	}
	return router
}
