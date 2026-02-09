package cli

import (
	"github.com/PhosFactum/expense-tracker/internal/usecase"
	"github.com/PhosFactum/expense-tracker/internal/model"
	"fmt"
)

// CLIHandler: структура для хэндлера для CLI
type CLIHandler struct {
	usecase *usecase.ExpenseUsecase
}

// NewCLIHandler: конструктор для сборки хэндлера
func NewCLIHandler(uc *usecase.ExpenseUsecase) *CLIHandler {
	return &CLIHandler {
		usecase: uc,
	}
}

// showExpensesHandler: ручка для вывода всех трат
func (h *CLIHandler) List() {
	expenses, err := h.usecase.ShowExpensesUsecase()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}

	for _, e := range expenses {
		fmt.Printf("%d. | %.2f | - %s\n", e.ID, e.Amount, e.Description) 
	}
}

// addExpenseHandler: ручка для создания новой траты
func (h *CLIHandler) Add(amount float64, desc string) {
	expense := model.Expense{
		Amount: amount,
		Description: desc,
	}

	err := h.usecase.AddExpenseUsecase(expense)
	if err != nil {
		fmt.Println("Error:", err)	
		return
	}

	fmt.Println("Expense added!")
}

// showExpensesHandler: ручка для вывода всех трат
func (h *CLIHandler) Update(id int, amount float64, desc string) {
	expense := model.Expense{
		Amount: amount,
		Description: desc,
	}
	
	err := h.usecase.EditExpenseUsecase(id, expense)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Expense updated!")
}

// showExpensesHandler: ручка для вывода всех трат
func (h *CLIHandler) Delete(id int) {
	err := h.usecase.DeleteExpenseUsecase(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Expense deleted!")
}
