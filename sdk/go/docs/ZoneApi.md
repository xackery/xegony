# \ZoneApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZone**](ZoneApi.md#GetZone) | **Get** /zone/{zoneId} | Get information about a specific zone
[**ListZones**](ZoneApi.md#ListZones) | **Get** /zones | List all zones based on parameters


# **GetZone**
> Zone GetZone($zoneId)

Get information about a specific zone


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int32**| The id of the zone to retrieve | 

### Return type

[**Zone**](Zone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListZones**
> []Zones ListZones($limit)

List all zones based on parameters


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| How many zones to return at once (max 10) | [optional] 

### Return type

[**[]Zones**](Zones.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

