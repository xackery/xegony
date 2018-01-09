package api

import (
	"net/http"
	"testing"

	"github.com/xackery/xegony/model"
)

func TestTaskEndpoints(t *testing.T) {
	initializeServer(t)

	doHTTPTest(t, Endpoint{
		name:         "GetTask",
		path:         "/api/task/456/details",
		method:       "GET",
		body:         ``,
		responseCode: 200,
		response:     nil,
		useAuth:      false,
	})

	doHTTPTest(t, Endpoint{
		name:         "GetTask",
		path:         "/api/task/1/details",
		method:       "GET",
		body:         ``,
		responseCode: 200,
		response:     nil,
		useAuth:      false,
	})

	doHTTPTest(t, Endpoint{
		name:         "CreateBlankTask",
		path:         "/api/task",
		method:       "POST",
		body:         model.Task{},
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateBlankTaskAdmin",
		path:         "/api/task",
		method:       "POST",
		body:         model.Task{},
		responseCode: 400,
		response:     `{"message":"String length must be greater than or equal to 3","fields":{"title":"String length must be greater than or equal to 3"}}`,
		useAuth:      true,
	})

	doHTTPTest(t, Endpoint{
		name:   "CreateTaskAdmin",
		path:   "/api/task",
		method: "POST",
		body: model.Task{
			Title: "Testing",
		},
		responseCode: http.StatusCreated,
		response: model.Task{
			ID:    1,
			Title: "Testing",
		},
		useAuth: true,
	})
}
