package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetUpTime(t *testing.T) {

	TIME_LAYOUT := "2006/01/02 15:04:05"
	BotStartTime, _ = time.Parse(TIME_LAYOUT, "2021/06/01 15:04:05")
	fmt.Println(BotStartTime.String()) // fake bot start time

	// uptime, _ := strconv.Atoi(GetUpTime(BotStartTime))
	// assert.Greater(t, uptime, 0)

	type args struct {
		startTime time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_GetUpTime_Success",
			args: args{
				startTime: BotStartTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := GetUpTime(BotStartTime)
			fmt.Println(got)

			if got == "0s" {
				t.Errorf("GetUpTime() == 0")
			}
		})
	}
}

func Test_GetOutBoundIP(t *testing.T) {

	// mock Normal(2xx) http server
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// generate and return response
		successResp := []byte(`140.113.195.73`)
		rw.WriteHeader(http.StatusOK)
		rw.Write(successResp)
	}))
	defer testServer.Close()

	// mock Bad(404) http server
	testBadServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		failedResp := []byte(`Server Broken!`)
		rw.WriteHeader(http.StatusNotFound)
		rw.Write(failedResp)
	}))
	defer testBadServer.Close()

	assert.Equal(t, GetOutBoundIPAddr(testServer.URL), "140.113.195.73")
	assert.Equal(t, GetOutBoundIPAddr(testBadServer.URL), "")
}
