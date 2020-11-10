package nijionair

import (
	"fmt"
	"log"
	"strings"

	"github.com/Ksld154/niji-discord-bot/pkg/nijiparser"
	"github.com/anaskhan96/soup"
	"github.com/bwmarrin/discordgo"
)

// OnAirStream contains info of a nijisanji on-air stream
type OnAirStream struct {
	Liver string
	URL   string
	Icon  string
}

func nijiOnAirScrapper() []OnAirStream {

	var Streams []OnAirStream
	var onAirLivers []string
	var onAirLiverIcons []string
	var OnAirLinks []string

	resp, err := soup.Get("https://nijisanji.vtubervideo.net/lives")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	doc := soup.HTMLParse(resp)

	onair := doc.FindAll("div", "class", "header")
	for _, liver := range onair {
		onAirLivers = append(onAirLivers, liver.Text())
	}

	// last item is not streamming liver
	if len(onAirLivers) > 0 {
		onAirLivers = onAirLivers[:len(onAirLivers)-1]
	}

	links := doc.FindAllStrict("div", "class", "player embed")
	for _, link := range links {

		youtubeID := strings.Split(link.Attrs()["id"], "player-")[1]
		// fmt.Println(youtubeID)
		OnAirLinks = append(OnAirLinks, youtubeID)
	}

	icons := doc.FindAllStrict("img", "class", "ui avatar image")
	for _, icon := range icons {
		onAirLiverIcons = append(onAirLiverIcons, icon.Attrs()["src"])
	}

	for i := 0; i < len(onAirLivers); i++ {
		var stream OnAirStream
		stream.Liver = onAirLivers[i]
		stream.URL = "https://youtu.be/" + OnAirLinks[i]
		stream.Icon = onAirLiverIcons[i]

		Streams = append(Streams, stream)
		// fmt.Println(stream)
	}

	return Streams
}

func buildOnAirEmbed(streams []OnAirStream) discordgo.MessageEmbed {
	var onAirEmbed discordgo.MessageEmbed
	onAirEmbed.Title = "🌈 2434 On-Air Streaming"
	onAirEmbed.URL = "https://nijisanji.vtubervideo.net/lives"
	onAirEmbed.Color = 0xffdd00

	var streamObj discordgo.MessageEmbedField
	streamObj.Name = "List"
	streamObj.Value = ""
	for _, stream := range streams {
		streamObjLink := fmt.Sprintf("[%s%s](%s) \n", nijiparser.GetLiverEmoji(stream.Liver), stream.Liver, stream.URL)
		streamObj.Value += streamObjLink
	}
	onAirEmbed.Fields = append(onAirEmbed.Fields, &streamObj)

	var footer discordgo.MessageEmbedFooter
	footer.Text = "にじさんじTool"
	footer.IconURL = "https://i.imgur.com/6DXmM6p.jpg"
	onAirEmbed.Footer = &footer

	return onAirEmbed
}

// GetNijiOnAir returns embed msg that contain nijisanji on-air streaming
func GetNijiOnAir() discordgo.MessageEmbed {
	streams := nijiOnAirScrapper()
	onair := buildOnAirEmbed(streams)

	return onair
}
