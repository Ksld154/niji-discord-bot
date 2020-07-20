package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ksld154/niji-discord-bot/pkg/nijiparser"
	"github.com/bwmarrin/discordgo"
)

var botToken string = os.Getenv("DISCORD_BOT_TOKEN")

func main() {
	fmt.Println("Hello 2434!")

	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Bot is running. Press Ctrl-C to exit.")

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		<-sc
		dg.Close()
	}()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.SessionID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "```Pong!```")
	}
	if m.Content == "$niji" {
		schedule := nijiparser.NijiScheduleParser()
		s.ChannelMessageSendEmbed(m.ChannelID, &schedule)
	}

}
