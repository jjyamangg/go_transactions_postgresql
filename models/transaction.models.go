package models

import (
	"time"
)

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"unique;not null" json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

type Currency struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"unique;not null" json:"code"`
	Name string `json:"name"`
}

type TransactionType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

type Transaction struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	AccountID         uint            `gorm:"not null" json:"account_id"`
	Account           Account         `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	Amount            int             `gorm:"not null" json:"amount"`
	CurrencyID        uint            `gorm:"not null" json:"currency_id"`
	Currency          Currency        `gorm:"foreignKey:CurrencyID" json:"currency,omitempty"`
	TypeTransactionID uint            `gorm:"not null" json:"type_transaction_id"`
	TypeTransaction   TransactionType `gorm:"foreignKey:TypeTransactionID" json:"type_transaction,omitempty"`
	Description       string          `json:"description,omitempty"`
	CreatedAt         time.Time       `json:"created_at"`
}

type TransactionResponse struct {
	ID                uint      `json:"id"`
	AccountID         uint      `json:"account_id"`
	Amount            int       `json:"amount"`
	CurrencyID        uint      `json:"currency_id"`
	TypeTransactionID uint      `json:"type_transaction_id"`
	Description       string    `json:"description,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	RunningBalance    int       `json:"running_balance"`
}