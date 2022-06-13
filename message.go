package lunobot

import (
	"github.com/bwmarrin/discordgo"
)

type Message struct {
	Command  Command
	Trailing []string
}

func parseMessage(m *discordgo.MessageCreate) *Message {
	msg := []rune(m.Message.Content)
	if string(msg[0]) != "$" {
		return nil
	}

	if len(msg) == 1 {
		return nil
	}

	inputs := leaderToCommand(msg)

	command, err := validCommand(inputs[0])
	if err != nil {
		return nil
	}
	return &Message{
		Command:  command,
		Trailing: inputs[1:],
	}
}

func leaderToCommand(bits []rune) []string {
	ret := make([]string, 0)
	stack := make([]rune, 0)
	for i := 1; i < len(bits); i++ {

		if i+1 == len(bits) {
			stack = append(stack, bits[i])
			ret = append(ret, string(stack))
			continue
		}

		if string(bits[i]) == " " {
			ret = append(ret, string(stack))
			stack = []rune{}
			continue
		}
		stack = append(stack, bits[i])
	}
	return ret
}
