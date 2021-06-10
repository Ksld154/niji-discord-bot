package ytpicker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type resource struct {
	VideoID string `json:"videoId"`
}

type snippet struct {
	StreamTitle string   `json:"title"`
	ResourceID  resource `json:"resourceId"`
}

type item struct {
	Snippet snippet `json:"snippet"`
}

type jsonResponse struct {
	NextPageToken string `json:"nextPageToken"`
	Items         []item `json:"items"`
}

var allVideoID []string

func parseJSON(httpResp []byte) string {
	var payload jsonResponse
	json.Unmarshal(httpResp, &payload)

	// fmt.Println(payload.NextPageToken)
	for _, video := range payload.Items {
		// fmt.Println(video.Snippet.StreamTitle)
		// fmt.Println(video.Snippet.ResourceID.VideoID)
		allVideoID = append(allVideoID, video.Snippet.ResourceID.VideoID)
	}

	return payload.NextPageToken
}

func getYoutubePlaylistItems(endPoint string, playListID string, pageToken string) (string, error) {

	client := &http.Client{}
	youtubeAPIKey := os.Getenv("YOUTUBE_API_KEY")
	// fmt.Println(youtubeAPIKey)

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()

	q.Add("key", youtubeAPIKey)
	q.Add("playlistId", playListID)
	q.Add("pageToken", pageToken)
	q.Add("part", "snippet")
	q.Add("maxResults", "50")
	q.Add("fields", "nextPageToken, items/snippet(title, resourceId/videoId)")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	nextPageToken := parseJSON(body)

	return nextPageToken, nil
}

func pickupVideo() string {

	rand.Seed(time.Now().Unix())
	idx := rand.Intn(len(allVideoID) - 1)

	videoID := allVideoID[idx]

	videoURL := "https://youtu.be/" + videoID

	return videoURL
}

// GetRandomSong randomly pick up a video from the given playlist
func GetRandomSong(endPoint string, playListID string) string {
	fmt.Println("suisui")

	// do all http request first
	pageToken := ""
	for {
		nextPageToken, err := getYoutubePlaylistItems(endPoint, playListID, pageToken)
		// fmt.Println(nextPageToken)
		if err != nil {
			log.Fatal(err)
		}

		if nextPageToken == "" {
			break
		}

		pageToken = nextPageToken
	}

	// pickup a video randomly
	videoURL := pickupVideo()
	// fmt.Println(videoURL)
	fmt.Println(len(allVideoID))

	allVideoID = nil

	return videoURL
}
