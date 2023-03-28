package site

import (
	"encoding/json"
	"github.com/mr-tron/base58"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/fairytale5571/punkz/internal/bots/ds"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/markbates/goth/gothic"
)

type DBProvider interface {
	InsertWallet(w WalletDatabase) error

	GetWallet(userID string) (WalletDatabase, error)
}

type Provider interface {
	HasAuth(c *gin.Context)
	OnDiscordPresent(c *gin.Context)

	Auth(c *gin.Context)
	GetUser(c *gin.Context)
	AuthCallback(c *gin.Context)

	CreateWallet(c *gin.Context)
}

type site struct {
	db DBProvider
	ds ds.Provider

	mux *sync.Mutex
}

func New(db DBProvider, discord ds.Provider) Provider {
	return &site{
		db:  db,
		ds:  discord,
		mux: &sync.Mutex{},
	}
}

type WalletRequest struct {
	Wallet string `json:"wallet"`
}

type WalletDatabase struct {
	Wallet   string `json:"name" bson:"name"`
	UserID   string `json:"user_id" bson:"user_id"`
	UserName string `json:"user_name" bson:"user_name"`
	Email    string `json:"email" bson:"email"`
}

func (s *site) CreateWallet(c *gin.Context) {
	s.mux.Lock()
	defer s.mux.Unlock()

	claims, _ := s.jwtClaims(c)

	var wr WalletRequest
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := json.Unmarshal(body, &wr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// user already has a wallet
	if wal, err := s.db.GetWallet(claims["id"].(string)); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "you already entered a wallet address",
			"wallet": wal.Wallet,
		})
		return
	}

	ok := isValidSolanaAddress(wr.Wallet)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid solana wallet address",
		})
		return
	}

	s.db.InsertWallet(WalletDatabase{
		Wallet:   wr.Wallet,
		UserID:   claims["id"].(string),
		UserName: claims["name"].(string),
		Email:    claims["email"].(string),
	})
	c.JSON(http.StatusOK, gin.H{
		"wallet": wr.Wallet,
	})
	return
}

func (s *site) Auth(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (s *site) GetUser(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *site) AuthCallback(c *gin.Context) {
	c.AddParam("provider", "discord")
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "could not complete user auth",
			"message": err.Error(),
		})
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.UserID
	claims["name"] = user.Name
	claims["avatar"] = user.AvatarURL
	claims["email"] = user.Email
	claims["access_token"] = user.AccessToken
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not sign token",
		})
		return
	}
	c.Set("user", user)
	if err := s.ds.AddMemberGuild(os.Getenv("DISCORD_GUILD"), user.UserID, user.AccessToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not add member to discord guild",
		})
		return
	}
	c.SetCookie("jwt", tokenString, int(128*time.Hour), "/", "", true, true)
	c.Redirect(http.StatusFound, os.Getenv("SITE_URL")+"?jwt="+tokenString)
}

func isValidSolanaAddress(address string) bool {
	if len(address) != 44 {
		return false
	}

	_, err := base58.Decode(address)
	if err != nil {
		return false
	}
	return true
}
