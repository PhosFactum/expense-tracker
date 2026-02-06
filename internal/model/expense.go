package model

import "time"

// Expense: представляет собой модель расхода
type Expense struct {
	ID			 int
	Description  string
	Amount       float64
	Category     string      // Позже: тип - структура Category
	Date 		 time.Time	
}


// IDEA: Возможно, понадобится модель Категории "Catergory"
