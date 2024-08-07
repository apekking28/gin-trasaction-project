package models

import "gorm.io/gorm"

type Account struct {
    gorm.Model
    UserID  uint    `json:"user_id"`
    Type    string  `json:"type"`
    Balance float64 `json:"balance"`
    Transactions []Transaction
}
