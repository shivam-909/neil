package main

import (
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestParseMessages(t *testing.T) {
	scenarios := []struct {
		input    string
		expected *Message
	}{
		{
			input:    "$stats eth mktcap",
			expected: &Message{Command: CommandStatistics, Trailing: []string{"eth", "mktcap"}},
		},
	}

	for _, sc := range scenarios {
		res := parseMessage(&discordgo.MessageCreate{
			Message: &discordgo.Message{
				Content: sc.input,
			},
		})

		if res == nil {
			t.Fatalf("was not expecting nil")
			return
		}

		if !reflect.DeepEqual(res, sc.expected) {
			t.Fatalf("result differed from expected result \n expected %v \n got %v", sc.expected, res)
		}
	}
}
