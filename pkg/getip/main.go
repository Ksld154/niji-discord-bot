package getip

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetOutBOundIPAddr() string {
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
