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

type DatabaseUserAPI interface {

	/*
		CreateDatabaseUser Create a database user

		Create a new database user with the specified name and email

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@return CreateDatabaseUserApiRequest
	*/
	/*
		CreateDatabaseUser Create a database user


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateDatabaseUserApiParams - Parameters for the request
		@return CreateDatabaseUserApiRequest
	*/
	CreateDatabaseUser(ctx context.Context, args *CreateDatabaseUserApiParams) CreateDatabaseUserApiRequest

	// Interface only available internally
	createDatabaseUserExecute(r CreateDatabaseUserApiRequest) (*DatabaseUserModel, *http.Response, error)

	/*
		DeleteDatabaseUser Delete a database user

		Delete a database user with the specified ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param username Username
		@return DeleteDatabaseUserApiRequest
	*/
	/*
		DeleteDatabaseUser Delete a database user


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteDatabaseUserApiParams - Parameters for the request
		@return DeleteDatabaseUserApiRequest
	*/
	DeleteDatabaseUser(ctx context.Context, args *DeleteDatabaseUserApiParams) DeleteDatabaseUserApiRequest

	// Interface only available internally
	deleteDatabaseUserExecute(r DeleteDatabaseUserApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		GetAllDatabaseUser Get all database users

		Get all database users

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@return GetAllDatabaseUserApiRequest
	*/
	/*
		GetAllDatabaseUser Get all database users


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAllDatabaseUserApiParams - Parameters for the request
		@return GetAllDatabaseUserApiRequest
	*/
	GetAllDatabaseUser(ctx context.Context, args *GetAllDatabaseUserApiParams) GetAllDatabaseUserApiRequest

	// Interface only available internally
	getAllDatabaseUserExecute(r GetAllDatabaseUserApiRequest) ([]DatabaseUserModel, *http.Response, error)

	/*
		GetDatabaseUser Get a database user

		Get a database user with the specified ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param username Username
		@return GetDatabaseUserApiRequest
	*/
	/*
		GetDatabaseUser Get a database user


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDatabaseUserApiParams - Parameters for the request
		@return GetDatabaseUserApiRequest
	*/
	GetDatabaseUser(ctx context.Context, args *GetDatabaseUserApiParams) GetDatabaseUserApiRequest

	// Interface only available internally
	getDatabaseUserExecute(r GetDatabaseUserApiRequest) (*http.Response, error)
}

// DatabaseUserAPIService DatabaseUserAPI service
type DatabaseUserAPIService service

type CreateDatabaseUserApiRequest struct {
	ctx              context.Context
	ApiService       DatabaseUserAPI
	projectId        string
	xMongoPublickey  *string
	xMongoPrivatekey *string
	inputModel       *DatabaseUserInputModel
}

type CreateDatabaseUserApiParams struct {
	ProjectId        string
	XMongoPublickey  *string
	XMongoPrivatekey *string
	InputModel       *DatabaseUserInputModel
}

func (a *DatabaseUserAPIService) CreateDatabaseUser(ctx context.Context, args *CreateDatabaseUserApiParams) CreateDatabaseUserApiRequest {
	return CreateDatabaseUserApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
		inputModel:       args.InputModel,
	}
}

// Public Key
func (r CreateDatabaseUserApiRequest) XMongoPublickey(xMongoPublickey string) CreateDatabaseUserApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r CreateDatabaseUserApiRequest) XMongoPrivatekey(xMongoPrivatekey string) CreateDatabaseUserApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r CreateDatabaseUserApiRequest) Execute() (*DatabaseUserModel, *http.Response, error) {
	return r.ApiService.createDatabaseUserExecute(r)
}

/*
CreateDatabaseUser Create a database user

Create a new database user with the specified name and email

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return CreateDatabaseUserApiRequest
*/

// Execute executes the request
//
//	@return DatabaseUserModel
func (a *DatabaseUserAPIService) createDatabaseUserExecute(r CreateDatabaseUserApiRequest) (*DatabaseUserModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseUserModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseUserAPIService.CreateDatabaseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/databaseUsers"
	localVarPath = strings.Replace(localVarPath, "{"+"ProjectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xMongoPublickey == nil {
		return localVarReturnValue, nil, reportError("xMongoPublickey is required and must be specified")
	}
	if r.xMongoPrivatekey == nil {
		return localVarReturnValue, nil, reportError("xMongoPrivatekey is required and must be specified")
	}
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-publickey", r.xMongoPublickey, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-privatekey", r.xMongoPrivatekey, "")
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

type DeleteDatabaseUserApiRequest struct {
	ctx              context.Context
	ApiService       DatabaseUserAPI
	projectId        string
	username         string
	xMongoPublickey  *string
	xMongoPrivatekey *string
}

type DeleteDatabaseUserApiParams struct {
	ProjectId        string
	Username         string
	XMongoPublickey  *string
	XMongoPrivatekey *string
}

