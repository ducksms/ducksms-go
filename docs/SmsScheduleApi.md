# \SmsScheduleApi

All URIs are relative to *https://ducksms.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSmsSchedule**](SmsScheduleApi.md#CancelSmsSchedule) | **Delete** /api/v1/sms/scheduled/{id} | Cancel Sms Schedule
[**ListSmsSchedule**](SmsScheduleApi.md#ListSmsSchedule) | **Get** /api/v1/sms/scheduled | List Sms Schedule



## CancelSmsSchedule

> CancelSmsSchedule CancelSmsSchedule(ctx, id)

Cancel Sms Schedule

Cancel existing sms schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**CancelSmsSchedule**](CancelSmsSchedule.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsSchedule

> ListSmsSchedule ListSmsSchedule(ctx, optional)

List Sms Schedule

List all the sms schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSmsScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSmsScheduleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number | 
 **filterSenderName** | **optional.String**| Filter by sender id name | 
 **filterType** | **optional.String**| Filter by sms type | 
 **filterMessageType** | **optional.String**| Filter by sms message type | 
 **filterStatus** | **optional.String**| Filter by sms status | 

### Return type

[**ListSmsSchedule**](ListSmsSchedule.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

