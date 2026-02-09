package app

import (
	"github.com/PhosFactum/expense-tracker/internal/repository"
	"github.com/PhosFactum/expense-tracker/internal/usecase"
	//"github.com/PhosFactum/expense-tracker/internal/ui/cli"
	"github.com/PhosFactum/expense-tracker/internal/ui/tui"
)

// App: структура приложения с инициализацией компонентов
type App struct {
	Repo *repository.ExpenseRepository
	UC *usecase.ExpenseUsecase
	//CLIHandler *cli.CLIHandler
	TUIHandler *tui.TUIHandler
}

// NewApp: конструктор приложения и сборщик компонентов
func NewApp() *App {
	repo := repository.NewExpenseRepository()
	uc := usecase.NewExpenseUsecase(repo)
	// Дробим хэндлеры на CLI-часть и TUI-часть
	//clihandler := cli.NewCLIHandler(uc)
	tuihandler := tui.NewTUIHandler(uc)

	return &App{
		Repo: repo,
		UC: uc,
		//CLIHandler: clihandler,
		TUIHandler: tuihandler,
	}
}
