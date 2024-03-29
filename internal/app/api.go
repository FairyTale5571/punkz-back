package app

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
)

func setupGoth() {
	goth.UseProviders(
		discord.New(
			os.Getenv("DISCORD_CLIENT"),
			os.Getenv("DISCORD_SECRET"),
			os.Getenv("DISCORD_CALLBACK"),
			discord.ScopeIdentify,
			discord.ScopeGuilds,
			discord.ScopeEmail,
			discord.ScopeJoinGuild,
			discord.ScopeConnections,
		),
	)
}

func (app *App) router() {

	setupGoth()

	app.Server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}),
		func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
		},
	)

	api := app.Server.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.GET("/auth", app.Site.Auth)
	api.GET("/auth/discord/callback", app.Site.AuthCallback)
	api.GET("/user", app.Site.HasAuth, app.Site.GetUser)
	api.POST("/wallet", app.Site.HasAuth, app.Site.OnDiscordPresent, app.Site.CreateWallet)
}
