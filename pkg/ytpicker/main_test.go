package ytpicker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	youtubeSuiseiID = "PLarzPAXT9RrpT2W_KUTOzNfaYZ4Qs9_-D"
)

func TestGetYoutubePlaylistItems(t *testing.T) {

	// testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

	// 	// generate and return response
	// 	data, _ := ioutil.ReadFile("../../test/successYoutubeResp.json")

	// 	successResp := []byte(data)
	// 	rw.WriteHeader(http.StatusOK)
	// 	rw.Write(successResp)
	// }))
	// defer testServer.Close()

	// pageToken, _ := getYoutubePlaylistItems(testServer.URL, youtubeSuiseiID, "")
	// fmt.Println(pageToken)
	// assert.Equal(t, pageToken, "CDIQAA")

}

func TestGetRandomSong(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		pageToken := req.URL.Query().Get("pageToken")
		fmt.Println(pageToken)
		// fmt.Println(req.RequestURI)

		// generate and return response
		var data []byte
		if pageToken == "CDIQAA" {
			data, _ = ioutil.ReadFile("../../test/lastYoutubeResp.json")
		} else {
			data, _ = ioutil.ReadFile("../../test/successYoutubeResp.json")
		}

		successResp := []byte(data)
		rw.WriteHeader(http.StatusOK)
		rw.Write(successResp)
	}))
	defer testServer.Close()

	// Normal Test case
	videoID := GetRandomSong(testServer.URL, youtubeSuiseiID)
	fmt.Println(videoID)
	assert.NotEqual(t, videoID, "")
}
