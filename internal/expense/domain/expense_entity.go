package domain

import "time"

type Expense struct {
	Id          ExpenseId `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Amount      int       `json:"amount"`
}

type ExpenseId int
