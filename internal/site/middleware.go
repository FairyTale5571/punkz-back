package site

import (
	"context"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

var jwtSecret = os.Getenv("SESSION_SECRET")

func (s *site) jwtClaims(c *gin.Context) (jwt.MapClaims, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

func (s *site) HasAuth(c *gin.Context) {

	claims, err := s.jwtClaims(c)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	c.Set("user", map[string]interface{}{
		"id":           claims["id"],
		"name":         claims["name"],
		"avatar":       claims["avatar"],
		"email":        claims["email"],
		"access_token": claims["access_token"],
		"exp":          claims["exp"],
	})
	c.Next()

	return
}

func (s *site) OnDiscordPresent(c *gin.Context) {

	user, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	guilds, err := FetchGuilds(user.(map[string]interface{})["access_token"].(string))
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
	for _, guild := range guilds {
		if guild.ID == os.Getenv("DISCORD_GUILD") {
			c.Next()
			return
		}
	}
	c.JSON(401, gin.H{
		"error": "need connect to discord",
	})
	c.Abort()
}

type guild struct {
	ID string `json:"id"`
}

func FetchGuilds(accessToken string) ([]guild, error) {
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	))

	resp, err := client.Get("https://discord.com/api/v10/users/@me/guilds")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var guilds []guild
	err = json.NewDecoder(resp.Body).Decode(&guilds)
	if err != nil {
		return nil, err
	}
	return guilds, nil
}
