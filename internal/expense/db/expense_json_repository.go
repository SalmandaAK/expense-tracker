package db

import (
	"cmp"
	"encoding/json"
	"errors"
	"io/fs"
	"maps"
	"os"
	"slices"
	"time"

	"github.com/SalmandaAK/expense-tracker/internal/expense/domain"
	"github.com/SalmandaAK/expense-tracker/internal/helper"
)

var (
	errEmptyExpenseList        = errors.New("expense list is empty")
	errTaskNotFound            = errors.New("task not found")
	errEmptyExpenseListByMonth = errors.New("no expenses found during this month")
)

type ExpenseJSONRepository struct {
	filePath string
	expenses map[domain.ExpenseId]*domain.Expense
}

func New(filePath string) *ExpenseJSONRepository {
	return &ExpenseJSONRepository{
		filePath: filePath,
		expenses: make(map[domain.ExpenseId]*domain.Expense),
	}
}

func (r *ExpenseJSONRepository) loadAllExpenses() error {
	jsonBlob, err := os.ReadFile(r.filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return errEmptyExpenseList
		} else {
			return err
		}
	}
	return json.Unmarshal(jsonBlob, &r.expenses)
}

func (r *ExpenseJSONRepository) saveAllExpenses() error {
	jsonBlob, err := json.MarshalIndent(&r.expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, jsonBlob, 0666)
}

func (r *ExpenseJSONRepository) AddExpense(e *domain.Expense) error {
	err := r.loadAllExpenses()
	if err != nil {
		if err != errEmptyExpenseList {
			return err
		}
	}
	id := domain.ExpenseId(helper.GenerateNumberId(r.expenses))
	e.Id = id
	r.expenses[e.Id] = e
	return r.saveAllExpenses()
}

func (r *ExpenseJSONRepository) FindAllExpenses() ([]*domain.Expense, error) {
	err := r.loadAllExpenses()
	if err != nil {
		return nil, err
	}
	if len(r.expenses) == 0 {
		return nil, errEmptyExpenseList
	}
	var expenses []*domain.Expense
	if len(r.expenses) > 0 {
		expenses = slices.SortedFunc(maps.Values(r.expenses), func(e1, e2 *domain.Expense) int {
			return cmp.Compare(e1.Id, e2.Id)
		})
	}
	return expenses, nil
}

func (r *ExpenseJSONRepository) FindExpenseById(id domain.ExpenseId) (*domain.Expense, error) {
	err := r.loadAllExpenses()
	if err != nil {
		return nil, err
	}
	e, ok := r.expenses[id]
	if !ok {
		return nil, errTaskNotFound
	}
	return e, nil
}

func (r *ExpenseJSONRepository) DeleteExpense(e *domain.Expense) error {
	delete(r.expenses, e.Id)
	return r.saveAllExpenses()
}

func (r *ExpenseJSONRepository) FindAllExpensesByMonth(month int) ([]*domain.Expense, error) {
	err := r.loadAllExpenses()
	if err != nil {
		return nil, err
	}
	if len(r.expenses) == 0 {
		return nil, errEmptyExpenseList
	}
	var expensesByMonth []*domain.Expense
	iter := maps.Values(r.expenses)
	for expense := range iter {
		if expense.CreatedAt.Month() == time.Month(month) {
			expensesByMonth = append(expensesByMonth, expense)
		}
	}
	if len(expensesByMonth) == 0 {
		return nil, errEmptyExpenseListByMonth
	}
	return expensesByMonth, nil
}
