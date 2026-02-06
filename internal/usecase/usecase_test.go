package usecase

import (
	"github.com/PhosFactum/expense-tracker/internal/repository"
	"github.com/PhosFactum/expense-tracker/internal/model"
	"testing"
	"fmt"
)

func TestCRUD(t *testing.T) {
	// Создадим тестовые репозиторий и юзкейс
	mockRepo := &repository.ExpenseRepository{}
	expenseUsecase := NewExpenseUsecase(mockRepo)

	// AddExpenseUsecase
	testExpense := model.Expense{
		Description: "Хлебушек",
		Amount: 200.1,
	}

	err := expenseUsecase.AddExpenseUsecase(testExpense)
	if err != nil {
		fmt.Println("error while creating")
	} else {
		fmt.Println("OK: (Create) expense for tests was created!")
	}

	// ShowExpensesUsecase
	expensesList, err := expenseUsecase.ShowExpensesUsecase()
	if err != nil {
		t.Errorf("error while showing the expenses list")
	} else {
		t.Log("OK: (ReadAll) expenses shown correctly: ", expensesList)
	}

	// UpdateExpenseUsecase
	updExpense := model.Expense{
		Description: "Бублик",
		Amount: 30.4,
	}

	err = expenseUsecase.EditExpenseUsecase(1, updExpense)
	if err != nil {
		t.Errorf("error while updating the expense")
	} else {
		t.Log("OK: (Update) expense updates successfully: ", expensesList)
	}

	// DeleteExpenseUsecase
	err = expenseUsecase.DeleteExpenseUsecase(1)
	if err != nil {
		t.Errorf("error while deleting the expense")
	} else {
		t.Log("OK: (Delete) expense deleted successfully: ", expensesList)
	}
}

