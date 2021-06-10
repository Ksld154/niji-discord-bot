package nijiparser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNijiScheduleParser(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// generate and return response
		// var data []byte
		data, _ := ioutil.ReadFile("../../test/nijiSchedule_sample.json")

		successResp := []byte(data)
		rw.WriteHeader(http.StatusOK)
		rw.Write(successResp)
	}))
	defer testServer.Close()

	scheduleEmbed := NijiScheduleParser(testServer.URL)
	fmt.Println(scheduleEmbed.Fields[0].Value)

	assert.Equal(t, scheduleEmbed.Fields[0].Value, "[📕アルス・アルマル](https://www.youtube.com/watch?v=er6BElZ2X9Y) \n")
}
