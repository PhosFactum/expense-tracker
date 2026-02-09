package usecase

import (
	"github.com/PhosFactum/expense-tracker/internal/repository"
	"github.com/PhosFactum/expense-tracker/internal/model"
	"errors"
)

// ExpenseUsecase: структура для репозитория, ссылающаяся на репозиторий
type ExpenseUsecase struct {
	repo *repository.ExpenseRepository
}

// NewExpenseUsecase: конструктор для создания юзкейса с ссылкой на репозиторий
func NewExpenseUsecase(r *repository.ExpenseRepository) *ExpenseUsecase {
	return &ExpenseUsecase{
		repo: r,
	}
}

//-----------------------------------------//
// === Методы юзкейсов (бизнес-логики) === //
//-----------------------------------------//
// Здесь обрабатываются ошибки бизнес-правил, 
// а не формата данных или сущностей.

// ShowExpensesUsecase сценарий с выводом всех трат юзера
func (u *ExpenseUsecase) ShowExpensesUsecase() ([]model.Expense, error) {
	return u.repo.GetAll()
}

// AddExpenseUsecase: сценарий с добавлением новой траты
func (u *ExpenseUsecase) AddExpenseUsecase(expense model.Expense) error {
	if expense.Amount <= 0 {
		return errors.New("сумма траты не может быть меньше нуля")
	}
	if expense.Amount >= 100000 {
		return errors.New("сумма траты слишком большая")
	}

	return u.repo.Create(expense)
}

// EditExpenseUsecase: сценарий с изменением данных траты
func (u *ExpenseUsecase) EditExpenseUsecase(id int, updExpense model.Expense) error {
	if updExpense.Amount <= 0 {
		return errors.New("сумма траты не может быть меньше нуля")
	}
	if updExpense.Amount >= 100000 {
		return errors.New("сумма траты слишком большая")
	}

	return u.repo.Update(id, updExpense)
}

// DeleteExpenseUsecase: сценарий удаления траты
func (u *ExpenseUsecase) DeleteExpenseUsecase(id int) error {
	return u.repo.Delete(id)
}

// ExpenseExistsUsecase: проверка существования траты
func (u *ExpenseUsecase) ExpenseExistsUsecase(id int) bool {
	return u.repo.Exists(id)
}

