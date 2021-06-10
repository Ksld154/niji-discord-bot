package bitly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func Test_shortenURL(t *testing.T) {

// 	type args struct {
// 		longURL string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "Test_shortenURL_Success",
// 			args: args{
// 				longURL: "https://www.youtube.com/",
// 			},
// 			want:    "https://bit.ly/3byvFEv",
// 			wantErr: false,
// 		},
// 		{
// 			name: "Test_shortenURL_Failed",
// 			args: args{
// 				longURL: "https://badurl.com/",
// 			},
// 			want:    "",
// 			wantErr: true,
// 		},
// 	}

// 	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		fmt.Println("test server")

// 		// get request body
// 		bodyBytes, err := ioutil.ReadAll(req.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// get long_url from request
// 		requestFields := make(map[string]interface{})
// 		err = json.Unmarshal(bodyBytes, &requestFields)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(requestFields["long_url"])

// 		// generate and return response
// 		successResp := []byte(`{"link": "https://bit.ly/3byvFEv"}`)
// 		failResp := []byte{}
// 		if requestFields["long_url"] == "https://www.youtube.com/" {
// 			rw.WriteHeader(http.StatusOK)
// 			rw.Write(successResp)
// 		} else {
// 			rw.WriteHeader(http.StatusNotFound)
// 			rw.Write(failResp)
// 		}
// 	}))
// 	defer testServer.Close()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			got, err := shortenURL(tt.args.longURL, testServer.URL)
// 			fmt.Println()

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("shortenURL() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("shortenURL() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetShortURL(t *testing.T) {
	type args struct {
		longURL string
		// endPoint string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test_GetShortURL_Success",
			args: args{
				longURL: "https://www.youtube.com/",
			},
			want:    "https://bit.ly/3byvFEv",
			wantErr: false,
		},
		// {
		// 	name: "Test_GetShortURL_Failed_1xx",
		// 	args: args{
		// 		longURL: "https://bad_1xx.com/",
		// 	},
		// 	want:    "",
		// 	wantErr: true,
		// },
		{
			name: "Test_GetShortURL_Failed_4xx",
			args: args{
				longURL: "https://bad_4xx.com/",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test_GetShortURL_InvalidURL",
			args: args{
				longURL: "5566.com",
			},
			want:    "",
			wantErr: true,
		},
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("test server")

		// get request body
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		// get long_url from request
		requestFields := make(map[string]interface{})
		err = json.Unmarshal(bodyBytes, &requestFields)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(requestFields["long_url"])

		// generate and return response
		// test "if resp.StatusCode > 299"
		successResp := []byte(`{"link": "https://bit.ly/3byvFEv"}`)
		failResp := []byte{}
		if requestFields["long_url"] == "https://www.youtube.com/" {
			rw.WriteHeader(http.StatusOK) // 200 < 299
			rw.Write(successResp)
		} else if requestFields["long_url"] == "https://bad_4xx.com/" {
			rw.WriteHeader(http.StatusTeapot) // 404 > 299
			rw.Write(failResp)
		}
	}))
	defer testServer.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetShortURL(tt.args.longURL, testServer.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
