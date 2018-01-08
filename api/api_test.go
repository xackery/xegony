package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/xackery/xegony/storage/mariadb"
)

var (
	apiKey string
)

type fieldResp struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"`
}

type Endpoint struct {
	name         string
	path         string
	method       string
	body         interface{}
	responseCode int
	response     interface{}
	useAuth      bool
}

func initializeServer(t *testing.T) {
	var err error

	s := &mariadb.Storage{}
	err = s.Initialize("root@tcp(127.0.0.1:3306)/eqemu_test?charset=utf8&parseTime=true", ioutil.Discard)
	assert.Nil(t, err)

	err = s.DropTables()
	assert.Nil(t, err)

	err = s.VerifyTables()
	assert.Nil(t, err)

	s.InsertTestData()
	assert.Nil(t, err)

	router := mux.NewRouter().StrictSlash(true)
	apiServer := API{}
	config := ""
	listen := ":8081"
	apiServer.Initialize(s, config, ioutil.Discard)
	assert.Nil(t, err)

	apiServer.ApplyRoutes(router)
	go http.ListenAndServe(listen, router)
	return
}

/*func getConfig() string {
	return "root@tcp(127.0.0.1:3306)/eqemu?charset=utf8&parseTime=true"
}*/

func getURL() string {
	return "http://localhost:8081"
}

func doHTTPTest(test Endpoint, t *testing.T) string {

	var req *http.Request
	client := &http.Client{}
	url := getURL()
	var err error
	var bData []byte

	switch v := test.body.(type) {
	case string:
		if len(v) == 0 {
			req, err = http.NewRequest(test.method, url+test.path, nil)
			assert.Nil(t, err)
		}
		if test.method == "POST" || test.method == "PUT" {
			req, err = http.NewRequest(test.method, url+test.path, strings.NewReader(v))
			assert.Nil(t, err)
			req.Header.Add("Content-Length", strconv.Itoa(len(v)))
		}
	case nil:
		req, err = http.NewRequest(test.method, url+test.path, nil)
		assert.Nil(t, err)
	default:
		bData, err = json.Marshal(v)
		assert.Nil(t, err)

		if test.method == "POST" || test.method == "PUT" {
			req, err = http.NewRequest(test.method, url+test.path, bytes.NewReader(bData))
			req.Header.Add("Content-Length", strconv.Itoa(len(bData)))
		}
	}

	if req == nil {
		req, err = http.NewRequest(test.method, url+test.path, nil)
	}

	if test.useAuth {
		if len(apiKey) == 0 {
			getAuthKey(t)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	assert.Nil(t, err)

	actual, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	actualStr := strings.TrimSpace(string(actual))
	assert.Equal(t, test.responseCode, resp.StatusCode, test.name, test.method, test.path, actualStr)

	response := ""
	switch v := test.response.(type) {
	case string:
		response = v
	case nil:
		if len(actualStr) > 4 {
			t.Fatalf("%s %s %s failed, expected blank, got %s", test.name, test.method, test.path, actualStr)
		}
		return ""
	default:
		bData, err = json.Marshal(v)
		if err != nil {
			t.Fatalf("%s %s %s failed marshalling response json: %s", test.name, test.method, test.path, err.Error())
		}

		response = string(bData)
	}

	assert.Equal(t, response, actualStr)
	return actualStr
}

func getAuthKey(t *testing.T) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", getURL()+"/api/login", strings.NewReader(`{"name":"Test","password":"somepass"}`))
	assert.Nil(t, err)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("%s %s %s failed: %s", "Get Login", "POST", "/login", err.Error())
	}
	if resp.StatusCode != 200 {
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Failed to get resp.body on authkey get")
		}
		t.Fatalf("Failed to get auth key, bad status code %d: %s", resp.StatusCode, string(actual))
	}

	loginResp := loginResponse{}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&loginResp)
	assert.Nil(t, err)
	apiKey = loginResp.APIKey
	if len(apiKey) < 1 {
		t.Fatal("Failed to get token (empty response)")
	}
	return
}

func TestRestEndpoints(t *testing.T) {
	var err error
	s := &mariadb.Storage{}
	s.Initialize("root@tcp(127.0.0.1:3306)/eqemu_test?charset=utf8&parseTime=true", ioutil.Discard)
	assert.Nil(t, err)

	initializeServer(t)

	tests := []Endpoint{
		{
			name:         "GetIndex",
			path:         "/api/",
			method:       "GET",
			body:         "",
			responseCode: 200,
			response:     `{"message":"Please refer to documentation for more details"}`,
			useAuth:      false,
		},
	}

	for _, test := range tests {
		doHTTPTest(test, t)
	}

}
