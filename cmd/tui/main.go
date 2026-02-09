package main

import (
	"github.com/PhosFactum/expense-tracker/internal/app"
)

func main() {
	app := app.NewApp()    // Создали экземпляр приложения с зависимостями
	app.TUIHandler.Run()   // Запустили метод хэндлера Run (TUI)
}
