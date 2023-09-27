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

type CloudBackupSnapshotAPI interface {

	/*
		CreateBackupSnapshot Create a cloud backup snapshot of a cluster

		Create a cloud backup snapshot of a cluster by project ID and cluster name

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param clusterName Cluster name
		@return CreateBackupSnapshotApiRequest
	*/
	/*
		CreateBackupSnapshot Create a cloud backup snapshot of a cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateBackupSnapshotApiParams - Parameters for the request
		@return CreateBackupSnapshotApiRequest
	*/
	CreateBackupSnapshot(ctx context.Context, args *CreateBackupSnapshotApiParams) CreateBackupSnapshotApiRequest

	// Interface only available internally
	createBackupSnapshotExecute(r CreateBackupSnapshotApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		GetAllBackupSnapshot Get all cloud backup snapshots of a cluster

		Get all cloud backup snapshots of a cluster by project ID and cluster name

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param clusterName Cluster name
		@return GetAllBackupSnapshotApiRequest
	*/
	/*
		GetAllBackupSnapshot Get all cloud backup snapshots of a cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAllBackupSnapshotApiParams - Parameters for the request
		@return GetAllBackupSnapshotApiRequest
	*/
	GetAllBackupSnapshot(ctx context.Context, args *GetAllBackupSnapshotApiParams) GetAllBackupSnapshotApiRequest

	// Interface only available internally
	getAllBackupSnapshotExecute(r GetAllBackupSnapshotApiRequest) (*AtlasResponse, *http.Response, error)
}

// CloudBackupSnapshotAPIService CloudBackupSnapshotAPI service
type CloudBackupSnapshotAPIService service

type CreateBackupSnapshotApiRequest struct {
	ctx              context.Context
	ApiService       CloudBackupSnapshotAPI
	projectId        string
	clusterName      string
	xMongoPublickey  *string
	xMongoPrivatekey *string
	description      *string
	retentionInDays  *string
}

type CreateBackupSnapshotApiParams struct {
	ProjectId        string
	ClusterName      string
	XMongoPublickey  *string
	XMongoPrivatekey *string
	Description      *string
	RetentionInDays  *string
}

func (a *CloudBackupSnapshotAPIService) CreateBackupSnapshot(ctx context.Context, args *CreateBackupSnapshotApiParams) CreateBackupSnapshotApiRequest {
	return CreateBackupSnapshotApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		clusterName:      args.ClusterName,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
		description:      args.Description,
		retentionInDays:  args.RetentionInDays,
	}
}

// Public Key
func (r CreateBackupSnapshotApiRequest) XMongoPublickey(xMongoPublickey string) CreateBackupSnapshotApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r CreateBackupSnapshotApiRequest) XMongoPrivatekey(xMongoPrivatekey string) CreateBackupSnapshotApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

// Description
func (r CreateBackupSnapshotApiRequest) Description(description string) CreateBackupSnapshotApiRequest {
	r.description = &description
	return r
}

// Retention in days
func (r CreateBackupSnapshotApiRequest) RetentionInDays(retentionInDays string) CreateBackupSnapshotApiRequest {
	r.retentionInDays = &retentionInDays
	return r
}

func (r CreateBackupSnapshotApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.createBackupSnapshotExecute(r)
}

/*
CreateBackupSnapshot Create a cloud backup snapshot of a cluster

Create a cloud backup snapshot of a cluster by project ID and cluster name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param clusterName Cluster name
 @return CreateBackupSnapshotApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *CloudBackupSnapshotAPIService) createBackupSnapshotExecute(r CreateBackupSnapshotApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupSnapshotAPIService.CreateBackupSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster/{ClusterName}/snapshot"
	localVarPath = strings.Replace(localVarPath, "{"+"ProjectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ClusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xMongoPublickey == nil {
		return localVarReturnValue, nil, reportError("xMongoPublickey is required and must be specified")
	}
	if r.xMongoPrivatekey == nil {
		return localVarReturnValue, nil, reportError("xMongoPrivatekey is required and must be specified")
	}

	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Description", r.description, "")
	} else {
		var defaultValue string = "\"\""
		r.description = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "Description", r.description, "")
	}
	if r.retentionInDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RetentionInDays", r.retentionInDays, "")
	} else {
		var defaultValue string = "\"\""
		r.retentionInDays = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "RetentionInDays", r.retentionInDays, "")
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

type GetAllBackupSnapshotApiRequest struct {
	ctx              context.Context
	ApiService       CloudBackupSnapshotAPI
	projectId        string
	clusterName      string
	xMongoPublickey  *string
	xMongoPrivatekey *string
}

type GetAllBackupSnapshotApiParams struct {
	ProjectId        string
	ClusterName      string
	XMongoPublickey  *string
	XMongoPrivatekey *string
}

func (a *CloudBackupSnapshotAPIService) GetAllBackupSnapshot(ctx context.Context, args *GetAllBackupSnapshotApiParams) GetAllBackupSnapshotApiRequest {
	return GetAllBackupSnapshotApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		clusterName:      args.ClusterName,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
	}
}

// Public Key
func (r GetAllBackupSnapshotApiRequest) XMongoPublickey(xMongoPublickey string) GetAllBackupSnapshotApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r GetAllBackupSnapshotApiRequest) XMongoPrivatekey(xMongoPrivatekey string) GetAllBackupSnapshotApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r GetAllBackupSnapshotApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.getAllBackupSnapshotExecute(r)
}

/*
GetAllBackupSnapshot Get all cloud backup snapshots of a cluster

Get all cloud backup snapshots of a cluster by project ID and cluster name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param clusterName Cluster name
 @return GetAllBackupSnapshotApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *CloudBackupSnapshotAPIService) getAllBackupSnapshotExecute(r GetAllBackupSnapshotApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupSnapshotAPIService.GetAllBackupSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster/{ClusterName}/snapshot"
	localVarPath = strings.Replace(localVarPath, "{"+"ProjectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ClusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

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
