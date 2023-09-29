# SafeListNumbersApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSafelist**](SafeListNumbersApi.md#CreateSafelist) | **Post** /2010-04-01/SafeList/Numbers.json | 
[**DeleteSafelist**](SafeListNumbersApi.md#DeleteSafelist) | **Delete** /2010-04-01/SafeList/Numbers.json | 
[**FetchSafelist**](SafeListNumbersApi.md#FetchSafelist) | **Get** /2010-04-01/SafeList/Numbers.json | 



## CreateSafelist

> ApiV2010Safelist CreateSafelist(ctx, optional)



Add a new phone number to SafeList.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number to be added in SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Return type

[**ApiV2010Safelist**](ApiV2010Safelist.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSafelist

> DeleteSafelist(ctx, optional)



Remove a phone number from SafeList.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number to be removed from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSafelist

> ApiV2010Safelist FetchSafelist(ctx, optional)



Check if a phone number exists in SafeList.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number to be fetched from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Return type

[**ApiV2010Safelist**](ApiV2010Safelist.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
