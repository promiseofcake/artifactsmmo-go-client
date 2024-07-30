/*
Artifacts API

 Artifacts is an API-based MMO game where you can manage 5 characters to explore, fight, gather resources, craft items and much more.  Website: https://artifactsmmo.com/  Documentation: https://docs.artifactsmmo.com/  OpenAPI Spec: https://api.artifactsmmo.com/openapi.json 

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type GrandExchangeAPI interface {

	/*
	GetAllGeItemsGeGet Get All Ge Items

	Fetch Grand Exchange items details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllGeItemsGeGetRequest
	*/
	GetAllGeItemsGeGet(ctx context.Context) ApiGetAllGeItemsGeGetRequest

	// GetAllGeItemsGeGetExecute executes the request
	//  @return DataPageGEItemSchema
	GetAllGeItemsGeGetExecute(r ApiGetAllGeItemsGeGetRequest) (*DataPageGEItemSchema, *http.Response, error)

	/*
	GetGeItemGeCodeGet Get Ge Item

	Retrieve the details of a Grand Exchange item.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param code The code of the item.
	@return ApiGetGeItemGeCodeGetRequest
	*/
	GetGeItemGeCodeGet(ctx context.Context, code string) ApiGetGeItemGeCodeGetRequest

	// GetGeItemGeCodeGetExecute executes the request
	//  @return GEItemResponseSchema
	GetGeItemGeCodeGetExecute(r ApiGetGeItemGeCodeGetRequest) (*GEItemResponseSchema, *http.Response, error)
}

// GrandExchangeAPIService GrandExchangeAPI service
type GrandExchangeAPIService service

type ApiGetAllGeItemsGeGetRequest struct {
	ctx context.Context
	ApiService GrandExchangeAPI
	page *int32
	size *int32
}

// Page number
func (r ApiGetAllGeItemsGeGetRequest) Page(page int32) ApiGetAllGeItemsGeGetRequest {
	r.page = &page
	return r
}

// Page size
func (r ApiGetAllGeItemsGeGetRequest) Size(size int32) ApiGetAllGeItemsGeGetRequest {
	r.size = &size
	return r
}

func (r ApiGetAllGeItemsGeGetRequest) Execute() (*DataPageGEItemSchema, *http.Response, error) {
	return r.ApiService.GetAllGeItemsGeGetExecute(r)
}

/*
GetAllGeItemsGeGet Get All Ge Items

Fetch Grand Exchange items details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllGeItemsGeGetRequest
*/
func (a *GrandExchangeAPIService) GetAllGeItemsGeGet(ctx context.Context) ApiGetAllGeItemsGeGetRequest {
	return ApiGetAllGeItemsGeGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DataPageGEItemSchema
func (a *GrandExchangeAPIService) GetAllGeItemsGeGetExecute(r ApiGetAllGeItemsGeGetRequest) (*DataPageGEItemSchema, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DataPageGEItemSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GrandExchangeAPIService.GetAllGeItemsGeGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ge/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "")
	} else {
		var defaultValue int32 = 50
		r.size = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetGeItemGeCodeGetRequest struct {
	ctx context.Context
	ApiService GrandExchangeAPI
	code string
}

func (r ApiGetGeItemGeCodeGetRequest) Execute() (*GEItemResponseSchema, *http.Response, error) {
	return r.ApiService.GetGeItemGeCodeGetExecute(r)
}

/*
GetGeItemGeCodeGet Get Ge Item

Retrieve the details of a Grand Exchange item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code The code of the item.
 @return ApiGetGeItemGeCodeGetRequest
*/
func (a *GrandExchangeAPIService) GetGeItemGeCodeGet(ctx context.Context, code string) ApiGetGeItemGeCodeGetRequest {
	return ApiGetGeItemGeCodeGetRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
//  @return GEItemResponseSchema
func (a *GrandExchangeAPIService) GetGeItemGeCodeGetExecute(r ApiGetGeItemGeCodeGetRequest) (*GEItemResponseSchema, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GEItemResponseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GrandExchangeAPIService.GetGeItemGeCodeGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ge/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}