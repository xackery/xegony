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
		name:         "CreateCharacterBadCharFail",
		path:         "/api/character",
		method:       "POST",
		body:         `{"accountID":1,"name":"Tsidofjs.øˆ∆ƒøˆ∆∂","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 400,
		response:     `{"message":"Does not match pattern '^[a-zA-Z]*$'","fields":{"name":"Does not match pattern '^[a-zA-Z]*$'"}}`,
		useAuth:      true,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateCharacterNotAdmin",
		path:         "/api/character",
		method:       "POST",
		body:         `{"name":"Test","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 401,
		response:     `{"message":"Administrator access required"}`,
		useAuth:      false,
	})
	doHTTPTest(t, Endpoint{
		name:         "CreateCharacter",
		path:         "/api/character",
		method:       "POST",
		body:         `{"name":"Test","accountID":1}`,
		responseCode: 201,
		response:     `{"Base":null,"Inventory":null,"id":1,"accountID":1,"name":"Test","lastName":"","title":"","suffix":"","zoneID":0,"zoneInstance":0,"y":0,"x":0,"z":0,"heading":0,"gender":0,"race":0,"class":0,"level":0,"deity":0,"birthday":0,"lastLogin":0,"timePlayed":0,"level2":0,"anon":0,"gm":0,"face":0,"hairColor":0,"hairStyle":0,"beard":0,"beardColor":0,"eyeColor1":0,"eyeColor2":0,"drakkinHeritage":0,"drakkinTattoo":0,"drakkinDetails":0,"abilityTimeSeconds":0,"abilityNumber":0,"abilityTimeMinutes":0,"abilityTimeHours":0,"exp":0,"aaPointsSpent":0,"aaExp":0,"aaPoints":0,"groupLeadershipExp":0,"raidLeadershipExp":0,"groupLeadershipPoints":0,"raidLeadershipPoints":0,"points":0,"curHp":0,"mana":0,"endurance":0,"intoxication":0,"str":0,"sta":0,"cha":0,"dex":0,"int":0,"agi":0,"wis":0,"zoneChangeCount":0,"toxicity":0,"hungerLevel":0,"thirstLevel":0,"abilityUp":0,"ldonPointsGuk":0,"ldonPointsMir":0,"ldonPointsMmc":0,"ldonPointsRuj":0,"ldonPointsTak":0,"ldonPointsAvailable":0,"tributeTimeRemaining":0,"careerTributePoints":0,"tributePoints":0,"tributeActive":0,"pvpStatus":0,"pvpKills":0,"pvpDeaths":0,"pvpCurrentPoints":0,"pvpCareerPoints":0,"pvpBestKillStreak":0,"pvpWorstDeathStreak":0,"pvpCurrentKillStreak":0,"pvp2":0,"pvpType":0,"showHelm":0,"groupAutoConsent":0,"raidAutoConsent":0,"guildAutoConsent":0,"leadershipExpOn":0,"RestTimer":0,"airRemaining":0,"autosplitEnabled":0,"lfp":0,"lfg":0,"mailkey":"","xtargets":0,"firstlogon":0,"eAaEffects":0,"ePercentToAa":0,"eExpendedAaSpent":0,"aaPointsSpentOld":0,"aaPointsOld":0,"eLastInvsnapshot":0}`,
		useAuth:      true,
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
		response:     `{"Base":null,"Inventory":null,"id":1,"accountID":1,"name":"Test","lastName":"","title":"","suffix":"","zoneID":0,"zoneInstance":0,"y":0,"x":0,"z":0,"heading":0,"gender":0,"race":0,"class":0,"level":0,"deity":0,"birthday":0,"lastLogin":0,"timePlayed":0,"level2":0,"anon":0,"gm":0,"face":0,"hairColor":0,"hairStyle":0,"beard":0,"beardColor":0,"eyeColor1":0,"eyeColor2":0,"drakkinHeritage":0,"drakkinTattoo":0,"drakkinDetails":0,"abilityTimeSeconds":0,"abilityNumber":0,"abilityTimeMinutes":0,"abilityTimeHours":0,"exp":0,"aaPointsSpent":0,"aaExp":0,"aaPoints":0,"groupLeadershipExp":0,"raidLeadershipExp":0,"groupLeadershipPoints":0,"raidLeadershipPoints":0,"points":0,"curHp":0,"mana":0,"endurance":0,"intoxication":0,"str":0,"sta":0,"cha":0,"dex":0,"int":0,"agi":0,"wis":0,"zoneChangeCount":0,"toxicity":0,"hungerLevel":0,"thirstLevel":0,"abilityUp":0,"ldonPointsGuk":0,"ldonPointsMir":0,"ldonPointsMmc":0,"ldonPointsRuj":0,"ldonPointsTak":0,"ldonPointsAvailable":0,"tributeTimeRemaining":0,"careerTributePoints":0,"tributePoints":0,"tributeActive":0,"pvpStatus":0,"pvpKills":0,"pvpDeaths":0,"pvpCurrentPoints":0,"pvpCareerPoints":0,"pvpBestKillStreak":0,"pvpWorstDeathStreak":0,"pvpCurrentKillStreak":0,"pvp2":0,"pvpType":0,"showHelm":0,"groupAutoConsent":0,"raidAutoConsent":0,"guildAutoConsent":0,"leadershipExpOn":0,"RestTimer":0,"airRemaining":0,"autosplitEnabled":0,"lfp":0,"lfg":0,"mailkey":"","xtargets":0,"firstlogon":0,"eAaEffects":0,"ePercentToAa":0,"eExpendedAaSpent":0,"aaPointsSpentOld":0,"aaPointsOld":0,"eLastInvsnapshot":0}`,
		useAuth:      false,
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
		name:         "EditCharacter",
		path:         "/api/character/1",
		method:       "PUT",
		body:         `{"id":1,"accountID":1,"name":"TestTwo","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 200,
		response:     `{"Base":null,"Inventory":null,"id":1,"accountID":1,"name":"TestTwo","lastName":"","title":"","suffix":"","zoneID":0,"zoneInstance":0,"y":0,"x":0,"z":0,"heading":0,"gender":0,"race":0,"class":0,"level":0,"deity":0,"birthday":0,"lastLogin":0,"timePlayed":0,"level2":0,"anon":0,"gm":0,"face":0,"hairColor":0,"hairStyle":0,"beard":0,"beardColor":0,"eyeColor1":0,"eyeColor2":0,"drakkinHeritage":0,"drakkinTattoo":0,"drakkinDetails":0,"abilityTimeSeconds":0,"abilityNumber":0,"abilityTimeMinutes":0,"abilityTimeHours":0,"exp":0,"aaPointsSpent":0,"aaExp":0,"aaPoints":0,"groupLeadershipExp":0,"raidLeadershipExp":0,"groupLeadershipPoints":0,"raidLeadershipPoints":0,"points":0,"curHp":0,"mana":0,"endurance":0,"intoxication":0,"str":0,"sta":0,"cha":0,"dex":0,"int":0,"agi":0,"wis":0,"zoneChangeCount":0,"toxicity":0,"hungerLevel":0,"thirstLevel":0,"abilityUp":0,"ldonPointsGuk":0,"ldonPointsMir":0,"ldonPointsMmc":0,"ldonPointsRuj":0,"ldonPointsTak":0,"ldonPointsAvailable":0,"tributeTimeRemaining":0,"careerTributePoints":0,"tributePoints":0,"tributeActive":0,"pvpStatus":0,"pvpKills":0,"pvpDeaths":0,"pvpCurrentPoints":0,"pvpCareerPoints":0,"pvpBestKillStreak":0,"pvpWorstDeathStreak":0,"pvpCurrentKillStreak":0,"pvp2":0,"pvpType":0,"showHelm":0,"groupAutoConsent":0,"raidAutoConsent":0,"guildAutoConsent":0,"leadershipExpOn":0,"RestTimer":0,"airRemaining":0,"autosplitEnabled":0,"lfp":0,"lfg":0,"mailkey":"","xtargets":0,"firstlogon":0,"eAaEffects":0,"ePercentToAa":0,"eExpendedAaSpent":0,"aaPointsSpentOld":0,"aaPointsOld":0,"eLastInvsnapshot":0}`,
		useAuth:      true,
	})
	//This should be a 304 / no content response
	doHTTPTest(t, Endpoint{
		name:         "NoContentEditCharacter",
		path:         "/api/character/1",
		method:       "PUT",
		body:         `{"id":1,"name":"TestTwo","image":"http://lfg.link/image.png","thumbnail":"http://lfg.link/thumbnail.png"}`,
		responseCode: 200,
		response:     `{"Base":null,"Inventory":null,"id":1,"accountID":0,"name":"TestTwo","lastName":"","title":"","suffix":"","zoneID":0,"zoneInstance":0,"y":0,"x":0,"z":0,"heading":0,"gender":0,"race":0,"class":0,"level":0,"deity":0,"birthday":0,"lastLogin":0,"timePlayed":0,"level2":0,"anon":0,"gm":0,"face":0,"hairColor":0,"hairStyle":0,"beard":0,"beardColor":0,"eyeColor1":0,"eyeColor2":0,"drakkinHeritage":0,"drakkinTattoo":0,"drakkinDetails":0,"abilityTimeSeconds":0,"abilityNumber":0,"abilityTimeMinutes":0,"abilityTimeHours":0,"exp":0,"aaPointsSpent":0,"aaExp":0,"aaPoints":0,"groupLeadershipExp":0,"raidLeadershipExp":0,"groupLeadershipPoints":0,"raidLeadershipPoints":0,"points":0,"curHp":0,"mana":0,"endurance":0,"intoxication":0,"str":0,"sta":0,"cha":0,"dex":0,"int":0,"agi":0,"wis":0,"zoneChangeCount":0,"toxicity":0,"hungerLevel":0,"thirstLevel":0,"abilityUp":0,"ldonPointsGuk":0,"ldonPointsMir":0,"ldonPointsMmc":0,"ldonPointsRuj":0,"ldonPointsTak":0,"ldonPointsAvailable":0,"tributeTimeRemaining":0,"careerTributePoints":0,"tributePoints":0,"tributeActive":0,"pvpStatus":0,"pvpKills":0,"pvpDeaths":0,"pvpCurrentPoints":0,"pvpCareerPoints":0,"pvpBestKillStreak":0,"pvpWorstDeathStreak":0,"pvpCurrentKillStreak":0,"pvp2":0,"pvpType":0,"showHelm":0,"groupAutoConsent":0,"raidAutoConsent":0,"guildAutoConsent":0,"leadershipExpOn":0,"RestTimer":0,"airRemaining":0,"autosplitEnabled":0,"lfp":0,"lfg":0,"mailkey":"","xtargets":0,"firstlogon":0,"eAaEffects":0,"ePercentToAa":0,"eExpendedAaSpent":0,"aaPointsSpentOld":0,"aaPointsOld":0,"eLastInvsnapshot":0}`,
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
		response:     `[{"Base":null,"Inventory":null,"id":1,"accountID":0,"name":"TestTwo","lastName":"","title":"","suffix":"","zoneID":0,"zoneInstance":0,"y":0,"x":0,"z":0,"heading":0,"gender":0,"race":0,"class":0,"level":0,"deity":0,"birthday":0,"lastLogin":0,"timePlayed":0,"level2":0,"anon":0,"gm":0,"face":0,"hairColor":0,"hairStyle":0,"beard":0,"beardColor":0,"eyeColor1":0,"eyeColor2":0,"drakkinHeritage":0,"drakkinTattoo":0,"drakkinDetails":0,"abilityTimeSeconds":0,"abilityNumber":0,"abilityTimeMinutes":0,"abilityTimeHours":0,"exp":0,"aaPointsSpent":0,"aaExp":0,"aaPoints":0,"groupLeadershipExp":0,"raidLeadershipExp":0,"groupLeadershipPoints":0,"raidLeadershipPoints":0,"points":0,"curHp":0,"mana":0,"endurance":0,"intoxication":0,"str":0,"sta":0,"cha":0,"dex":0,"int":0,"agi":0,"wis":0,"zoneChangeCount":0,"toxicity":0,"hungerLevel":0,"thirstLevel":0,"abilityUp":0,"ldonPointsGuk":0,"ldonPointsMir":0,"ldonPointsMmc":0,"ldonPointsRuj":0,"ldonPointsTak":0,"ldonPointsAvailable":0,"tributeTimeRemaining":0,"careerTributePoints":0,"tributePoints":0,"tributeActive":0,"pvpStatus":0,"pvpKills":0,"pvpDeaths":0,"pvpCurrentPoints":0,"pvpCareerPoints":0,"pvpBestKillStreak":0,"pvpWorstDeathStreak":0,"pvpCurrentKillStreak":0,"pvp2":0,"pvpType":0,"showHelm":0,"groupAutoConsent":0,"raidAutoConsent":0,"guildAutoConsent":0,"leadershipExpOn":0,"RestTimer":0,"airRemaining":0,"autosplitEnabled":0,"lfp":0,"lfg":0,"mailkey":"","xtargets":0,"firstlogon":0,"eAaEffects":0,"ePercentToAa":0,"eExpendedAaSpent":0,"aaPointsSpentOld":0,"aaPointsOld":0,"eLastInvsnapshot":0}]`,
		useAuth:      false,
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
