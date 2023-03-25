package ds

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type Provider interface {
	Start()
	Stop()

	AddMemberGuild(guildID, userID, accessToken string) error
}

type ds struct {
	session *discordgo.Session
}

func New() Provider {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return nil
	}

	return &ds{
		session: s,
	}
}

func (d *ds) Start() {
	d.session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err := d.session.Open()
	if err != nil {
		return
	}
}

func (d *ds) Stop() {
	err := d.session.Close()
	if err != nil {
		return
	}
}

func (d *ds) AddMemberGuild(guildID, userID, accessToken string) error {
	return d.session.GuildMemberAdd(guildID, userID, &discordgo.GuildMemberAddParams{
		AccessToken: accessToken,
	})
}
