package api

import (
	"testing"

	"github.com/xackery/xegony/model"
)

func TestCharacterEndpoints(t *testing.T) {
	initializeServer(t)

	type fieldResp struct {
		Message string            `json:"message"`
		Fields  map[string]string `json:"fields"`
	}

	doHTTPTest(t, Endpoint{
		name:         "CreateCharacterInvalidNameReq",
		path:         "/api/character",
		method:       "POST",
		body:         &model.Character{AccountID: 1},
		responseCode: 400,
		response: &fieldResp{
			Message: "String length must be greater than or equal to 3",
			Fields: map[string]string{
				"name": "String length must be greater than or equal to 3",
			},
		},
		useAuth: true,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateCharacterFailDecode",
		path:         "/api/character",
		method:       "POST",
		body:         `{"id":"abc",name":"Test"}`,
		responseCode: 405,
		response:     `{"message":"Failed to decode body"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateCharacterMinNameFail",
		path:         "/api/character",
		method:       "POST",
		body:         `{"accountID":1,"name":"Te"}`,
		responseCode: 400,
		response: &fieldResp{
			Message: "String length must be greater than or equal to 3",
			Fields: map[string]string{
				"name": "String length must be greater than or equal to 3",
			},
		},
		useAuth: true,
	})

	doHTTPTest(t, Endpoint{
		name:         "CreateCharacterMaxNameFail",
		path:         "/api/character",
		method:       "POST",
		body:         `{"name":"Tsidofjsdoifjsdofijsdofijsdofisjdfoisj","accountID":1,"image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"String length must be less than or equal to 32","fields":{"name":"String length must be less than or equal to 32"}}`,
		useAuth:      true,
	})

	doHTTPTest(t, Endpoint{
		name:   "CreateCharacterBadCharFail",
		path:   "/api/character",
		method: "POST",
		body: &model.Character{
			ID:        1,
			AccountID: 1,
			Name:      "Testø∆£ø´∆™",
		},
		responseCode: 400,
		response:     `{"message":"Does not match pattern '^[a-zA-Z]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z]*$'"}}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:   "CreateCharacterNotAdmin",
		path:   "/api/character",
		method: "POST",
		body: &model.Character{
			ID:        1,
			AccountID: 1,
			Name:      "Test",
		},
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:   "CreateCharacter",
		path:   "/api/character",
		method: "POST",
		body: &model.Character{
			AccountID: 1,
			Name:      "Test",
		},
		responseCode: 201,
		response: &model.Character{
			ID:        1,
			AccountID: 1,
			Name:      "Test",
		},
		useAuth: true,
	})
	doHTTPTest(t, Endpoint{
		name:         "GetCharacterInvalidCharacterId",
		path:         "/api/character/invalid",
		method:       "GET",
		body:         "",
		responseCode: 400,
		response:     `{"message":"characterID argument is required: Invalid arguments provided"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "GetCharacterNoResults",
		path:         "/api/character/2",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response:     nil,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "GetCharacter",
		path:         "/api/character/1",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response: &model.Character{
			ID:        1,
			AccountID: 1,
			Name:      "Test",
		},
		useAuth: false,
	})

	doHTTPTest(t, Endpoint{
		name:         "GetCharacterByName",
		path:         "/api/character/byname/Test",
		method:       "GET",
		body:         ``,
		responseCode: 200,
		response: &model.Character{
			ID:        1,
			Name:      "Test",
			AccountID: 1,
		},
		useAuth: false,
	})

	doHTTPTest(t, Endpoint{
		name:         "NoTokenEditCharacter",
		path:         "/api/character/invalid",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 401,
		response:     `{"message":"Moderator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "InvalidEditCharacter",
		path:         "/api/character/invalid",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"characterID argument is required: Invalid arguments provided"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "RequestFailEditCharacter",
		path:         "/api/character/4",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"Does not match pattern '^[a-zA-Z]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z]*$'"}}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:   "EditCharacter",
		path:   "/api/character/1",
		method: "PUT",
		body: &model.Character{
			ID:   1,
			Name: "TestTwo",
		},
		responseCode: 200,
		response: &model.Character{
			ID:   1,
			Name: "TestTwo",
		},
		useAuth: true,
	})
	//This should be a 304 / no content response
	doHTTPTest(t, Endpoint{
		name:         "ErrNoContentEditCharacter",
		path:         "/api/character/1",
		method:       "PUT",
		body:         `{"id":1,"name":"TestTwo","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 304,
		response:     nil,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "EditCharacterInvalid",
		path:         "/api/character/1",
		method:       "PUT",
		body:         `{"id":"abc",name":"Test"}`,
		responseCode: 405,
		response:     `{"message":"Request error: Failed to decode body"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "ListCharacter",
		path:         "/api/character",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response: []*model.Character{
			&model.Character{
				ID:   1,
				Name: "TestTwo",
			},
		},
		useAuth: false,
	})
	doHTTPTest(t, Endpoint{
		name:         "InvalidIdDeleteCharacter",
		path:         "/api/character/{invalid}",
		method:       "DELETE",
		body:         "",
		responseCode: 400,
		response:     `{"message":"characterID argument is required: Invalid arguments provided"}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "RequestFailDeleteCharacter",
		path:         "/api/character/3",
		method:       "DELETE",
		body:         "",
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "DeleteCharacterNotLoggedIn",
		path:         "/api/character/1",
		method:       "DELETE",
		body:         "",
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "DeleteCharacter",
		path:         "/api/character/1",
		method:       "DELETE",
		body:         "",
		responseCode: 204,
		response:     ``,
		useAuth:      true,
	})

}
