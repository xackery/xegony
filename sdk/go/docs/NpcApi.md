# \NpcApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNpc**](NpcApi.md#AddNpc) | **Post** /npc | Add a new npc to server
[**FindNpcsByName**](NpcApi.md#FindNpcsByName) | **Get** /npc/findByName | List all npcs that match search
[**FindNpcsByZoneId**](NpcApi.md#FindNpcsByZoneId) | **Get** /npc/findByZoneId | List all npcs found in provided zoneid
[**GetNpc**](NpcApi.md#GetNpc) | **Get** /npc/{npcId} | Get information about a specific npc
[**UpdateNpc**](NpcApi.md#UpdateNpc) | **Put** /npc | Update an existing npc


# **AddNpc**
> AddNpc($body)

Add a new npc to server




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Npc**](Npc.md)| Npc data | 

### Return type

void (empty response body)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNpcsByName**
> []Npcs FindNpcsByName($name, $limit)

List all npcs that match search


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **int32**| The name of the npc | 
 **limit** | **int32**| How many npcs to return at once (max 10) | [optional] 

### Return type

[**[]Npcs**](Npcs.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNpcsByZoneId**
> []Npcs FindNpcsByZoneId($zoneId, $limit)

List all npcs found in provided zoneid


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int32**| The id of the zone to retrieve npcs based on | 
 **limit** | **int32**| How many npcs to return at once (max 10) | [optional] 

### Return type

[**[]Npcs**](Npcs.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNpc**
> Npc GetNpc($npcId)

Get information about a specific npc


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **npcId** | **int32**| The id of the npc to retrieve | 

### Return type

[**Npc**](Npc.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNpc**
> UpdateNpc($body)

Update an existing npc




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Npc**](Npc.md)| Npc data | 

### Return type

void (empty response body)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

