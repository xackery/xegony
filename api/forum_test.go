package api

import (
	"testing"
)

func TestForumEndpoints(t *testing.T) {
	initializeServer(t)

	tests := []Endpoint{
		{
			name:         "CreateForumInvalidNameReq",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"forumId":1}`,
			responseCode: 400,
			response:     `{"message":"String length must be greater than or equal to 3","fields":{"name":"String length must be greater than or equal to 3"}}`,
			useAuth:      true,
		},
		{
			name:         "CreateForumFailDecode",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"id":"abc",name":"Test"}`,
			responseCode: 405,
			response:     `{"message":"Failed to decode body"}`,
			useAuth:      true,
		},
		{
			name:         "CreateForumMinNameFail",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"forumId":1,"name":"Te"}`,
			responseCode: 400,
			response:     ``,
			useAuth:      true,
		},
		{
			name:         "CreateForumMaxNameFail",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"name":"Tsidofjsdoifjsdofijsdofijsdofisjdfoisj","forumId":1,"image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"String length must be less than or equal to 32","fields":{"name":"String length must be less than or equal to 32"}}`,
			useAuth:      true,
		},
		{
			name:         "CreateForumBadCharFail",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"forumId":1,"name":"Tsidofjs.øˆ∆ƒøˆ∆∂","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"Does not match pattern '^[a-zA-Z' ]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z' ]*$'"}}`,
			useAuth:      true,
		},
		{
			name:         "CreateForumNotAdmin",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"name":"Test","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 401,
			response:     `{"message":"Administrator access required"}`,
			useAuth:      false,
		},
		{
			name:         "CreateForum",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"name":"Test"}`,
			responseCode: 201,
			response:     `{"id":1,"name":"Test","ownerId":1,"description":"","icon":""}`,
			useAuth:      true,
		},
		{
			name:         "CreateForumTwo",
			path:         "/api/forum",
			method:       "POST",
			body:         `{"name":"Test Again"}`,
			responseCode: 201,
			response:     `{"id":2,"name":"Test Again","ownerId":1,"description":"","icon":""}`,
			useAuth:      true,
		},
		{
			name:         "GetForumInvalidForumId",
			path:         "/api/forum/invalid",
			method:       "GET",
			body:         "",
			responseCode: 400,
			response:     `{"message":"forumId argument is required: Invalid arguments provided"}`,
			useAuth:      false,
		},
		{
			name:         "GetForumNoResults",
			path:         "/api/forum/2",
			method:       "GET",
			body:         "",
			responseCode: 200,
			response:     ``,
			useAuth:      false,
		},
		{
			name:         "GetForum",
			path:         "/api/forum/1",
			method:       "GET",
			body:         "",
			responseCode: 200,
			response:     `{"id":1,"name":"Test","ownerId":1,"description":"","icon":""}`,
			useAuth:      false,
		},
		{
			name:         "NoTokenEditForum",
			path:         "/api/forum/invalid",
			method:       "PUT",
			body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 401,
			response:     `{"message":"Moderator access required"}`,
			useAuth:      false,
		},
		{
			name:         "InvalidEditForum",
			path:         "/api/forum/invalid",
			method:       "PUT",
			body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"forumId argument is required: Invalid arguments provided"}`,
			useAuth:      true,
		},
		{
			name:         "RequestFailEditForum",
			path:         "/api/forum/4",
			method:       "PUT",
			body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"Does not match pattern '^[a-zA-Z' ]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z' ]*$'"}}`,
			useAuth:      true,
		},

		{
			name:         "EditForum",
			path:         "/api/forum/2",
			method:       "PUT",
			body:         `{"id":1,"ownerId":1,"name":"TestTwo"}`,
			responseCode: 200,
			response:     `{"id":2,"name":"TestTwo","ownerId":1,"description":"","icon":""}`,
			useAuth:      true,
		},
		{
			name:         "NoContentEditForum",
			path:         "/api/forum/2",
			method:       "PUT",
			body:         `{"id":2,"status":0,"name":"TestTwo"}`,
			responseCode: 304,
			response:     ``,
			useAuth:      true,
		},
		{
			name:         "EditForumInvalid",
			path:         "/api/forum/1",
			method:       "PUT",
			body:         `{"id":"abc",name":"Test"}`,
			responseCode: 405,
			response:     `{"message":"Request error: Failed to decode body"}`,
			useAuth:      true,
		},
		{
			name:         "ListForum",
			path:         "/api/forum",
			method:       "GET",
			body:         "",
			responseCode: 200,
			response:     `[{"id":2,"name":"TestTwo","ownerId":1,"description":"","icon":""},{"id":1,"name":"Test","ownerId":1,"description":"","icon":""}]`,
			useAuth:      false,
		},
		{
			name:         "InvalidIdDeleteForum",
			path:         "/api/forum/{invalid}",
			method:       "DELETE",
			body:         "",
			responseCode: 400,
			response:     `{"message":"forumId argument is required: Invalid arguments provided"}`,
			useAuth:      true,
		},
		{
			name:         "RequestFailDeleteForum",
			path:         "/api/forum/3",
			method:       "DELETE",
			body:         "",
			responseCode: 304,
			response:     ``,
			useAuth:      true,
		},
		{
			name:         "DeleteForumNotLoggedIn",
			path:         "/api/forum/1",
			method:       "DELETE",
			body:         "",
			responseCode: 401,
			response:     `{"message":"Administrator access required"}`,
			useAuth:      false,
		},
		{
			name:         "DeleteForum",
			path:         "/api/forum/1",
			method:       "DELETE",
			body:         "",
			responseCode: 204,
			response:     ``,
			useAuth:      true,
		},
	}

	for _, test := range tests {
		doHttpTest(test, t)
	}
}
