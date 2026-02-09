package main

import (
	"github.com/PhosFactum/expense-tracker/internal/app"
)

func main() {
	app := app.NewApp()   // Создали экземпляр приложения (со всеми зависимостями)
	app.CLIHandler.Run()  // Запустили метод хэндлера Run (CLI)
}
