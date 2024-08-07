package models

import "gorm.io/gorm"

type Transaction struct {
    gorm.Model
    AccountID uint    `json:"account_id"`
    Amount    float64 `json:"amount"`
    Status    string  `json:"status"`
    ToAddress string  `json:"to_address"`
}
