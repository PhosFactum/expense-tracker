package main

import (
	"github.com/PhosFactum/expense-tracker/internal/app"
)

func main() {
	app := NewApp()
	app.TUIHandler.Run()
}
