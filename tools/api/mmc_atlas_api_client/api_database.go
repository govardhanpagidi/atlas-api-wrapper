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

type DatabaseAPI interface {

	/*
		CreateDatabase Create a new database

		Create a new database with the specified name and owner

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return CreateDatabaseApiRequest
	*/
	/*
		CreateDatabase Create a new database


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateDatabaseApiParams - Parameters for the request
		@return CreateDatabaseApiRequest
	*/
	CreateDatabase(ctx context.Context, args *CreateDatabaseApiParams) CreateDatabaseApiRequest

	// Interface only available internally
	createDatabaseExecute(r CreateDatabaseApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		DeleteDatabase Delete a database

		Delete a database with the specified ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param databaseName databaseName
		@return DeleteDatabaseApiRequest
	*/
	/*
		DeleteDatabase Delete a database


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteDatabaseApiParams - Parameters for the request
		@return DeleteDatabaseApiRequest
	*/
	DeleteDatabase(ctx context.Context, args *DeleteDatabaseApiParams) DeleteDatabaseApiRequest

	// Interface only available internally
	deleteDatabaseExecute(r DeleteDatabaseApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		ReadAllDatabase lists all database

		read all the databases using hostname,username and password

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ReadAllDatabaseApiRequest
	*/
	/*
		ReadAllDatabase lists all database


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ReadAllDatabaseApiParams - Parameters for the request
		@return ReadAllDatabaseApiRequest
	*/
	ReadAllDatabase(ctx context.Context, args *ReadAllDatabaseApiParams) ReadAllDatabaseApiRequest

	// Interface only available internally
	readAllDatabaseExecute(r ReadAllDatabaseApiRequest) (*AtlasResponse, *http.Response, error)
}

// DatabaseAPIService DatabaseAPI service
type DatabaseAPIService service

type CreateDatabaseApiRequest struct {
	ctx        context.Context
	ApiService DatabaseAPI
	inputModel *DatabaseInputModel
}

type CreateDatabaseApiParams struct {
	InputModel *DatabaseInputModel
}

func (a *DatabaseAPIService) CreateDatabase(ctx context.Context, args *CreateDatabaseApiParams) CreateDatabaseApiRequest {
	return CreateDatabaseApiRequest{
		ApiService: a,
		ctx:        ctx,
		inputModel: args.InputModel,
	}
}

func (r CreateDatabaseApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.createDatabaseExecute(r)
}

/*
CreateDatabase Create a new database

Create a new database with the specified name and owner

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CreateDatabaseApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *DatabaseAPIService) createDatabaseExecute(r CreateDatabaseApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseAPIService.CreateDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database"

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

type DeleteDatabaseApiRequest struct {
	ctx          context.Context
	ApiService   DatabaseAPI
	databaseName string
	hostName     *string
}

type DeleteDatabaseApiParams struct {
	DatabaseName string
	HostName     *string
}

func (a *DatabaseAPIService) DeleteDatabase(ctx context.Context, args *DeleteDatabaseApiParams) DeleteDatabaseApiRequest {
	return DeleteDatabaseApiRequest{
		ApiService:   a,
		ctx:          ctx,
		databaseName: args.DatabaseName,
		hostName:     args.HostName,
	}
}

// hostName
func (r DeleteDatabaseApiRequest) HostName(hostName string) DeleteDatabaseApiRequest {
	r.hostName = &hostName
	return r
}

func (r DeleteDatabaseApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.deleteDatabaseExecute(r)
}

/*
DeleteDatabase Delete a database

Delete a database with the specified ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseName databaseName
 @return DeleteDatabaseApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *DatabaseAPIService) deleteDatabaseExecute(r DeleteDatabaseApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseAPIService.DeleteDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database/{DatabaseName}"
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

type ReadAllDatabaseApiRequest struct {
	ctx        context.Context
	ApiService DatabaseAPI
	hostName   *string
}

type ReadAllDatabaseApiParams struct {
	HostName *string
}

func (a *DatabaseAPIService) ReadAllDatabase(ctx context.Context, args *ReadAllDatabaseApiParams) ReadAllDatabaseApiRequest {
	return ReadAllDatabaseApiRequest{
		ApiService: a,
		ctx:        ctx,
		hostName:   args.HostName,
	}
}

// hostName
func (r ReadAllDatabaseApiRequest) HostName(hostName string) ReadAllDatabaseApiRequest {
	r.hostName = &hostName
	return r
}

func (r ReadAllDatabaseApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.readAllDatabaseExecute(r)
}

/*
ReadAllDatabase lists all database

read all the databases using hostname,username and password

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ReadAllDatabaseApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *DatabaseAPIService) readAllDatabaseExecute(r ReadAllDatabaseApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseAPIService.ReadAllDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/database"

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
