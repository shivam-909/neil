package lunobot

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot := NewBot()
	bot.Start()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	os.Exit(1)
}
