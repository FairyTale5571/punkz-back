package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fairytale5571/punkz/internal/bots/ds"
	"github.com/fairytale5571/punkz/internal/bots/tg"
	"github.com/fairytale5571/punkz/internal/db"
	"github.com/fairytale5571/punkz/internal/site"
	"github.com/gin-gonic/gin"
)

type App struct {
	DS   ds.Provider
	TG   tg.Provider
	DB   db.Provider
	Site site.Provider

	Server *gin.Engine
}

func NewApp() (*App, error) {

	database, err := db.NewProvider(os.Getenv("MONGO_URL"))
	if err != nil {
		return nil, err
	}
	discord := ds.New()
	return &App{
		Server: gin.Default(),
		Site:   site.New(database, discord),
		DS:     discord,
		TG:     tg.New(database),
		DB:     database,
	}, nil
}

func (app *App) Start() error {
	app.router()
	go app.DS.Start()
	go app.TG.Start()

	return app.Server.Run(":" + os.Getenv("PORT"))
}

func (app *App) Stop() error {
	app.DS.Stop()
	app.TG.Stop()
	return nil
}

func WaitForIntOrTerm() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	log.Println("Received signal: ", sig)
}
