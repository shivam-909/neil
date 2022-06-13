package lunobot

import (
	"context"
	"flag"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
	"github.com/shivam-909/lunobot/ex/luno"
	"gopkg.in/yaml.v2"
)

var (
	cfg = flag.String("config file", "./config.yaml", "location of configuration file")
)

type config struct {
	DiscordAccessToken string
	Luno               luno.Config
}

type bot struct {
	Id      string
	Config  *config
	Session *discordgo.Session
}

func NewBot() *bot {

	bot := &bot{}

	bot.LoadConfig()

	session, err := discordgo.New(bot.Config.DiscordAccessToken)
	if err != nil {
		panic(err)
	}

	bot.Session = session

	user, err := bot.Session.User("@me")
	if err != nil {
		panic(err)
	}

	bot.Id = user.ID

	luno.Init(&bot.Config.Luno)

	return bot
}

func (b *bot) Start() {
	b.Session.AddHandler(b.MessageHandler)
	err := b.Session.Open()
	if err != nil {
		panic(err)
	}
}

func (b *bot) LoadConfig() {
	file, err := ioutil.ReadFile(*cfg)
	if err != nil {
		panic(err)
	}

	c := &config{}

	if err := yaml.Unmarshal(file, c); err != nil {
		panic(err)
	}

	b.Config = c
}
func (b *bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	ctx := context.Background()

	if m.Author.ID == b.Id {
		return
	}

	msg := parseMessage(m)
	if msg == nil {
		return
	}

	svc := supportedCommands[msg.Command]
	_ = svc(ctx, b, msg)
}
