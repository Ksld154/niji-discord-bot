package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// BotStartTime is start time of this bot
var BotStartTime time.Time

// GetUpTime return the execution time of bot
func GetUpTime() string {
	nowTime := time.Now()
	diff := nowTime.Sub(BotStartTime)

	return diff.Round(time.Second).String()
}

// GetOutBoundIPAddr return IP address of bot
func GetOutBoundIPAddr() string {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(body)
}
