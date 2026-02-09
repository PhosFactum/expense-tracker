package repository 

import (
	"github.com/PhosFactum/expense-tracker/internal/model"
)

// Repository: интерфейс с методами репозитория расходов
type Repository interface {
	GetAll() ([]model.Expense, error)
	Create(expense model.Expense) error
	Update(id int, updExpense model.Expense) error
	Delete(id int) error
	Exists(id int) bool
	// SummarizeAll()
	// SummarizeByMonth()
}
