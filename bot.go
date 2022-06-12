package main

import (
	"flag"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

var (
	cfg = flag.String("config file", "./config.yaml", "location of configuration file")
)

type config struct {
	AccessToken string
}

type bot struct {
	Id      string
	Config  *config
	Session *discordgo.Session
}

func NewBot() *bot {
	session, err := discordgo.New("")
	if err != nil {
		panic(err)
	}

	bot := &bot{
		Session: session,
	}

	bot.LoadConfig()

	user, err := bot.Session.User("@me")
	if err != nil {
		panic(err)
	}

	bot.Id = user.ID

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
	if m.Author.ID == b.Id {
		return
	}
	msg := parseMessage(m)
	switch msg.Command {
	default:
		return
	}
}
