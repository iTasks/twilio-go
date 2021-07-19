# FlowsEngagementsStepsApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchStep**](FlowsEngagementsStepsApi.md#FetchStep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{Sid} | 
[**ListStep**](FlowsEngagementsStepsApi.md#ListStep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps | 



## FetchStep

> StudioV1FlowEngagementStep FetchStep(ctx, FlowSidEngagementSidSid)



Retrieve a Step.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**EngagementSid** | **string** | The SID of the Engagement with the Step to fetch.
**Sid** | **string** | The SID of the Step resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchStepParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1FlowEngagementStep**](StudioV1FlowEngagementStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStep

> ListStepResponse ListStep(ctx, FlowSidEngagementSidoptional)



Retrieve a list of all Steps for an Engagement.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to read.
**EngagementSid** | **string** | The SID of the Engagement with the Step to read.

### Other Parameters

Other parameters are passed through a pointer to a ListStepParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListStepResponse**](ListStepResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
