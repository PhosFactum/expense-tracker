package cli

import (
	"github.com/PhosFactum/expense-tracker/internal/usecase"
	"strings"
	"fmt"
	"os"
)

// RunCLI: функция для запуска CLI
func RunCLI(expenseUsecase *usecase.ExpenseUsecase) error {
	if len(os.Args) < 2 {
		printHelp()
		return nil
	}


}
