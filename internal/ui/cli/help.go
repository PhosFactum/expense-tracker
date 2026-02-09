package cli

import (
	"fmt"
)

// printHelp: руководство по CLI-командам
func printHelp() {
	fmt.Println("=== Expense Tracker CLI ===")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("	list")
	fmt.Println("	add <amount> <description>   ---  show all expenses")
	fmt.Println("	update <id> <amount> <desc>  ---  add expense")
	fmt.Println("	delete <id>                  ---  delete expense")
	fmt.Println("   help                         ---  show help")
	fmt.Println("   exit                         ---  quit")
}
