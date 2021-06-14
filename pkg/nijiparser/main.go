package nijiparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"time"

	"github.com/bwmarrin/discordgo"
)

// const url string = "https://api.itsukaralink.jp/v1.2/events.json"

type liverInfo struct {
	Name string `json:"name"`
}

type streamInfo struct {
	Title     string      `json:"name"`
	StartDate string      `json:"start_date"`
	URL       string      `json:"url"`
	Liver     []liverInfo `json:"livers"`
}

type eventInfo struct {
	Events []streamInfo `json:"events"`
}

type jsonResponse struct {
	Status string    `json:"status"`
	Data   eventInfo `json:"data"`
}

type streamInfoSimple struct {
	Liver string
	URL   string
}

// GetLiverEmoji return the emoji of given nijisanji-liver's JP name
func GetLiverEmoji(liver string) string {

	emojiMap := map[string]string{
		"月ノ美兎":         "🐰",
		"勇気ちひろ":        "🎀💙",
		"える":           "🗼",
		"樋口楓":          "🍁",
		"静凛":           "💜",
		"渋谷ハジメ":        "🌱",
		"鈴谷アキ":         "🐈",
		"モイラ":          "🐶",
		"鈴鹿詩子":         "🎶",
		"宇志海いちご":       "🍓",
		"家長むぎ":         "🌷",
		"夕陽リリ":         "🌇",
		"物述有栖":         "♥️♠️♦️♣️",
		"文野環":          "🐟",
		"伏見ガク":         "🦊",
		"ギルザレンⅢ世":      "🏰🌕",
		"剣持刀也":         "⚔",
		"森中花咲":         "🐻",
		"叶":            "🔫",
		"赤羽葉子":         "💀",
		"笹木咲":          "🎋",
		"本間ひまわり":       "🌻",
		"魔界ノりりむ":       "🍼",
		"葛葉":           "🎲",
		"椎名唯華":         "👻",
		"ドーラ":          "🔥",
		"出雲霞":          "🦑",
		"轟京子":          "🐐",
		"シスター・クレア":     "🔔",
		"花畑チャイカ":       "🌵",
		"社築":           "🖥",
		"安土桃":          "🍑",
		"鈴木勝":          "☪",
		"緑仙":           "🐼",
		"卯月コウ":         "🌙",
		"神田笑一":         "🔪",
		"飛鳥ひな":         "🐤",
		"春崎エアル":        "🍭",
		"雨森小夜":         "☔",
		"鷹宮リオン":        "🦅",
		"舞元啓介":         "👨‍🌾",
		"竜胆尊":          "🍶⚜️",
		"でびでび・でびる":     "🚪👿",
		"桜凛月":          "🌸",
		"町田ちま":         "🐹",
		"ジョー・力一":       "🤡",
		"成瀬鳴":          "🎙",
		"ベルモンド・バンデラス":  "🥃",
		"矢車りね":         "🌽",
		"夢追翔":          "🎤",
		"黒井しば":         "🐕🐾",
		"童田明治":         "🐺🍎",
		"郡道美玲":         "🐽",
		"夢月ロア":         "🌖",
		"小野町春香":        "♨",
		"語部紡":          "🧂📘",
		"瀬戸美夜子":        "📷💚",
		"御伽原江良":        "🏰🕛",
		"戌亥とこ":         "🍹",
		"アンジュ・カトリーナ":   "⚖",
		"リゼ・ヘルエスタ":     "👑",
		"三枝明那":         "🌶",
		"愛園愛美":         "💕",
		"鈴原るる":         "🎨",
		"雪城眞尋":         "🌐💫",
		"エクス・アルビオ":     "🛡",
		"レヴィ・エリファ":     "🔲",
		"葉山舞鈴":         "🍃🗻",
		"ニュイ・ソシエール":    "🎃",
		"葉加瀬冬雪":        "⚗",
		"加賀美ハヤト":       "🏢",
		"夜見れな":         "🎩🐤",
		"黛灰":           "💻💙",
		"アルス・アルマル":     "📕",
		"相羽ういは":        "🍮💎",
		"天宮こころ":        "🎐",
		"エリー・コニファー":    "🌲",
		"ラトナ・プティ":      "🐻💎",
		"早瀬走":          "🚴‍♀️",
		"健屋花那":         "💉💘",
		"シェリン・バーガンディ":  "🧐",
		"フミ":           "🔖",
		"星川サラ":         "🌟",
		"山神カルタ":        "🎴",
		"えま★おうがすと":     "★",
		"ルイス・キャミー":     "❤️🦋",
		"魔使マオ":         "💥",
		"不破湊":          "🥂✨",
		"白雪巴":          "👠⛓",
		"グウェル・オス・ガール":  "😎",
		"ましろ":          "🧷",
		"奈羅花":          "✖🍳",
		"来栖夏芽":         "🐏🎵",
		"フレン・E・ルスタリオ":  "🎠",
		"メリッサ・キンレンカ":   "🐝",
		"イブラヒム":        "💧",
		"長尾景":          "☯",
		"弦月藤士郎":        "🎻🛵",
		"甲斐田晴":         "🌞",
		"空星きらめ":        "🌌",
		"金魚坂めいろ":       "🩰",
		"にじさんじ公式チャンネル": "🌈🕒",
		"朝日南アカネ":       "🦖🎖",
		"周央サンゴ":        "💞🦩",
		"東堂コハク":        "🍯",
		"北小路ヒスイ":       "❇",
		"西園チグサ":        "🐬🌱",
	}

	return emojiMap[liver]
}

