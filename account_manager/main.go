package main

import (
    "github.com/gin-gonic/gin"
    "github.com/apekking28/gin_project/account_manager/controllers"
    "github.com/apekking28/gin_project/account_manager/database"
    "github.com/apekking28/gin_project/account_manager/middlewares"
)

func main() {
    r := gin.Default()
    database.ConnectDB()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    authenticated := r.Group("/")
    authenticated.Use(middlewares.AuthMiddleware())
    {
        authenticated.GET("/accounts", controllers.GetAccounts)
        authenticated.GET("/transactions", controllers.GetTransactions)
        authenticated.POST("/accounts", controllers.CreateAccount)
    }

    r.Run(":8080")
}
