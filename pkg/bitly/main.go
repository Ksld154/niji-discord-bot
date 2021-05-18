package bitly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// BitlyToken is bitly's api token
var BitlyToken string

const (
	urlRegexp     = `https?:\/\/[\S]*`
	bitlyEndPoint = "https://api-ssl.bitly.com/v4/shorten"
)

type shortURLObject struct {
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	Link      string `json:"link"`
	LongURL   string `json:"long_url"`
}

// GetShortURL returns the URL shorten by Bitly service
func GetShortURL(longURL string, endPoint string) (string, error) {

	// fmt.Println(longURL)
	urlPattern, err := regexp.Compile(urlRegexp)
	if err != nil {
		return "", err
	}

	if urlPattern.MatchString(longURL) {
		shortURL, err := shortenURL(longURL, endPoint)
		if err != nil {
			return "", err
		}
		return shortURL, nil
	}

	fmt.Println("url formal err")
	err = fmt.Errorf("url format error")
	return "", err
}

func shortenURL(longURL string, endPoint string) (string, error) {

	data := map[string]string{"long_url": longURL}
	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	fmt.Println(endPoint)

	// send POST request to bitly api
	client := &http.Client{}
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BitlyToken)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode > 299 {
		err = fmt.Errorf("[Error] bitly resp code: %d", resp.StatusCode)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var shortURLObj shortURLObject
	err = json.Unmarshal(body, &shortURLObj)
	if err != nil {
		return "", err
	}

	shortURL := shortURLObj.Link
	return shortURL, nil
}
