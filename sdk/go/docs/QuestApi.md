# \QuestApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuest**](QuestApi.md#GetQuest) | **Get** /quest/{questId} | Get information about a specific quest
[**ListQuests**](QuestApi.md#ListQuests) | **Get** /quests | List all quests based on parameters


# **GetQuest**
> Quest GetQuest($questId)

Get information about a specific quest


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questId** | **int32**| The id of the quest to retrieve | 

### Return type

[**Quest**](Quest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuests**
> []Quests ListQuests($limit)

List all quests based on parameters


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| How many quests to return at once (max 10) | [optional] 

### Return type

[**[]Quests**](Quests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

