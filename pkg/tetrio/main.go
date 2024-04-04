package tetrio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	tetrioApiEndpoint = "https://ch.tetr.io/api"
)

type Response struct {
	Success bool `json:"success"`
	Data    User `json:"data"`
}

type User struct {
	User UserInfo `json:"user"`
}

type UserInfo struct {
	Username  string `json:"username"`
	ID        string `json:"_id"`
	Role      string `json:"role"`
	Timestamp string `json:"ts"`
	// Badges         []Badge                `json:"badges"`
	XP             float64                `json:"xp"`
	GamesPlayed    int                    `json:"gamesplayed"`
	GamesWon       int                    `json:"gameswon"`
	GameTime       float64                `json:"gametime"`
	Country        string                 `json:"country"`
	SupporterTier  int                    `json:"supporter_tier"`
	Verified       bool                   `json:"verified"`
	League         League                 `json:"league"`
	AvatarRevision int64                  `json:"avatar_revision"`
	Connections    map[string]interface{} `json:"connections"`
	FriendCount    int                    `json:"friend_count"`
}

type Badge struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Group     string `json:"group"`
	Timestamp string `json:"ts"`
}

type League struct {
	GamesPlayed    int     `json:"gamesplayed"`
	GamesWon       int     `json:"gameswon"`
	Rating         float64 `json:"rating"`
	Glicko         float64 `json:"glicko"`
	RD             float64 `json:"rd"`
	Rank           string  `json:"rank"`
	BestRank       string  `json:"bestrank"`
	APM            float64 `json:"apm"`
	PPS            float64 `json:"pps"`
	VS             float64 `json:"vs"`
	Decaying       bool    `json:"decaying"`
	Standing       int     `json:"standing"`
	Percentile     float64 `json:"percentile"`
	StandingLocal  int     `json:"standing_local"`
	PreviousRank   string  `json:"prev_rank"`
	PreviousAt     int     `json:"prev_at"`
	NextRank       string  `json:"next_rank"`
	NextAt         int     `json:"next_at"`
	PercentileRank string  `json:"percentile_rank"`
}

func GetUserInfo(userName string) (UserInfo, error) {
	userName = strings.ToLower(userName)
	url := fmt.Sprintf("%s/users/%s", tetrioApiEndpoint, userName)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return UserInfo{}, err
	}

	var apiResp Response
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return UserInfo{}, err
	}

	return apiResp.Data.User, nil
}

func retrieveUsefulFields(userInfo UserInfo) ([]string, map[string]string) {
	userInfoFields := make(map[string]string)
	fields := []string{"Rank", "Rating", "Standing", "Local Standing", "Games Played", "Games Won", "Percentile", "APM", "PPS", "VS"}

	userInfoFields["Rank"] = strings.ToUpper(userInfo.League.Rank)
	userInfoFields["Rating"] = fmt.Sprintf("%.2f", userInfo.League.Rating)
	userInfoFields["Standing"] = fmt.Sprintf("#%d", userInfo.League.Standing)
	userInfoFields["Local Standing"] = fmt.Sprintf("#%d", userInfo.League.StandingLocal)
	userInfoFields["Games Played"] = strconv.Itoa(userInfo.League.GamesPlayed)
	userInfoFields["Games Won"] = fmt.Sprintf("%d (%.2f%%)", userInfo.League.GamesWon, float64(userInfo.League.GamesWon)/float64(userInfo.League.GamesPlayed)*100)
	userInfoFields["Percentile"] = fmt.Sprintf("%.2f%%", userInfo.League.Percentile*100)
	userInfoFields["APM"] = fmt.Sprintf("%.2f", userInfo.League.APM)
	userInfoFields["PPS"] = fmt.Sprintf("%.2f", userInfo.League.PPS)
	userInfoFields["VS"] = fmt.Sprintf("%.2f", userInfo.League.VS)

	return fields, userInfoFields
}

func generateUserAvatar(userID string, avatarRevision int64) string {
	return fmt.Sprintf("https://tetr.io/user-content/avatars/%s.jpg?rv=%d", userID, avatarRevision)
}

func buildEmbedMsg(userInfo UserInfo, keys []string, userInfoFields map[string]string) discordgo.MessageEmbed {
	var userInfoEmbed discordgo.MessageEmbed
	userInfoEmbed.Title = "Tetrio User Info"
	userInfoEmbed.Color = 0x00ff00

	var embedMsgThumbnail discordgo.MessageEmbedThumbnail
	embedMsgThumbnail.URL = generateUserAvatar(userInfo.ID, userInfo.AvatarRevision)
	userInfoEmbed.Thumbnail = &embedMsgThumbnail

	userInfoEmbed.Title = userInfo.Username
	userInfoEmbed.URL = fmt.Sprintf("https://ch.tetr.io/u/%s", userInfo.ID)

	for _, key := range keys {
		var field discordgo.MessageEmbedField
		field.Name = key
		field.Value = userInfoFields[key]
		field.Inline = true

		userInfoEmbed.Fields = append(userInfoEmbed.Fields, &field)
	}

	return userInfoEmbed
}

func GetUserInfoEmbed(userInfo UserInfo) discordgo.MessageEmbed {
	keys, userInfoFields := retrieveUsefulFields(userInfo)
	embedMsg := buildEmbedMsg(userInfo, keys, userInfoFields)
	return embedMsg
}
