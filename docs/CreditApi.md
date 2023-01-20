# \CreditApi

All URIs are relative to *https://ducksms.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditBalance**](CreditApi.md#CreditBalance) | **Get** /api/v1/credits/balance | Credit Balance
[**CreditHistory**](CreditApi.md#CreditHistory) | **Get** /api/v1/credits/history | Credit History



## CreditBalance

> CreditBalance CreditBalance(ctx, )

Credit Balance

Get available credit balance

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CreditBalance**](CreditBalance.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditHistory

> CreditHistory CreditHistory(ctx, optional)

Credit History

Get all credit history

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreditHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreditHistoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number | 
 **filterType** | **optional.String**| Filter by credit type | 
 **filterSmsSmsid** | **optional.Int32**| Filter by sms id | 

### Return type

[**CreditHistory**](CreditHistory.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

