package api

import (
	"testing"
)

func TestAccountEndpoints(t *testing.T) {
	initializeServer(t)

	test := Endpoint{
		name:         "CreateAccountInvalidNameReq",
		path:         "/api/account",
		method:       "POST",
		body:         `{"accountID":1}`,
		responseCode: 400,
		response:     `{"message":"String length must be greater than or equal to 3","fields":{"name":"String length must be greater than or equal to 3"}}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "CreateAccountFailDecode",
		path:         "/api/account",
		method:       "POST",
		body:         `{"id":"abc",name":"Test"}`,
		responseCode: 405,
		response:     `{"message":"Failed to decode body"}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "CreateAccountMinNameFail",
		path:         "/api/account",
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
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "CreateAccountMaxNameFail",
		path:         "/api/account",
		method:       "POST",
		body:         `{"name":"Tsidofjsdoifjsdofijsdofijsdofisjdfoisj","accountID":1,"image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"String length must be less than or equal to 30","fields":{"name":"String length must be less than or equal to 30"}}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "CreateAccountBadCharFail",
		path:         "/api/account",
		method:       "POST",
		body:         `{"accountID":1,"name":"Tsidofjs.øˆ∆ƒøˆ∆∂","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"Does not match pattern '^[a-zA-Z]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z]*$'"}}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "CreateAccountNotAdmin",
		path:         "/api/account",
		method:       "POST",
		body:         `{"name":"Test","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "CreateAccount",
		path:         "/api/account",
		method:       "POST",
		body:         `{"name":"Test","status":10}`,
		responseCode: 201,
		response:     `{"id":82152,"name":"Test","charname":"","sharedplat":0,"password":"","status":10,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "GetAccountInvalidAccountID",
		path:         "/api/account/invalid",
		method:       "GET",
		body:         "",
		responseCode: 400,
		response:     `{"message":"accountID argument is required: Invalid arguments provided"}`,
		useAuth:      false,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "GetAccountNoResults",
		path:         "/api/account/2",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response:     nil,
		useAuth:      false,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "GetAccount",
		path:         "/api/account/1",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response:     `{"id":1,"name":"Shin","charname":"","sharedplat":0,"password":"","status":200,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""}`,
		useAuth:      false,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "NoTokenEditAccount",
		path:         "/api/account/invalid",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 401,
		response:     `{"message":"Moderator access required"}`,
		useAuth:      false,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "InvalidEditAccount",
		path:         "/api/account/invalid",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"accountID argument is required: Invalid arguments provided"}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "RequestFailEditAccount",
		path:         "/api/account/4",
		method:       "PUT",
		body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"Does not match pattern '^[a-zA-Z]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z]*$'"}}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "EditAccount",
		path:         "/api/account/2",
		method:       "PUT",
		body:         `{"id":1,"accountID":1,"name":"TestTwo","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "NoContentEditAccount",
		path:         "/api/account/2",
		method:       "PUT",
		body:         `{"id":2,"status":0,"name":"TestTwo"}`,
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "EditAccountInvalid",
		path:         "/api/account/1",
		method:       "PUT",
		body:         `{"id":"abc",name":"Test"}`,
		responseCode: 405,
		response:     `{"message":"Request error: Failed to decode body"}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "ListAccount",
		path:         "/api/account",
		method:       "GET",
		body:         "",
		responseCode: 200,
		response:     `[{"id":82152,"name":"Test","charname":"","sharedplat":0,"password":"","status":10,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""},{"id":1,"name":"Shin","charname":"","sharedplat":0,"password":"","status":200,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""}]`,
		useAuth:      false,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "InvalidIdDeleteAccount",
		path:         "/api/account/{invalid}",
		method:       "DELETE",
		body:         "",
		responseCode: 400,
		response:     `{"message":"accountID argument is required: Invalid arguments provided"}`,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "RequestFailDeleteAccount",
		path:         "/api/account/3",
		method:       "DELETE",
		body:         "",
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "DeleteAccountNotLoggedIn",
		path:         "/api/account/1",
		method:       "DELETE",
		body:         "",
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "DeleteAccount",
		path:         "/api/account/1",
		method:       "DELETE",
		body:         "",
		responseCode: 204,
		response:     ``,
		useAuth:      true,
	}
	doHTTPTest(test, t)
	test = Endpoint{
		name:         "DeleteAccountInvalidID",
		path:         "/api/account/5",
		method:       "DELETE",
		body:         "",
		responseCode: 304,
		response:     ``,
		useAuth:      true,
	}
	doHTTPTest(test, t)
}
