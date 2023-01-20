# \SmsApi

All URIs are relative to *https://ducksms.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmsSend**](SmsApi.md#SmsSend) | **Post** /api/v1/sms/send | Send Sms



## SmsSend

> PreviewSmsSend SmsSend(ctx, optional)

Send Sms

Create a new sms

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmsSendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SmsSendOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsSchema** | [**optional.Interface of SmsSchema**](SmsSchema.md)| Create a new sms | 

### Return type

[**PreviewSmsSend**](PreviewSmsSend.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

