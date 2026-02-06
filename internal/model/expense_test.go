package model

import (
	"testing"
	"time"
)

func TestExpense(t *testing.T) {
	// Сидим тестовые данные
	expectedID		 := 1
	expectedDesc	 := "Покупка продуктов"
	expectedAmount	 := 2500.02
	expectedCategory := "продукты"

	expense := &Expense{
		ID: 		 expectedID,
		Description: expectedDesc,
		Amount:      expectedAmount, 
		Category:    expectedCategory,
		Date:		 time.Now(),
	}

	// Проверка 1: поля заполнены правильно
	if expense.ID != expectedID {
		t.Errorf("ID don't match: expected %d, receive %d.", expectedID, expense.ID)
	}
	if expense.Description != expectedDesc {
		t.Errorf("Descriptions don't match: expected %s, receive %s.", expectedDesc, expense.Description)
	}
	if expense.Amount != expectedAmount {
		t.Errorf("Amounts don't match: expected %.2f, receive %.2f.", expectedAmount, expense.Amount)
	}
	if expense.Category != expectedCategory {
		t.Errorf("Category don't match: expected %s, receive %s.", expectedCategory, expense.Category)
	}

	t.Log("OK: created structure is correct!")
}
