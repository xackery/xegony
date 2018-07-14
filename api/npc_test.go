package api

import (
	"net/http"
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

func TestNPCEndpoints(t *testing.T) {
	initializeServer(t)

	test := baloo.New("http://localhost:8081")
	test.Get("/api/npc").
		Expect(t).
		Status(http.StatusOK).
		Type("json").
		JSON("{\"npcs\":[{\"ID\":1008,\"aggroRadius\":55,\"agility\":156,\"ammoIDFile\":\"IT10\",\"armorClass\":311,\"attackCount\":-1,\"attackDelay\":32,\"attackSpeed\":-10,\"bodyTypeID\":1,\"charisma\":156,\"class\":{\"ID\":41,\"bit\":0,\"icon\":\"xa xa-shield\",\"name\":\"Shopkeeper\",\"shortName\":\"\"},\"classID\":41,\"cleanName\":\"Topala Xenem\",\"coldResistance\":18,\"corruptionResistance\":28,\"dexterity\":156,\"diseaseResistance\":18,\"exclude\":1,\"experience\":345515,\"findable\":1,\"fireResistance\":18,\"gender\":1,\"healScale\":100,\"helmTexture\":1,\"hitpoints\":5875,\"hpRegenRate\":12,\"intelligence\":156,\"lastName\":{\"String\":\"Bard Songs\",\"Valid\":true},\"level\":45,\"luclinBeard\":255,\"magicResistance\":18,\"manaRegenRate\":12,\"maximumDamage\":139,\"merchantID\":1008,\"minimumDamage\":36,\"name\":\"Topala_Xenem\",\"npcFactionID\":144,\"physicalResistance\":10,\"poisonResistance\":18,\"primaryMeleeTypeID\":28,\"race\":{\"ID\":71,\"female\":\"QCF\",\"icon\":\"xa xa-octopus\",\"male\":\"QCM\",\"name\":\"Human\"},\"raceID\":71,\"rangedTypeID\":7,\"runspeed\":1.325,\"scaleRate\":100,\"secondaryMeleeTypeID\":28,\"seeInvisibleUndead\":1,\"showName\":1,\"size\":6,\"specialAbilitiesRaw\":{\"String\":\"\",\"Valid\":true},\"spellScale\":100,\"stamina\":156,\"strength\":156,\"texture\":1,\"trackable\":1,\"trapTemplate\":{\"Int64\":0,\"Valid\":true},\"wisdom\":156,\"zone\":{\"ID\":1,\"MaxClip\":1600,\"MinClip\":50,\"canBind\":2,\"canCombat\":1,\"canLevitate\":1,\"castOutdoor\":1,\"expansion\":{\"ID\":50,\"name\":\"Unknown\",\"shortName\":\"UNK\"},\"expansionID\":50,\"fileName\":{\"String\":\"\",\"Valid\":true},\"fogBlue\":128,\"fogBlue2\":128,\"fogBlue3\":128,\"fogBlue4\":128,\"fogDensity\":0.33,\"fogMaxClip\":1600,\"fogMaxClip1\":2000,\"fogMaxClip2\":600,\"fogMaxClip3\":600,\"fogMaxClip4\":600,\"fogMinClip\":10,\"fogMinClip1\":450,\"fogMinClip2\":10,\"fogMinClip3\":10,\"fogMinClip4\":10,\"fogRed\":64,\"fogRed2\":64,\"fogRed3\":64,\"fogRed4\":64,\"gravity\":0.4,\"longName\":\"The Abysmal Sea\",\"mapFileName\":{\"String\":\"\",\"Valid\":false},\"minStatus\":250,\"modifier\":2,\"note\":{\"String\":\"\",\"Valid\":true},\"peqZone\":1,\"ruleset\":1,\"safeY\":-199,\"safeZ\":140,\"shortName\":{\"String\":\"abysmal\",\"Valid\":true},\"shutdownDelay\":5000,\"timeType\":2,\"type\":5,\"underworld\":-1000,\"walkSpeed\":0.4,\"zoneExpMultiplier\":1,\"zoneIDNumber\":279}}],\"page\":{\"isDescending\":0,\"limit\":1,\"offset\":0,\"orderBy\":\"id\",\"total\":1}}").
		Done()
	/*&NpcsResponse{
		Npcs: []*model.Npc{
			{
				ID:          1008,
				AggroRadius: 55,
			},
		},
		Page: &model.Page{
			Limit:   1,
			OrderBy: "id",
			Total:   1,
		},
	}*/

	/*test.Post("/api/account").
	Expect(t).
	JSON(`{"accountID":1}`).
	Status(400).
	Type("json").
	JSON(&model.ErrValidation{
		Message: "String length must be greater than or equal to 3",
		Reasons: map[string]string{
			"name": "String length must be greater than or equal to 3",
		},
	}).Done()*/

	/*doHTTPTest(t, Endpoint{
		name:         "CreateAccountInvalidNameReq",
		path:         "/api/account",
		method:       "POST",
		body:         `{"accountID":1}`,
		responseCode: 400,
		response:     `{"message":"String length must be greater than or equal to 3","fields":{"name":"String length must be greater than or equal to 3"}}`,
		useAuth:      true,
	})*/
	/*
		doHTTPTest(t, Endpoint{
			name:         "CreateAccountFailDecode",
			path:         "/api/account",
			method:       "POST",
			body:         `{"id":"abc",name":"Test"}`,
			responseCode: 405,
			response:     `{"message":"Failed to decode body"}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
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
		})
		doHTTPTest(t, Endpoint{
			name:         "CreateAccountMaxNameFail",
			path:         "/api/account",
			method:       "POST",
			body:         `{"name":"Tsidofjsdoifjsdofijsdofijsdofisjdfoisj","accountID":1,"image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"String length must be less than or equal to 30","fields":{"name":"String length must be less than or equal to 30"}}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "CreateAccountBadCharFail",
			path:         "/api/account",
			method:       "POST",
			body:         `{"accountID":1,"name":"Tsidofjs.øˆ∆ƒøˆ∆∂","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"Does not match pattern '^[a-zA-Z]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z]*$'"}}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "CreateAccountNotAdmin",
			path:         "/api/account",
			method:       "POST",
			body:         `{"name":"Test","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 401,
			response:     `{"message":"Administrator access required"}`,
			useAuth:      false,
		})
		doHTTPTest(t, Endpoint{
			name:         "CreateAccount",
			path:         "/api/account",
			method:       "POST",
			body:         `{"name":"Test","status":10}`,
			responseCode: 201,
			response:     `{"id":82152,"name":"Test","charname":"","sharedplat":0,"password":"","status":10,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "GetAccountInvalidAccountID",
			path:         "/api/account/invalid",
			method:       "GET",
			body:         "",
			responseCode: 400,
			response:     `{"message":"accountID argument is required: Invalid arguments provided"}`,
			useAuth:      false,
		})
		doHTTPTest(t, Endpoint{
			name:         "GetAccountNoResults",
			path:         "/api/account/2",
			method:       "GET",
			body:         "",
			responseCode: 200,
			response:     nil,
			useAuth:      false,
		})
		doHTTPTest(t, Endpoint{
			name:         "GetAccount",
			path:         "/api/account/1",
			method:       "GET",
			body:         "",
			responseCode: 200,
			response:     `{"id":1,"name":"Shin","charname":"","sharedplat":0,"password":"","status":200,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""}`,
			useAuth:      false,
		})
		doHTTPTest(t, Endpoint{
			name:         "NoTokenEditAccount",
			path:         "/api/account/invalid",
			method:       "PUT",
			body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 401,
			response:     `{"message":"Moderator access required"}`,
			useAuth:      false,
		})
		doHTTPTest(t, Endpoint{
			name:         "InvalidEditAccount",
			path:         "/api/account/invalid",
			method:       "PUT",
			body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"accountID argument is required: Invalid arguments provided"}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "RequestFailEditAccount",
			path:         "/api/account/4",
			method:       "PUT",
			body:         `{"id":1,"name":"Test2","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 400,
			response:     `{"message":"Does not match pattern '^[a-zA-Z]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z]*$'"}}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "EditAccount",
			path:         "/api/account/2",
			method:       "PUT",
			body:         `{"id":1,"accountID":1,"name":"TestTwo","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
			responseCode: 304,
			response:     ``,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "ErrNoContentEditAccount",
			path:         "/api/account/2",
			method:       "PUT",
			body:         `{"id":2,"status":0,"name":"TestTwo"}`,
			responseCode: 304,
			response:     ``,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "EditAccountInvalid",
			path:         "/api/account/1",
			method:       "PUT",
			body:         `{"id":"abc",name":"Test"}`,
			responseCode: 405,
			response:     `{"message":"Request error: Failed to decode body"}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "ListAccount",
			path:         "/api/account",
			method:       "GET",
			body:         "",
			responseCode: 200,
			response:     `[{"id":82152,"name":"Test","charname":"","sharedplat":0,"password":"","status":10,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""},{"id":1,"name":"Shin","charname":"","sharedplat":0,"password":"","status":200,"lsaccountID":{"Int64":0,"Valid":false},"gmspeed":0,"revoked":0,"karma":0,"miniloginIp":"","hideme":0,"rulesflag":0,"suspendeduntil":"0001-01-01T00:00:00Z","timeCreation":0,"expansion":0,"banReason":"","suspendReason":""}]`,
			useAuth:      false,
		})
		doHTTPTest(t, Endpoint{
			name:         "InvalidIdDeleteAccount",
			path:         "/api/account/{invalid}",
			method:       "DELETE",
			body:         "",
			responseCode: 400,
			response:     `{"message":"accountID argument is required: Invalid arguments provided"}`,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "RequestFailDeleteAccount",
			path:         "/api/account/3",
			method:       "DELETE",
			body:         "",
			responseCode: 304,
			response:     ``,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "DeleteAccountNotLoggedIn",
			path:         "/api/account/1",
			method:       "DELETE",
			body:         "",
			responseCode: 401,
			response:     `{"message":"Administrator access required"}`,
			useAuth:      false,
		})
		doHTTPTest(t, Endpoint{
			name:         "DeleteAccount",
			path:         "/api/account/1",
			method:       "DELETE",
			body:         "",
			responseCode: 204,
			response:     ``,
			useAuth:      true,
		})
		doHTTPTest(t, Endpoint{
			name:         "DeleteAccountInvalidID",
			path:         "/api/account/5",
			method:       "DELETE",
			body:         "",
			responseCode: 304,
			response:     ``,
			useAuth:      true,
		})*/
}