func getScheduleThumbnail() string {
	thumbnailList := []string{
		"https://i.imgur.com/PfSZpsi.gif",
		"https://i.imgur.com/Dd0X71h.gif",
		"https://i.imgur.com/OdODO3N.gif",
		"https://i.imgur.com/wuqaSbs.gif",
		"https://i.imgur.com/aonsNxP.gif",
		"https://i.imgur.com/zf2l5aO.gif",
	}

	idx := rand.Intn(len(thumbnailList))

	return thumbnailList[idx]
}

func parseJSON(jsonBody []byte) ([]string, map[string][]streamInfoSimple) {

	schedule := make(map[string][]streamInfoSimple, 100)

	oneHourAgo := time.Now().Add(-1 * time.Hour)
	twTimeZone, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatal(err)
	}

	var payload jsonResponse
	json.Unmarshal(jsonBody, &payload)

	for _, stream := range payload.Data.Events {

		if len(stream.Liver) < 1 {
			continue
		}

		streamTime, err := time.Parse(time.RFC3339, stream.StartDate)
		if err != nil {
			log.Fatal(err)
			continue
		}
		// fmt.Println(streamTime)
		// change streamTime obj's timezone first
		twStreamTime := streamTime.In(twTimeZone)
		if streamTime.Before(oneHourAgo) {
			continue
		}
		// fmt.Println(streamTime)

		formattedStreamTime := fmt.Sprintf("%02d/%02d %02d:%02d",
			twStreamTime.Month(), twStreamTime.Day(),
			twStreamTime.Hour(), twStreamTime.Minute(),
		)

		var streamObj streamInfoSimple
		streamObj.Liver = stream.Liver[0].Name
		streamObj.URL = stream.URL

		schedule[formattedStreamTime] = append(schedule[formattedStreamTime], streamObj)
	}

	// sort the struct by stream_start_time
	var keys []string
	for k := range schedule {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys, schedule
}

func buildEmbedMsg(timeKeys []string, schedule map[string][]streamInfoSimple) discordgo.MessageEmbed {
	var scheduleEmbed discordgo.MessageEmbed
	scheduleEmbed.Title = "🌈 2434 Schedule"
	scheduleEmbed.URL = "https://www.itsukaralink.jp/"
	scheduleEmbed.Color = 0xffdd00

	var embedMsgThumbnail discordgo.MessageEmbedThumbnail
	embedMsgThumbnail.URL = getScheduleThumbnail()
	scheduleEmbed.Thumbnail = &embedMsgThumbnail

	for _, key := range timeKeys {
		var streamObj discordgo.MessageEmbedField

		streamObj.Name = key
		streamObj.Value = ""
		for _, stream := range schedule[key] {
			streamObjLink := fmt.Sprintf("[%s%s](%s) \n", GetLiverEmoji(stream.Liver), stream.Liver, stream.URL)
			streamObj.Value += streamObjLink
		}

		scheduleEmbed.Fields = append(scheduleEmbed.Fields, &streamObj)
	}

	var embedMsgFooter discordgo.MessageEmbedFooter
	embedMsgFooter.Text = "Ichikara Inc."
	embedMsgFooter.IconURL = "https://i.imgur.com/ipsQ3gX.jpg"
	scheduleEmbed.Footer = &embedMsgFooter

	return scheduleEmbed
}

// NijiScheduleParser should return embed msg that contain nijisanji schedule
func NijiScheduleParser(endpoint string) discordgo.MessageEmbed {
	fmt.Println("### 2434 Schedule Parser! ###")

	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
		return discordgo.MessageEmbed{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return discordgo.MessageEmbed{}
	}

	timeKeys, schedule := parseJSON(body)
	scheduleEmbed := buildEmbedMsg(timeKeys, schedule)

	return scheduleEmbed
}
