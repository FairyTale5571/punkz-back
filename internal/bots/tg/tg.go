package tg

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/fairytale5571/punkz/internal/site"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DBProvider interface {
	GetWallets() ([]site.WalletDatabase, error)
}

type Provider interface {
	Start()
	Stop()
}

type tg struct {
	bot  *tgbotapi.BotAPI
	db   DBProvider
	site Provider
}

func New(db DBProvider) Provider {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	return &tg{
		bot: bot,
		db:  db,
	}
}

func (tg *tg) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			tg.handleMessage(update.Message)
		}
	}
}

func (tg *tg) Stop() {
	tg.bot.StopReceivingUpdates()
}

func (tg *tg) handleMessage(msg *tgbotapi.Message) {

	switch msg.Text {
	case "/start":
		tg.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Welcome to Punkz!"))
	case "кошельки":
		tg.sendWallets(msg)
	}
}

func (tg *tg) sendWallets(msg *tgbotapi.Message) {
	wallets, err := tg.db.GetWallets()
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.Create("wallets.csv")
	if err != nil {
		tg.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Не получается создать файл"))
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	header := []string{"Кошелек", "ID Discord", "Discord", "Email Discord"}
	_ = writer.Write(header)

	for _, wallet := range wallets {
		_ = writer.Write([]string{wallet.Wallet, wallet.UserID, wallet.UserName, wallet.Email})
	}
	writer.Flush()

	file, err = os.Open("wallets.csv")
	if err != nil {
		tg.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Не получается открыть файл"))
		return
	}
	defer file.Close()

	tg.bot.Send(tgbotapi.NewDocument(msg.Chat.ID, tgbotapi.FileReader{
		Name:   "wallets.csv",
		Reader: file,
	}))
}
