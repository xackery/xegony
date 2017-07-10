# \CharacterApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCharacter**](CharacterApi.md#AddCharacter) | **Post** /character | Add a new character to server
[**DeleteCharacter**](CharacterApi.md#DeleteCharacter) | **Delete** /character | Deletes a character
[**FindCharactersByAccountId**](CharacterApi.md#FindCharactersByAccountId) | **Get** /character/findByAccountId | List all characters owned by provided account
[**FindCharactersByName**](CharacterApi.md#FindCharactersByName) | **Get** /character/findByName | List all characters that match search
[**GetCharacter**](CharacterApi.md#GetCharacter) | **Get** /character/{characterId} | Get information about a specific character
[**UpdateCharacter**](CharacterApi.md#UpdateCharacter) | **Put** /character | Update an existing character


# **AddCharacter**
> AddCharacter($body)

Add a new character to server




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Character**](Character.md)| Character data | 

### Return type

void (empty response body)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCharacter**
> DeleteCharacter($body, $apiKey)

Deletes a character




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Character**](Character.md)| Character data | 
 **apiKey** | **string**|  | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCharactersByAccountId**
> []Characters FindCharactersByAccountId($accountId, $limit)

List all characters owned by provided account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| The id of the account to retrieve characters based on | 
 **limit** | **int32**| How many characters to return at once (max 10) | [optional] 

### Return type

[**[]Characters**](Characters.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCharactersByName**
> []Characters FindCharactersByName($name, $limit)

List all characters that match search


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **int32**| The name of the character | 
 **limit** | **int32**| How many characters to return at once (max 10) | [optional] 

### Return type

[**[]Characters**](Characters.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharacter**
> Character GetCharacter($characterId)

Get information about a specific character


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| The id of the character to retrieve | 

### Return type

[**Character**](Character.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCharacter**
> UpdateCharacter($body)

Update an existing character




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Character**](Character.md)| Character data | 

### Return type

void (empty response body)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

