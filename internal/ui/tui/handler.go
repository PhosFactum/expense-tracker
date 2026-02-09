package tui

import (
	"github.com/PhosFactum/expense-tracker/internal/usecase"
	"github.com/PhosFactum/expense-tracker/internal/model"
	"github.com/PhosFactum/expense-tracker/pkg/input"
	"fmt"
)

// TUIHandler: структура для хэндлера для TUI
type TUIHandler struct {
	usecase *usecase.ExpenseUsecase
}

// NewTUIHandler: конструктор для сборки хэндлера
func NewTUIHandler(uc *usecase.ExpenseUsecase) *TUIHandler {
	return &TUIHandler {
		usecase: uc,
	}
}

// showExpensesHandler: ручка для вывода всех трат
func (h *TUIHandler) showExpensesHandler() {
	fmt.Println("--- List of all expenses ---")

	expenses, err := h.usecase.ShowExpensesUsecase()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}
	for i, e := range expenses {
		fmt.Printf("%d. %s - %.2f\n", i+1, e.Description, e.Amount) 
	}
}

// addExpenseHandler: ручка для создания новой траты
func (h *TUIHandler) addExpenseHandler() {
	fmt.Println("--- Adding new expense ---")

	fmt.Print("Type amount: ")
	amount, err := input.GetFloat() 
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	} 
	fmt.Print("Type description: ")
	description, err := input.GetString() 
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	expense := model.Expense{
		Amount: amount,
		Description: description,
	}

	err = h.usecase.AddExpenseUsecase(expense)
	if err != nil {
		fmt.Printf("Error: %v", err)	
		return
	}
	
	fmt.Println("Expense added successfully!")
}

// updateExpenseHandler: ручка для обновления траты
func (h *TUIHandler) updateExpenseHandler() {
	fmt.Println("--- Updating old expense by id ---")

	fmt.Print("ID of expense: ")
	id, err := input.GetInt()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	if !h.usecase.ExpenseExistsUsecase(id) {
		fmt.Println("Expense with this ID does not exist")
		return
	}

	fmt.Print("Type new amount: ")
	newAmount, err := input.GetFloat() 
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	} 
	fmt.Print("Type new description: ")
	newDesc, err := input.GetString() 
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	updExpense := model.Expense{
		Amount: newAmount,
		Description: newDesc,
	}
	
	err = h.usecase.EditExpenseUsecase(id, updExpense)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("Expense updated successfully!")
}

// showExpensesHandler: ручка для вывода всех трат
func (h *TUIHandler) deleteExpenseHandler() {
	fmt.Println("--- Deleting old expense by id ---")
	
	fmt.Print("ID of expense: ")
	id, err := input.GetInt()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	if !h.usecase.ExpenseExistsUsecase(id) {
		fmt.Println("Expense with this ID does not exist")
		return
	}

	err = h.usecase.DeleteExpenseUsecase(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("Expense deleted successfully!")
}
