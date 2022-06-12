package main

import "fmt"

type Command string

const (
	CommandStatistics  Command = "stats"
	CommandNews        Command = "news"
	CommandDailyDigest Command = "daily"
	CommandChart       Command = "chart"
	CommandTrade       Command = "chart"
	CommandTransaction Command = "tx"
	CommandGas         Command = "gas"
)

var (
	supportedCommands = map[Command]bool{
		CommandStatistics: true,
	}
)

func validCommand(s string) (Command, error) {
	if _, ok := supportedCommands[Command(s)]; ok {
		return Command(s), nil
	}
	return "", fmt.Errorf("invalid command")
}
