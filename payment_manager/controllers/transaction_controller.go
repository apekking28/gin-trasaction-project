package controllers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/apekking28/gin_project/payment_manager/database"
    "github.com/apekking28/gin_project/payment_manager/models"
    "strconv"
    "gorm.io/gorm"
)

func processTransaction(transaction models.Transaction) (models.Transaction, error) {
    time.Sleep(30 * time.Second) // Simulate long running process

    transaction.Status = "processed"
    if err := database.DB.Save(&transaction).Error; err != nil {
        return transaction, err
    }

    return transaction, nil
}

func Send(c *gin.Context) {
    var transaction models.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate if account exists
    var account models.Account
    if err := database.DB.Where("id = ?", transaction.AccountID).First(&account).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Account does not exist"})
        return
    }

    transaction.Status = "pending"
    if err := database.DB.Create(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    go processTransaction(transaction)

    c.JSON(http.StatusOK, gin.H{"message": "transaction processing started", "transaction": transaction})
}

func Withdraw(c *gin.Context) {
    var transaction models.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate if account exists
    var account models.Account
    if err := database.DB.Where("id = ?", transaction.AccountID).First(&account).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Account does not exist"})
        return
    }

    transaction.Status = "pending"
    transaction.Amount = -transaction.Amount
    if err := database.DB.Create(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    go processTransaction(transaction)

    c.JSON(http.StatusOK, gin.H{"message": "transaction processing started", "transaction": transaction})
}

func GetTransactions(c *gin.Context) {
    accountID := c.Query("account_id")

    // Validate if account_id is provided
    if accountID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
        return
    }

    // Validate if account_id is a valid number
    if _, err := strconv.Atoi(accountID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Account ID"})
        return
    }

    // Validate if account exists
    var account models.Account
    if err := database.DB.Where("id = ?", accountID).First(&account).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Account does not exist"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        }
        return
    }

    var transactions []models.Transaction
    if err := database.DB.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, transactions)
}