func (a *DatabaseUserAPIService) DeleteDatabaseUser(ctx context.Context, args *DeleteDatabaseUserApiParams) DeleteDatabaseUserApiRequest {
	return DeleteDatabaseUserApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		username:         args.Username,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
	}
}

// Public Key
func (r DeleteDatabaseUserApiRequest) XMongoPublickey(xMongoPublickey string) DeleteDatabaseUserApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r DeleteDatabaseUserApiRequest) XMongoPrivatekey(xMongoPrivatekey string) DeleteDatabaseUserApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r DeleteDatabaseUserApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.deleteDatabaseUserExecute(r)
}

/*
DeleteDatabaseUser Delete a database user

Delete a database user with the specified ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param username Username
 @return DeleteDatabaseUserApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *DatabaseUserAPIService) deleteDatabaseUserExecute(r DeleteDatabaseUserApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseUserAPIService.DeleteDatabaseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/databaseUsers/{Username}"
	localVarPath = strings.Replace(localVarPath, "{"+"ProjectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"Username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xMongoPublickey == nil {
		return localVarReturnValue, nil, reportError("xMongoPublickey is required and must be specified")
	}
	if r.xMongoPrivatekey == nil {
		return localVarReturnValue, nil, reportError("xMongoPrivatekey is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-publickey", r.xMongoPublickey, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-privatekey", r.xMongoPrivatekey, "")
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

type GetAllDatabaseUserApiRequest struct {
	ctx              context.Context
	ApiService       DatabaseUserAPI
	projectId        string
	xMongoPublickey  *string
	xMongoPrivatekey *string
}

type GetAllDatabaseUserApiParams struct {
	ProjectId        string
	XMongoPublickey  *string
	XMongoPrivatekey *string
}

func (a *DatabaseUserAPIService) GetAllDatabaseUser(ctx context.Context, args *GetAllDatabaseUserApiParams) GetAllDatabaseUserApiRequest {
	return GetAllDatabaseUserApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
	}
}

// Public Key
func (r GetAllDatabaseUserApiRequest) XMongoPublickey(xMongoPublickey string) GetAllDatabaseUserApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r GetAllDatabaseUserApiRequest) XMongoPrivatekey(xMongoPrivatekey string) GetAllDatabaseUserApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r GetAllDatabaseUserApiRequest) Execute() ([]DatabaseUserModel, *http.Response, error) {
	return r.ApiService.getAllDatabaseUserExecute(r)
}

/*
GetAllDatabaseUser Get all database users

Get all database users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return GetAllDatabaseUserApiRequest
*/

// Execute executes the request
//
//	@return []DatabaseUserModel
func (a *DatabaseUserAPIService) getAllDatabaseUserExecute(r GetAllDatabaseUserApiRequest) ([]DatabaseUserModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []DatabaseUserModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseUserAPIService.GetAllDatabaseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/databaseUsers"
	localVarPath = strings.Replace(localVarPath, "{"+"ProjectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xMongoPublickey == nil {
		return localVarReturnValue, nil, reportError("xMongoPublickey is required and must be specified")
	}
	if r.xMongoPrivatekey == nil {
		return localVarReturnValue, nil, reportError("xMongoPrivatekey is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-publickey", r.xMongoPublickey, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-privatekey", r.xMongoPrivatekey, "")
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

type GetDatabaseUserApiRequest struct {
	ctx              context.Context
	ApiService       DatabaseUserAPI
	projectId        string
	username         string
	xMongoPublickey  *string
	xMongoPrivatekey *string
}

type GetDatabaseUserApiParams struct {
	ProjectId        string
	Username         string
	XMongoPublickey  *string
	XMongoPrivatekey *string
}

func (a *DatabaseUserAPIService) GetDatabaseUser(ctx context.Context, args *GetDatabaseUserApiParams) GetDatabaseUserApiRequest {
	return GetDatabaseUserApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		username:         args.Username,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
	}
}

// Public Key
func (r GetDatabaseUserApiRequest) XMongoPublickey(xMongoPublickey string) GetDatabaseUserApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r GetDatabaseUserApiRequest) XMongoPrivatekey(xMongoPrivatekey string) GetDatabaseUserApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r GetDatabaseUserApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.getDatabaseUserExecute(r)
}

/*
GetDatabaseUser Get a database user

Get a database user with the specified ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param username Username
 @return GetDatabaseUserApiRequest
*/

// Execute executes the request
func (a *DatabaseUserAPIService) getDatabaseUserExecute(r GetDatabaseUserApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseUserAPIService.GetDatabaseUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/databaseUsers/{Username}"
	localVarPath = strings.Replace(localVarPath, "{"+"ProjectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"Username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xMongoPublickey == nil {
		return nil, reportError("xMongoPublickey is required and must be specified")
	}
	if r.xMongoPrivatekey == nil {
		return nil, reportError("xMongoPrivatekey is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-publickey", r.xMongoPublickey, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-mongo-privatekey", r.xMongoPrivatekey, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	return localVarHTTPResponse, nil
}
