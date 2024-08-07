package main

import (
    "github.com/gin-gonic/gin"
    "github.com/apekking28/gin_project/payment_manager/controllers"
    "github.com/apekking28/gin_project/payment_manager/database"
)

func main() {
    r := gin.Default()
    database.ConnectDB()

    r.POST("/send", controllers.Send)
    r.POST("/withdraw", controllers.Withdraw)
    r.GET("/transactions", controllers.GetTransactions)

    r.Run(":8081")
}
