package config

import (
	"path/filepath"

	"github.com/SalmandaAK/expense-tracker/internal/expense/db"
	"github.com/SalmandaAK/expense-tracker/internal/expense/service"
)

var FilePath = filepath.Join(filepath.Dir(""), "expense.json")

func InitiateExpenseErvice(filePath string) *service.ExpenseService {
	repo := db.New(filePath)
	service := service.New(repo)
	return service
}
