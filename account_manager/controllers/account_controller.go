package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/apekking28/gin_project/account_manager/database"
    "github.com/apekking28/gin_project/account_manager/models"
    "github.com/apekking28/gin_project/account_manager/middlewares"
    "gorm.io/gorm"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var storedUser models.User
    if err := database.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil || storedUser.Password != user.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    token, err := middlewares.GenerateToken(storedUser.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}

func GetAccounts(c *gin.Context) {
    username, _ := c.Get("username")
    var user models.User
    if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var accounts []models.Account
    if err := database.DB.Where("user_id = ?", user.ID).Find(&accounts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, accounts)
}

func GetTransactions(c *gin.Context) {
    accountID := c.Query("account_id")

    if accountID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
        return
    }

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


func CreateAccount(c *gin.Context) {
    var account models.Account
    if err := c.ShouldBindJSON(&account); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate if user exists
    var user models.User
    if err := database.DB.Where("id = ?", account.UserID).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        }
        return
    }

    if err := database.DB.Create(&account).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "account creation successful", "account": account})
}