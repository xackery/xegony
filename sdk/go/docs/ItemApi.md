# \ItemApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItem**](ItemApi.md#GetItem) | **Get** /item/{itemId} | Get information about a specific item
[**ListItems**](ItemApi.md#ListItems) | **Get** /items | List all items based on parameters


# **GetItem**
> Item GetItem($itemId)

Get information about a specific item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **int32**| The id of the item to retrieve | 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListItems**
> []Items ListItems($limit)

List all items based on parameters


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| How many items to return at once (max 10) | [optional] 

### Return type

[**[]Items**](Items.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

