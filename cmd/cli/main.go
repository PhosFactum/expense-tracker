package main

import (
	"github.com/PhosFactum/expense-tracker/internal/app"
)

func main() {
	app := app.NewApp()
	app.CLIHandler.Run()
}
