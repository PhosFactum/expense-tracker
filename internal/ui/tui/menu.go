package tui

import (
	"github.com/PhosFactum/expense-tracker/pkg/input"
	"fmt"
)

// RunTUI: функция для запуска TUI
func RunTUI() {
	fmt.Println("=== Expense Tracker ===\n")

	var choice int
	for {
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

		choice, err := input.GetInt()
		if choice < 0 || choice > 7 {
			fmt.Println("Invalid choice, try again!")
			continue
		}

		switch(choice) {
			case 0:
				fmt.Println("--- Program was terminated ---")
				break
			case 1:
				showExpensesHandler()
			case 2:
				addExpenseHandler()
			case 3:
				updateExpenseHandler()
			case 4:
				deleteExpenseHandler()
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
