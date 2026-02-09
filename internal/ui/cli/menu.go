package cli

import (
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"os"
)

// RunCLI: функция для запуска CLI
func (h *CLIHandler) Run() {
	fmt.Println("=== Expense Tracker CLI ===")
	fmt.Println("Type 'help' to see commands, 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		
		args := strings.Split(line, " ")

		switch args[0] {

			case "exit":
				fmt.Println("--- Program was completed ---")
				return

			case "help":
				printHelp()

			case "list":
				h.List()

			case "add":
				if len(args) < 3 {
					fmt.Println("Usage: add <amount> <description>")
					return
				}
				amount, err := strconv.ParseFloat(args[1], 64)
				if err != nil {
					fmt.Println("Invalid amount!")
					continue
				}
				desc := strings.Join(args[2:], " ")
				h.Add(amount, desc)

			case "update":
				if len(args) < 4 {
					fmt.Println("Usage: update <id> <amount> <description>")
					return
				}
				id, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("Invalid id!")
					continue
				}
				amount, err := strconv.ParseFloat(args[2], 64)
				if err != nil {
					fmt.Println("Invalid amount!")
					continue
				}
				desc := strings.Join(args[3:], " ")
				h.Update(id, amount, desc)

			case "delete":
				if len(args) < 2 {
					fmt.Println("Usage: delete <id>")
					return
				}
				id, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("Invalid id!")
					continue
				}
				h.Delete(id)

			default:
				fmt.Println("Unknown command. Type 'help'!")
		}
	}
}
