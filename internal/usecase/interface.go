package usecase

import (
	"github.com/PhosFactum/expense-tracker/internal/model"	
)

// Usecase: интерфейс сценариев использования (слой бизнес-логики)
type Usecase interface {
	ShowExpensesUsecase() ([]model.Expense, error)
	AddExpenseUsecase(expense model.Expense) error
	EditExpenseUsecase(id int, updExpense model.Expense) error
	DeleteExpenseUsecase(id int) error
}
