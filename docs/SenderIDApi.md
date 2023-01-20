# \SenderIDApi

All URIs are relative to *https://ducksms.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSender**](SenderIDApi.md#CreateSender) | **Post** /api/v1/senders | Create a Sender ID
[**DeleteSender**](SenderIDApi.md#DeleteSender) | **Delete** /api/v1/senders/{id} | Delete a Sender ID
[**GetSender**](SenderIDApi.md#GetSender) | **Get** /api/v1/senders/{id} | Get a single Sender ID
[**ListSender**](SenderIDApi.md#ListSender) | **Get** /api/v1/senders | List Sender ID
[**UpdateSender**](SenderIDApi.md#UpdateSender) | **Post** /api/v1/senders/{id} | Update a Sender ID



## CreateSender

> CreatedSender CreateSender(ctx, optional)

Create a Sender ID

Create a new sender id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSenderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sender** | [**optional.Interface of Sender**](Sender.md)| Create a new sender | 

### Return type

[**CreatedSender**](CreatedSender.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSender

> DeletedSender DeleteSender(ctx, id)

Delete a Sender ID

Delete an existing sender id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**DeletedSender**](DeletedSender.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSender

> GetSender GetSender(ctx, id)

Get a single Sender ID

Get details on a single sender id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**GetSender**](GetSender.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSender

> ListSender ListSender(ctx, optional)

List Sender ID

List all the senders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSenderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number | 
 **filterName** | **optional.String**| Filter by sender name | 
 **filterStatus** | **optional.String**| Filter by sender status | 

### Return type

[**ListSender**](ListSender.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSender

> UpdatedSender UpdateSender(ctx, id, optional)

Update a Sender ID

Update an existing sender id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***UpdateSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSenderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sender** | [**optional.Interface of Sender**](Sender.md)| Update an existing sender id | 

### Return type

[**UpdatedSender**](UpdatedSender.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

