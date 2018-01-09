package api

import (
	"net/http"
	"testing"

	"github.com/xackery/xegony/model"
)

func TestActivityEndpoints(t *testing.T) {
	initializeServer(t)

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

	doHTTPTest(t, Endpoint{
		name:         "GetActivityBadID",
		path:         "/api/task/abc/asd",
		method:       "GET",
		body:         ``,
		responseCode: 400,
		response:     `{"message":"activityID argument is required: Invalid arguments provided"}`,
		useAuth:      false,
	})

	doHTTPTest(t, Endpoint{
		name:         "GetActivityBadTaskID",
		path:         "/api/task/abc/456",
		method:       "GET",
		body:         ``,
		responseCode: 400,
		response:     `{"message":"taskID argument is required: Invalid arguments provided"}`,
		useAuth:      false,
	})

	doHTTPTest(t, Endpoint{
		name:         "GetActivity",
		path:         "/api/task/1/1",
		method:       "GET",
		body:         ``,
		responseCode: 200,
		response:     nil,
		useAuth:      false,
	})

	doHTTPTest(t, Endpoint{
		name:         "CreateBlankActivity",
		path:         "/api/task/1",
		method:       "POST",
		body:         model.Activity{},
		responseCode: 401,
		response: &resp{
			Message: "Administrator access required",
		},
		useAuth: false,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateBlankActivityAdmin",
		path:         "/api/task/1",
		method:       "POST",
		body:         model.Activity{},
		responseCode: 500,
		response: resp{
			Message: "Failed to verify TaskID: sql: no rows in result set",
		},
		useAuth: true,
	})

	doHTTPTest(t, Endpoint{
		name:         "CreateActivityBadAdmin",
		path:         "/api/task/1",
		method:       "POST",
		body:         "asdf",
		responseCode: 405,
		response:     `{"message":"Failed to decode body"}`,
		useAuth:      true,
	})

	doHTTPTest(t, Endpoint{
		name:   "CreateActivityAdmin",
		path:   "/api/task/1",
		method: "POST",
		body: model.Activity{
			TaskID:       1,
			ZoneID:       101,
			ActivityType: 1,
		},
		responseCode: http.StatusCreated,
		response: model.Activity{
			TaskID:       1,
			ZoneID:       101,
			ActivityType: 1,
		},
		useAuth: true,
	})

	doHTTPTest(t, Endpoint{
		name:         "GetActivity",
		path:         "/api/task/1/2",
		method:       "GET",
		body:         ``,
		responseCode: 200,
		response:     nil,
		useAuth:      false,
	})
}
