// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CollectionAPI interface {

	/*
		CreateCollection Create a new collection

		Create a new collection with the specified name and description

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param databaseName databaseName
		@return CreateCollectionApiRequest
	*/
	/*
		CreateCollection Create a new collection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateCollectionApiParams - Parameters for the request
		@return CreateCollectionApiRequest
	*/
	CreateCollection(ctx context.Context, args *CreateCollectionApiParams) CreateCollectionApiRequest

	// Interface only available internally
	createCollectionExecute(r CreateCollectionApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		DeleteCollection Delete a collection

		Delete a collection with the specified ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param databaseName databaseName
		@param collectionName collectionName
		@return DeleteCollectionApiRequest
	*/
	/*
		DeleteCollection Delete a collection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteCollectionApiParams - Parameters for the request
		@return DeleteCollectionApiRequest
	*/
	DeleteCollection(ctx context.Context, args *DeleteCollectionApiParams) DeleteCollectionApiRequest

	// Interface only available internally
	deleteCollectionExecute(r DeleteCollectionApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		ListCollection Lists all the collection

		Lists all the collections in the database

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param databaseName databaseName
		@return ListCollectionApiRequest
	*/
	/*
		ListCollection Lists all the collection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListCollectionApiParams - Parameters for the request
		@return ListCollectionApiRequest
	*/
	ListCollection(ctx context.Context, args *ListCollectionApiParams) ListCollectionApiRequest

	// Interface only available internally
	listCollectionExecute(r ListCollectionApiRequest) (*AtlasResponse, *http.Response, error)
}

// CollectionAPIService CollectionAPI service
type CollectionAPIService service

type CreateCollectionApiRequest struct {
	ctx          context.Context
	ApiService   CollectionAPI
	databaseName string
	inputModel   *CollectionInputModel
}

type CreateCollectionApiParams struct {
	DatabaseName string
	InputModel   *CollectionInputModel
}

func (a *CollectionAPIService) CreateCollection(ctx context.Context, args *CreateCollectionApiParams) CreateCollectionApiRequest {
	return CreateCollectionApiRequest{
		ApiService:   a,
		ctx:          ctx,
		databaseName: args.DatabaseName,
		inputModel:   args.InputModel,
	}
}

func (r CreateCollectionApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.createCollectionExecute(r)
}

/*
CreateCollection Create a new collection

Create a new collection with the specified name and description

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseName databaseName
 @return CreateCollectionApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *CollectionAPIService) createCollectionExecute(r CreateCollectionApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionAPIService.CreateCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database/{DatabaseName}/collections"
	localVarPath = strings.Replace(localVarPath, "{"+"DatabaseName"+"}", url.PathEscape(parameterValueToString(r.databaseName, "databaseName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inputModel == nil {
		return localVarReturnValue, nil, reportError("inputModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.inputModel
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {

		return localVarReturnValue, localVarHTTPResponse, err
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteCollectionApiRequest struct {
	ctx            context.Context
	ApiService     CollectionAPI
	databaseName   string
	collectionName string
	hostName       *string
}

type DeleteCollectionApiParams struct {
	DatabaseName   string
	CollectionName string
	HostName       *string
}

func (a *CollectionAPIService) DeleteCollection(ctx context.Context, args *DeleteCollectionApiParams) DeleteCollectionApiRequest {
	return DeleteCollectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		databaseName:   args.DatabaseName,
		collectionName: args.CollectionName,
		hostName:       args.HostName,
	}
}

// hostName
func (r DeleteCollectionApiRequest) HostName(hostName string) DeleteCollectionApiRequest {
	r.hostName = &hostName
	return r
}

func (r DeleteCollectionApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.deleteCollectionExecute(r)
}

/*
DeleteCollection Delete a collection

Delete a collection with the specified ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseName databaseName
 @param collectionName collectionName
 @return DeleteCollectionApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *CollectionAPIService) deleteCollectionExecute(r DeleteCollectionApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionAPIService.DeleteCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database/{DatabaseName}/collection/{CollectionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"DatabaseName"+"}", url.PathEscape(parameterValueToString(r.databaseName, "databaseName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"CollectionName"+"}", url.PathEscape(parameterValueToString(r.collectionName, "collectionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hostName == nil {
		return localVarReturnValue, nil, reportError("hostName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "HostName", r.hostName, "")
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {

		return localVarReturnValue, localVarHTTPResponse, err
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListCollectionApiRequest struct {
	ctx          context.Context
	ApiService   CollectionAPI
	databaseName string
	hostName     *string
}

type ListCollectionApiParams struct {
	DatabaseName string
	HostName     *string
}

func (a *CollectionAPIService) ListCollection(ctx context.Context, args *ListCollectionApiParams) ListCollectionApiRequest {
	return ListCollectionApiRequest{
		ApiService:   a,
		ctx:          ctx,
		databaseName: args.DatabaseName,
		hostName:     args.HostName,
	}
}

// hostName
func (r ListCollectionApiRequest) HostName(hostName string) ListCollectionApiRequest {
	r.hostName = &hostName
	return r
}

func (r ListCollectionApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.listCollectionExecute(r)
}

/*
ListCollection Lists all the collection

Lists all the collections in the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseName databaseName
 @return ListCollectionApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *CollectionAPIService) listCollectionExecute(r ListCollectionApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionAPIService.ListCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database/{DatabaseName}/collections"
	localVarPath = strings.Replace(localVarPath, "{"+"DatabaseName"+"}", url.PathEscape(parameterValueToString(r.databaseName, "databaseName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hostName == nil {
		return localVarReturnValue, nil, reportError("hostName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "HostName", r.hostName, "")
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {

		return localVarReturnValue, localVarHTTPResponse, err
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
