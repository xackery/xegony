# \RecipeApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecipe**](RecipeApi.md#GetRecipe) | **Get** /recipe/{recipeId} | Get information about a specific recipe
[**ListRecipes**](RecipeApi.md#ListRecipes) | **Get** /recipes | List all recipes based on parameters


# **GetRecipe**
> Recipe GetRecipe($recipeId)

Get information about a specific recipe


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recipeId** | **int32**| The id of the recipe to retrieve | 

### Return type

[**Recipe**](Recipe.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecipes**
> []Recipes ListRecipes($limit)

List all recipes based on parameters


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| How many recipes to return at once (max 10) | [optional] 

### Return type

[**[]Recipes**](Recipes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

