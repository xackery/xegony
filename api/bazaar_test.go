package api

import (
	"testing"

	"github.com/xackery/xegony/model"
)

func TestBazaarEndpoints(t *testing.T) {
	initializeServer(t)

	doHTTPTest(t, Endpoint{
		name:         "CreateBazaarFailDecode",
		path:         "/api/bazaar",
		method:       "POST",
		body:         `{"id":"abc",name":"Test"}`,
		responseCode: 405,
		response:     `{"message":"Failed to decode body"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateBazaarMinNameFail",
		path:         "/api/bazaar",
		method:       "POST",
		body:         `{"bazaarID":1,"name":"Te"}`,
		responseCode: 400,
		response: &fieldResp{
			Message: "Must be greater than or equal to 1",
			Fields: map[string]string{
				"accountID": "Must be greater than or equal to 1",
				"itemID":    "Must be greater than or equal to 1",
				"price":     "Must be greater than or equal to 1",
			},
		},
		useAuth: true,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateBazaarNotAdmin",
		path:         "/api/bazaar",
		method:       "POST",
		body:         `{"name":"Test","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:   "CreateBazaar",
		path:   "/api/bazaar",
		method: "POST",
		body: &model.Bazaar{
			ItemID:    1,
			AccountID: 1,
			Price:     1,
		},
		responseCode: 201,
		response: &model.Bazaar{
			ID:        1,
			ItemID:    1,
			AccountID: 1,
			Price:     1,
		},
		useAuth: true,
	})
	doHTTPTest(t, Endpoint{
		name:         "GetBazaarInvalidBazaarID",
		path:         "/api/bazaar/invalid",
		method:       "GET",
		body:         "",
		responseCode: 400,
		response:     `{"message":"bazaarID argument is required: Invalid arguments provided"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "GetBazaarNoResults",
		path:         "/api/bazaar/2",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response:     nil,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "GetBazaar",
		path:         "/api/bazaar/1",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response: &model.Bazaar{
			ID:        1,
			ItemID:    1,
			AccountID: 1,
			Price:     1,
		},
		useAuth: false,
	})
	doHTTPTest(t, Endpoint{
		name:         "NoTokenEditBazaar",
		path:         "/api/bazaar/invalid",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 401,
		response:     `{"message":"Moderator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "InvalidEditBazaar",
		path:         "/api/bazaar/invalid",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"bazaarID argument is required: Invalid arguments provided"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:   "RequestFailEditBazaar",
		path:   "/api/bazaar/4",
		method: "PUT",
		body: &model.Bazaar{
			ItemID:    0,
			AccountID: 1,
			Price:     1,
		},
		responseCode: 400,
		response:     `{"message":"Must be greater than or equal to 1","fields":{"itemID":"Must be greater than or equal to 1"}}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:   "EditBazaar",
		path:   "/api/bazaar/2",
		method: "PUT",
		body: &model.Bazaar{
			ItemID:    1,
			AccountID: 1,
			Price:     1,
		},
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:   "NoContentEditBazaar",
		path:   "/api/bazaar/2",
		method: "PUT",
		body: &model.Bazaar{
			ItemID:    1,
			AccountID: 1,
			Price:     1,
		},
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "EditBazaarInvalid",
		path:         "/api/bazaar/1",
		method:       "PUT",
		body:         `{"id":"abc",name":"Test"}`,
		responseCode: 405,
		response:     `{"message":"Request error: Failed to decode body"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "ListBazaar",
		path:         "/api/bazaar",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response:     `[{"Item":null,"id":1,"itemID":1,"accountID":1,"price":1,"createDate":"0001-01-01T00:00:00Z"}]`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "InvalidIdDeleteBazaar",
		path:         "/api/bazaar/{invalid}",
		method:       "DELETE",
		body:         "",
		responseCode: 400,
		response:     `{"message":"bazaarID argument is required: Invalid arguments provided"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "RequestFailDeleteBazaar",
		path:         "/api/bazaar/3",
		method:       "DELETE",
		body:         "",
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "DeleteBazaarNotLoggedIn",
		path:         "/api/bazaar/1",
		method:       "DELETE",
		body:         "",
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "DeleteBazaar",
		path:         "/api/bazaar/1",
		method:       "DELETE",
		body:         "",
		responseCode: 204,
		response:     ``,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "DeleteBazaarInvalidID",
		path:         "/api/bazaar/5",
		method:       "DELETE",
		body:         "",
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	})
}
