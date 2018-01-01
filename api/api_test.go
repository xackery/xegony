package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/storage/mariadb"
)

var (
	apiKey string
)

type Endpoint struct {
	name         string
	path         string
	method       string
	body         string
	responseCode int
	response     string
	useAuth      bool
}

func initializeServer(t *testing.T) {
	var err error

	s := &mariadb.Storage{}
	if err = s.Initialize("root@tcp(127.0.0.1:3306)/eqemu_test?charset=utf8&parseTime=true"); err != nil {
		t.Fatalf("Failed to initialize: %s", err.Error())
	}
	if err = s.DropTables(); err != nil {
		t.Fatalf("Failed to drop tables: %s", err.Error())
	}

	if err = s.VerifyTables(); err != nil {
		t.Fatalf("Failed to verify tables: %s", err.Error())
	}
	if err = s.InsertTestData(); err != nil {
		t.Fatalf("Failed to insert test data: %s", err.Error())
	}
	router := mux.NewRouter().StrictSlash(true)
	apiServer := API{}
	config := ""
	listen := ":8081"
	if err = apiServer.Initialize(s, config); err != nil {
		t.Fatal("Failed to initialize apiServer:", err.Error())
	}
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
	if test.method == "POST" || test.method == "PUT" {
		req, err = http.NewRequest(test.method, url+test.path, strings.NewReader(test.body))
		req.Header.Add("Content-Length", strconv.Itoa(len(test.body)))
	} else {
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
	if err != nil {
		t.Fatalf("%s %s %s failed: %s", test.name, test.method, test.path, err.Error())
	}

	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("%s %s %s failed getting body: %s", test.name, test.method, test.path, err.Error())
	}
	actualStr := strings.TrimSpace(string(actual))
	if resp.StatusCode != test.responseCode {
		t.Fatalf("%s %s %s failed response: Expected %d, got %d: %s", test.name, test.method, url+test.path, test.responseCode, resp.StatusCode, actualStr)
	}

	if strings.Index(actualStr, test.response) != 0 {
		t.Fatalf("%s %s %s failed body: Expected %s, got %s", test.name, test.method, test.path, test.response, actualStr)
	}
	return actualStr
}

func getAuthKey(t *testing.T) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", getURL()+"/api/login", strings.NewReader(`{"name":"Test","password":"somepass"}`))
	if err != nil {
		t.Fatal("Failed to get auth key", err.Error())
	}
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
	if err != nil {
		t.Fatalf("Failed to decode login response: %s", err.Error())
	}
	apiKey = loginResp.APIKey
	if len(apiKey) < 1 {
		t.Fatal("Failed to get token (empty response)")
	}
	return
}

func TestRestEndpoints(t *testing.T) {
	var err error
	s := &mariadb.Storage{}
	if err = s.Initialize("root@tcp(127.0.0.1:3306)/eqemu_test?charset=utf8&parseTime=true"); err != nil {
		t.Fatalf("Failed to initialize: %s", err.Error())
	}
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
