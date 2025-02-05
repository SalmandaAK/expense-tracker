package db

import (
	"cmp"
	"encoding/json"
	"errors"
	"io/fs"
	"maps"
	"os"
	"slices"

	"github.com/SalmandaAK/expense-tracker/internal/expense/domain"
	"github.com/SalmandaAK/expense-tracker/internal/helper"
)

var (
	errEmptyExpenseList = errors.New("expense list is empty")
	errTaskNotFound     = errors.New("task not found")
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
		return err
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
	expenses := slices.SortedFunc(maps.Values(r.expenses), func(e1, e2 *domain.Expense) int {
		return cmp.Compare(e1.Id, e2.Id)
	})
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
