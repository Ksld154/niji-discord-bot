package helpmsg

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func getScheduleThumbnail() string {
	thumbnailList := []string{
		"https://i.imgur.com/BRokk77.png",
		"https://i.imgur.com/uylS55f.png",
		"https://i.imgur.com/pbz7OXY.png",
	}

	idx := rand.Intn(len(thumbnailList))

	return thumbnailList[idx]
}

// BuildHelpMsg returns a embed of cmd description
func BuildHelpMsg() discordgo.MessageEmbed {
	var helpEmbed discordgo.MessageEmbed
	helpEmbed.Title = "$maru"
	helpEmbed.Color = 0xffdd00

	var embedMsgThumbnail discordgo.MessageEmbedThumbnail
	embedMsgThumbnail.URL = getScheduleThumbnail()
	helpEmbed.Thumbnail = &embedMsgThumbnail

	var helpAuthor discordgo.MessageEmbedAuthor
	helpAuthor.Name = "Marumaru"
	helpAuthor.URL = "https://www.youtube.com/channel/UCdpUojq0KWZCN9bxXnZwz5w"
	helpAuthor.IconURL = "https://i.imgur.com/BRokk77.png"
	helpEmbed.Author = &helpAuthor

	botCmd := []string{
		"$niji",
		"$onair",
		"$sui",
		"$short [url]",
		"$avatar [@user]",
		"$ip",
		"$uptime",
		"$activity [status]",
		"$kick [@user]",
	}
	botCmdDescription := map[string]string{
		"$niji":              "List 2434 schedule",
		"$onair":             "List 2434 on-air streaming",
		"$sui":               "Randomly pickup a suisei's song",
		"$short [url]":       "Shorten URL by Bitly",
		"$avatar [@user]":    "Avatar of a user",
		"$ip":                "IP address of marumaru",
		"$uptime":            "Uptime of marumaru",
		"$activity [status]": "Set marumaru status",
		"$kick [@user]":      "Ayame kick",
	}

	for _, cmd := range botCmd {
		var cmdObj discordgo.MessageEmbedField

		cmdObj.Name = cmd
		cmdObj.Value = botCmdDescription[cmd]

		helpEmbed.Fields = append(helpEmbed.Fields, &cmdObj)
	}

	var embedFooter discordgo.MessageEmbedFooter
	embedFooter.Text = "ksld154 Inc."
	embedFooter.IconURL = "https://i.imgur.com/q8r4Ue6.jpg"
	helpEmbed.Footer = &embedFooter

	return helpEmbed
}
