package tui

import (
	"github.com/PhosFactum/expense-tracker/internal/usecase"
	"github.com/PhosFactum/expense-tracker/internal/model"
	"github.com/PhosFactum/expense-tracker/pkg/input"
	"fmt"
)

// TUIHandler: структура для хэндлера для CLI
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
func (h *CLIHandler) showExpensesHandler() {
	fmt.Println("--- List of all expenses ---")

	expenses, err := h.usecase.ShowExpensesUsecase()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return

	for i, e := range expenses {
		fmt.Printf("%d. %s - %.2f\n", i, e.Description, e.Amount) 
	}
}

// addExpenseHandler: ручка для создания новой траты
func (h *TUIHandler) addExpenseHandler() {
	fmt.Println("--- Adding new expense ---")

	fmt.Print("Type amount: ")
	amount, err := input.GetFloat() 
	if err != nil {
		fmt.Printf("Error: %v", err)
	} 
	fmt.Print("Type description: ")
	description, err := input.GetString() 
	if err != nil {
		fmt.Printf("Error: %v", err)
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
	fmt.Println("New expense added!")
}

// showExpensesHandler: ручка для вывода всех трат
func (h *TUIHandler) updateExpenseHandler() {
	fmt.Println("--- Updating old expense by id ---")

	fmt.Print("ID of expense: ")
	id, err := input.GetFloat()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Print("Type new amount: ")
	newAmount, err := input.GetFloat() 
	if err != nil {
		fmt.Printf("Error: %v", err)
	} 
	fmt.Print("Type new description: ")
	newDesc, err := input.GetString() 
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	updExpense := model.Expense{
		Amount: newAmount,
		Description: newDesc,
	}
	
	err = h.usecase.EditExpenseUsecase(id, updExpense)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

// showExpensesHandler: ручка для вывода всех трат
func (h *TUIHandler) deleteExpenseHandler() {
	fmt.Println("--- Deleting old expense by id ---")
	
	fmt.Print("ID of expense: ")
	id, err := input.GetFloat()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	err = h.usecase.DeleteExpenseUsecase(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
