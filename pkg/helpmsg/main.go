package helpmsg

import (
	"github.com/bwmarrin/discordgo"
)

// BuildHelpMsg returns a embed of cmd description
func BuildHelpMsg() discordgo.MessageEmbed {
	var helpEmbed discordgo.MessageEmbed
	helpEmbed.Title = "$help"

	var helpAuthor discordgo.MessageEmbedAuthor
	helpAuthor.Name = "Marumaru"
	helpAuthor.URL = "https://www.youtube.com/channel/UCdpUojq0KWZCN9bxXnZwz5w"
	helpAuthor.IconURL = "https://i.imgur.com/xPiRLCv.png"
	helpEmbed.Author = &helpAuthor

	botCmd := []string{
		"$niji",
		"$ip",
		"$gif",
	}
	botCmdDescription := map[string]string{
		"$niji": "List 2434 schedule",
		"$ip":   "IP address of marumaru",
		"$gif":  "WIP",
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

// func main() {
// 	fmt.Println("Hello, help")
// }
