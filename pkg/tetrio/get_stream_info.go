package tetrio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type GetStreamInfoResponse struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Records []Record `json:"records"`
}

type Record struct {
	ID         string       `json:"_id"`
	Owner      Owner        `json:"user"`
	Timestamp  string       `json:"ts"`
	EndContext []EndContext `json:"endcontext"`
}

type Owner struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}

type EndContext struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Wins     int    `json:"wins"`
}

type GameResult struct {
	Result        bool   `json:"result"`
	Score         int    `json:"score"`
	Opponent      string `json:"opponent"`
	OpponentScore int    `json:"opponent_score"`
	Timestamp     string `json:"ts"`
}

// GetStreamInfo returns the stream info of the user
func getRecords(userID string) ([]Record, error) {
	streamID := fmt.Sprintf("league_userrecent_%s", userID)
	url := fmt.Sprintf("%s/streams/%s", tetrioApiEndpoint, streamID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp GetStreamInfoResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return nil, err
	}

	if !apiResp.Success {
		return nil, err
	}

	return apiResp.Data.Records, nil
}

func getPlayerInfo(records []Record) Owner {
	if len(records) == 0 {
		return Owner{}
	}
	return records[0].Owner
}

func formatTimestamp(timestamp string) string {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return ""
	}

	taiwan, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return ""
	}

	t = t.In(taiwan)

	return t.Format("1/2 15:04")
}

func parseIntoGameResult(records []Record) []GameResult {
	GameResults := make([]GameResult, 0)
	for _, record := range records {
		// fmt.Println(record)

		if len(record.EndContext) != 2 {
			continue
		}

		timestamp := formatTimestamp(record.Timestamp)

		var gameResult GameResult

		if record.EndContext[0].Username != record.Owner.Username {
			record.EndContext[0], record.EndContext[1] = record.EndContext[1], record.EndContext[0]
		}

		playerInfo := record.EndContext[0]
		opponentInfo := record.EndContext[1]
		gameResult.Timestamp = timestamp
		gameResult.Opponent = opponentInfo.Username
		gameResult.OpponentScore = opponentInfo.Wins
		gameResult.Score = playerInfo.Wins

		if playerInfo.Wins < opponentInfo.Wins {
			gameResult.Result = false
		} else {
			gameResult.Result = true
		}

		GameResults = append(GameResults, gameResult)
	}

	return GameResults
}

func buildStreamEmbedMsg(user Owner, GameResults []GameResult) discordgo.MessageEmbed {
	var gameResultEmbed discordgo.MessageEmbed
	gameResultEmbed.Title = "Tetrio User Records"
	gameResultEmbed.Color = 0x00ff00

	gameResultEmbed.Title = user.Username
	gameResultEmbed.URL = fmt.Sprintf("https://ch.tetr.io/s/league_userrecent_%s", user.ID)

	for _, gameResult := range GameResults {
		// fmt.Println(gameResult)

		var field discordgo.MessageEmbedField
		field.Name = generateGameResultField(gameResult)
		field.Value = gameResult.Timestamp

		gameResultEmbed.Fields = append(gameResultEmbed.Fields, &field)
	}

	return gameResultEmbed
}

func generateGameResultField(gameResult GameResult) string {
	var result string
	if gameResult.Result {
		result = "👑"
	} else {
		result = "🥀"
	}

	result += fmt.Sprintf(" %d-%d vs %s", gameResult.Score, gameResult.OpponentScore, gameResult.Opponent)

	return result
}

func GetStreamInfo(userID string) discordgo.MessageEmbed {
	// 1. Send GET request to Tetrio API
	// 2. Parse the response body into []Record
	// 3. Format the []Record into a embed message
	records, err := getRecords(userID)
	if err != nil {
		return discordgo.MessageEmbed{}
	}

	player := getPlayerInfo(records)
	GameResults := parseIntoGameResult(records)

	// 4. Return the embed message
	embedMsg := buildStreamEmbedMsg(player, GameResults)

	return embedMsg
}
