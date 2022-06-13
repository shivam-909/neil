package luno

import (
	lunosdk "github.com/luno/luno-go"
)

var (
	Client *lunosdk.Client
)

type Config struct {
	Id     string
	Secret string
}

func Init(config *Config) {
	client := lunosdk.NewClient()
	client.SetAuth(config.Id, config.Secret)
}
