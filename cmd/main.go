package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/Ksld154/niji-discord-bot/pkg/bitly"
	"github.com/Ksld154/niji-discord-bot/pkg/helpmsg"
	"github.com/Ksld154/niji-discord-bot/pkg/nijionair"
	"github.com/Ksld154/niji-discord-bot/pkg/nijiparser"
	"github.com/Ksld154/niji-discord-bot/pkg/utils"
	"github.com/Ksld154/niji-discord-bot/pkg/ytpicker"

	"github.com/bwmarrin/discordgo"
)

const (
	bitlyEndPoint = "https://api-ssl.bitly.com/v4/shorten"
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

	utils.BotStartTime = time.Now()
	bitly.BitlyToken = os.Getenv("BITLY_TOKEN")

	dg.UpdateStatus(0, "$maru")

	fmt.Println("Bot is running.")
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		<-sc
		dg.Close()
	}()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	m.Content = strings.TrimSpace(m.Content)
	messageArgs := strings.Split(m.Content, " ")

	m.Content = strings.ToLower(m.Content)
	// fmt.Println(m.Content)
	// fmt.Println(messageArgs)
	// fmt.Println(len(m.Mentions))

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
		localIPAddr := utils.GetOutBoundIPAddr("https://ifconfig.me")
		s.ChannelMessageSend(m.ChannelID, localIPAddr)
	} else if m.Content == "$demo" {
		s.ChannelMessageSend(m.ChannelID, ":white_check_mark: "+"<@"+m.Author.ID+">")
	} else if m.Content == "$sui" {
		s.ChannelMessageSend(m.ChannelID, ytpicker.GetRandomSong("sui", "PLarzPAXT9RrpT2W_KUTOzNfaYZ4Qs9_-D"))
	} else if m.Content == "$watame" {
		s.ChannelMessageSend(m.ChannelID, ytpicker.GetRandomSong("watame", "PLZ34fLWik_iB_cdmHivl8xhMJW6JwIkNn"))
	} else if m.Content == "$inui" {
		s.ChannelMessageSend(m.ChannelID, ytpicker.GetRandomSong("inui", "PLp93VJ2iFLMJiHM_0FXfjljA8oyu_lXuS"))
	} else if m.Content == "$gif" {
		s.ChannelMessageSend(m.ChannelID, "Please upgrade to monthly plan <:Arisu:735409267133382659> \nhttps://www.youtube.com/channel/UCdpUojq0KWZCN9bxXnZwz5w/join")
	} else if m.Content == "$onair" || m.Content == "$o" {
		onair := nijionair.GetNijiOnAir()
		s.ChannelMessageSendEmbed(m.ChannelID, &onair)
	} else if m.Content == "$uptime" {
		s.ChannelMessageSend(m.ChannelID, utils.GetUpTime(utils.BotStartTime))
	} else if m.Content == "$maru" || m.Content == "$marumaru" {
		s.ChannelMessageSend(m.ChannelID, "<:ars2:736167952348479508>")
		help := helpmsg.BuildHelpMsg()
		s.ChannelMessageSendEmbed(m.ChannelID, &help)
	} else if messageArgs[0] == "$kick" && len(m.Mentions) == 1 {
		s.ChannelMessageSend(m.ChannelID, "https://imgur.com/q90jFeM")
		s.ChannelMessageSend(m.ChannelID, ":white_check_mark: "+"<@!"+m.Mentions[0].ID+">"+" is kicked ")
	} else if messageArgs[0] == "$avatar" && len(m.Mentions) == 1 {
		s.ChannelMessageSend(m.ChannelID, m.Mentions[0].AvatarURL("512"))
	} else if messageArgs[0] == "$short" && len(messageArgs) == 2 {

		// client := &http.Client{}
		shortURL, err := bitly.GetShortURL(messageArgs[1], bitlyEndPoint)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "<:LizeCry:734715144323727451>")
		}
		s.ChannelMessageSend(m.ChannelID, shortURL)
	} else if messageArgs[0] == "$activity" && len(messageArgs) == 2 {
		s.UpdateStatus(0, messageArgs[1])
		s.ChannelMessageSend(m.ChannelID, ":white_check_mark: ")
	} else if ok, _ := regexp.MatchString("^\\$.+", m.Content); ok {
		s.ChannelMessageSend(m.ChannelID, "<:LizeCry:734715144323727451>")
	}
}
