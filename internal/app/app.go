package app

import (
	"github.com/PhosFactum/expense-tracker/internal/repository"
	"github.com/PhosFactum/expense-tracker/internal/usecase"
	"github.com/PhosFactum/expense-tracker/internal/ui/cli"
	"github.com/PhosFactum/expense-tracker/internal/ui/tui"
)

// App: структура приложения с инициализацией компонентов
type App struct {
	ExpenseRepository *repository.ExpenseRepository
	ExpenseUsecase *usecase.ExpenseUsecase
	CLIHandler cli.handler
	TUIHandler tui.handler
}

// NewApp: конструктор приложения и сборщик компонентов
func NewApp() *App {
	repo := repository.NewExpenseRepository()
	uc := usecase.NewExpenseUsecase(repo)
	// Дробим хэндлеры на CLI-часть и TUI-часть
	cli_handler := cli.NewExpenseHandler(expenseUsecase)
	tui_handler := tui.NewExpenseHandler(expenseUsecase)

	return &App{
		ExpenseRepository: repo,
		ExpenseUsecase: uc,
		CLIHandler: cli_handler,
		TUIHandler: tui_handler,
	}
}
