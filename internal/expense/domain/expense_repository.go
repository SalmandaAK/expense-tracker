package domain

type ExpenseRepository interface {
	AddExpense(*Expense) error
	FindAllExpenses() ([]*Expense, error)
	FindExpenseById(ExpenseId) (*Expense, error)
	DeleteExpense(*Expense) error
	FindAllExpensesByMonth(int) ([]*Expense, error)
}
