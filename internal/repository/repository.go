package repository

import (
	"github.com/PhosFactum/expense-tracker/internal/model"
	"errors"
)

// ExpenseRepository: репозиторий, содержащий траты (само хранилище данных)
type ExpenseRepository struct {
	expenses []model.Expense	
}

// NewExpenseRepository: конструктор для создания репозитория трат
func NewExpenseRepository() *ExpenseRepository {
	return &ExpenseRepository{
		expenses: []model.Expense{},
	}
}

//----------------------------------------------------//
// === Методы хранилища данных (репозитория трат) === //
//----------------------------------------------------//
// Ошибки здесь обрабатываются не на уровни бизнес-логики, 
// а на уровне работы с данными (сущности, domain-level)

// GetAll: метод для вывода всех трат
func (r *ExpenseRepository) GetAll() ([]model.Expense, error) {
	return r.expenses, nil
}

// Create: метод для создания новой траты
func (r *ExpenseRepository) Create(expense model.Expense) error {
	// Проверка того, что трата не пустая
	if expense == (model.Expense{}) {
		return errors.New("нельзя создать пустую трату")
	}
	expense.ID = len(r.expenses) + 1
	
	r.expenses = append(r.expenses, expense)	
	return nil
}

// Update: метод для обновления старой траты по конкретным полям
func (r *ExpenseRepository) Update(id int, updExpense model.Expense) error {
	// Итерируемся по всем тратам и, найдя нужную, редактируем её, предварительно валидируя
	for i, e := range r.expenses {
		if e.ID == id {
			// Проверяем валидность нового описания и новой суммы траты
			if updExpense.Description != "" {
				r.expenses[i].Description = updExpense.Description
			}
			if updExpense.Amount != 0 {
				r.expenses[i].Amount = updExpense.Amount
			}
			return nil
		}
	}
	return errors.New("expense not found")
}

// Delete: метод для удаления траты по идентификатору
func (r *ExpenseRepository) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid index: must be positive")
	}

	// Удаляем запись через складывание двух срезов
	for i, e := range r.expenses {
		if e.ID == id {
			r.expenses = append(r.expenses[:i], r.expenses[i+1:]...)	
			return nil
		}
	}
	return errors.New("expense not found")
}
