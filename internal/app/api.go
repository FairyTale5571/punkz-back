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

	app.Server.Use(cors.Default(),
		func(c *gin.Context) {

			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
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
	api.GET("/auth/callback", app.Site.AuthCallback)
	api.GET("/user", app.Site.HasAuth, app.Site.GetUser)
	api.POST("/wallet", app.Site.HasAuth, app.Site.OnDiscordPresent, app.Site.CreateWallet)
	api.GET("/wallets", app.Site.GetWallets)
}
