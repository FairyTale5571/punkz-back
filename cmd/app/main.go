package main

import (
	"log"
	"os"

	"github.com/fairytale5571/punkz/internal/app"
)

func main() {
	newApp, err := app.NewApp()
	if err != nil {
		log.Fatalf("start application failed: %v", err)
		return
	}
	newApp.Start()
	app.WaitForIntOrTerm()
	newApp.Stop()
	os.Exit(0)
}
