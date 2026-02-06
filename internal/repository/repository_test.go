package repository

import (
	"github.com/PhosFactum/expense-tracker/internal/model"
	"testing"
)

func TestCRUD(t *testing.T) {
	repo := NewExpenseRepository()

	emptyExpense := model.Expense{}
	testExpense := model.Expense{
		Amount: 20.2,
		Description: "Bread",
	}
	
	updExpense := model.Expense{
		Amount: 0.01,
		Description: "Trash",
	}

	// Тест - Create
	t.Logf("FAKE FAIL:")
	err := repo.Create(emptyExpense)
	if err != nil {
		t.Logf("OK: (Create) empty expense rejected: %v", emptyExpense)
	} else {
		t.Error("FAIL: expected error, got nil")
	}

	// Дважды создаём нормальные структуры (1, 2 - индексы)
	err = repo.Create(testExpense)
	if err != nil {
		t.Errorf("error while creating new expense: %v", err)
	} else {
		t.Logf("OK: (Create) created expense rightly!")
	}

	err = repo.Create(testExpense)
	if err != nil {
		t.Errorf("error while creating new expense: %v", err)
	} else {
		t.Logf("OK: (Create) created expense rightly!")
	}


	// Тест - GetAll
	expenses, err := repo.GetAll()
	if err != nil {
		t.Errorf("error while getting expenses: %v", err)
	} else {
		t.Logf("OK: (GetAll) readAll worked rightly, returned %d expenses", len(expenses))
	}
	t.Logf("Expenses (1, 2): %v", expenses)

	// Тест - Update
	t.Logf("FAKE FAIL:")
	err = repo.Update(0, updExpense)
	if err != nil {
		t.Errorf("error while updating expense: %v", err)
	} else {
		t.Log("OK: (Update 0) update worked rightly!")
	}

	err = repo.Update(1, updExpense)
	if err != nil {
		t.Errorf("error while updating expense: %v", err)
	} else {
		t.Log("OK: (Update 1) update worked rightly!")
	}

	// Тест - Delete (TODO)
	err = repo.Delete(1)
	if err != nil {
		t.Errorf("error while deleting expense: %v", err)
	} else {
		t.Log("OK: (Delete 1) delete worked rightly!")
	}

	t.Logf("FAKE FAIL:")
	err = repo.Delete(3)
	if err != nil {
		t.Errorf("error while deleting expense: %v", err)
	} else {
		t.Log("OK: (Delete 3) delete worked rightly!")
	}
}

