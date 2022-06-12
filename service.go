package main

import "github.com/bwmarrin/discordgo"

type Service func(s *discordgo.Session, m *Message)
