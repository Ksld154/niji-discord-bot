package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/Ksld154/niji-discord-bot/pkg/getip"
	"github.com/Ksld154/niji-discord-bot/pkg/helpmsg"
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
	} else if m.Content == "$ping" {
		s.ChannelMessageSend(m.ChannelID, "```Pong!```")
	} else if m.Content == "$help" {
		help := helpmsg.BuildHelpMsg()
		s.ChannelMessageSendEmbed(m.ChannelID, &help)
	} else if m.Content == "$niji" {
		schedule := nijiparser.NijiScheduleParser()
		s.ChannelMessageSendEmbed(m.ChannelID, &schedule)
	} else if m.Content == "$ip" {
		localIPAddr := getip.GetOutBOundIPAddr()
		s.ChannelMessageSend(m.ChannelID, localIPAddr)
	} else if m.Content == "$demo" {
		s.ChannelMessageSend(m.ChannelID, "🐰")
		s.ChannelMessageSend(m.ChannelID, ":rabbit:")
	} else if m.Content == "$gif" {
		s.ChannelMessageSend(m.ChannelID, "Please upgrade to monthly plan <:Arisu:735409267133382659> \nhttps://www.youtube.com/channel/UCdpUojq0KWZCN9bxXnZwz5w/join")
	} else if ok, _ := regexp.MatchString("^\\$.+", m.Content); ok {
		s.ChannelMessageSend(m.ChannelID, "<:LizeCry:734715144323727451>")
	}
}
