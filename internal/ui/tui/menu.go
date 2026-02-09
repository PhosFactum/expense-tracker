package tui

import (
	"github.com/PhosFactum/expense-tracker/pkg/input"
	"fmt"
)

// RunTUI: функция для запуска TUI
func (h *TUIHandler) Run() {
	fmt.Println("=== Expense Tracker ===\n")

	for {
		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Println("| 0. Exit                |")
		fmt.Println("| 1. Show expenses       |")
		fmt.Println("| 2. Add expense         |")
		fmt.Println("| 3. Update expense      |")
		fmt.Println("| 4. Delete expense      |")
		fmt.Println("| 5. Summary by all time |")
		fmt.Println("| 6. Summary by month    |")
		fmt.Println("| 7. Export data to CSV  |")
		fmt.Println("--------------------------")
		fmt.Println("|")
		  fmt.Print("|> Your choice: ")
		choice, err := input.GetInt()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if choice < 0 || choice > 7 {
			fmt.Println("Invalid choice, try again!")
			continue
		}

		switch(choice) {
			case 0:
				fmt.Println("--- Program was terminated ---")
				return
			case 1:
				h.showExpensesHandler()
			case 2:
				h.addExpenseHandler()
			case 3:
				h.updateExpenseHandler()
			case 4:
				h.deleteExpenseHandler()
			case 5:
				fmt.Println("Summary by all time: will be soon...")
			case 6:
				fmt.Println("Summary by month: will be soon...")
			case 7:
				fmt.Println("Exporting to CSV: will be soon...")
			default:
				fmt.Println("Invalid choice, try again!")
				continue
		}
	}
}
