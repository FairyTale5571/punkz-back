package tg

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Provider interface {
	Start()
	Stop()
}

type tg struct {
	bot  *tgbotapi.BotAPI
	site Provider
}

func New() Provider {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	return &tg{
		bot: bot,
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

}
