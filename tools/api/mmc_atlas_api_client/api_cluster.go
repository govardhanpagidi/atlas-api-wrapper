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

type ClusterAPI interface {

	/*
		CreateCluster CreateCluster handles the POST requests to create the cluster with the provided TshirtSize

		create the cluster with the provided TshirtSize

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@return CreateClusterApiRequest
	*/
	/*
		CreateCluster CreateCluster handles the POST requests to create the cluster with the provided TshirtSize


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateClusterApiParams - Parameters for the request
		@return CreateClusterApiRequest
	*/
	CreateCluster(ctx context.Context, args *CreateClusterApiParams) CreateClusterApiRequest

	// Interface only available internally
	createClusterExecute(r CreateClusterApiRequest) (*ClusterModel, *http.Response, error)

	/*
		DeleteCluster Delete a cluster

		Delete a cluster by project ID and cluster name

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param clusterName Cluster name
		@return DeleteClusterApiRequest
	*/
	/*
		DeleteCluster Delete a cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteClusterApiParams - Parameters for the request
		@return DeleteClusterApiRequest
	*/
	DeleteCluster(ctx context.Context, args *DeleteClusterApiParams) DeleteClusterApiRequest

	// Interface only available internally
	deleteClusterExecute(r DeleteClusterApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		GetAllClusters Get all clusters

		Get all clusters along with their advanced configuration by project ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@return GetAllClustersApiRequest
	*/
	/*
		GetAllClusters Get all clusters


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAllClustersApiParams - Parameters for the request
		@return GetAllClustersApiRequest
	*/
	GetAllClusters(ctx context.Context, args *GetAllClustersApiParams) GetAllClustersApiRequest

	// Interface only available internally
	getAllClustersExecute(r GetAllClustersApiRequest) ([]ClusterModel, *http.Response, error)

	/*
		GetCluster Get the state of a cluster

		Get the state of a cluster by project ID and cluster name

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param clusterName Cluster name
		@return GetClusterApiRequest
	*/
	/*
		GetCluster Get the state of a cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterApiParams - Parameters for the request
		@return GetClusterApiRequest
	*/
	GetCluster(ctx context.Context, args *GetClusterApiParams) GetClusterApiRequest

	// Interface only available internally
	getClusterExecute(r GetClusterApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		UpdateCluster UpdateCluster handles the Put requests to create the provided MongoDbVersion and MongoDBMajorVersion

		Update the cluster with the provided MongoDbVersion and MongoDBMajorVersion

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@return UpdateClusterApiRequest
	*/
	/*
		UpdateCluster UpdateCluster handles the Put requests to create the provided MongoDbVersion and MongoDBMajorVersion


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateClusterApiParams - Parameters for the request
		@return UpdateClusterApiRequest
	*/
	UpdateCluster(ctx context.Context, args *UpdateClusterApiParams) UpdateClusterApiRequest

	// Interface only available internally
	updateClusterExecute(r UpdateClusterApiRequest) (*ClusterModel, *http.Response, error)
}

// ClusterAPIService ClusterAPI service
type ClusterAPIService service

type CreateClusterApiRequest struct {
	ctx              context.Context
	ApiService       ClusterAPI
	projectId        string
	xMongoPublickey  *string
	xMongoPrivatekey *string
	inputModel       *ClusterInputModel
}

type CreateClusterApiParams struct {
	ProjectId        string
	XMongoPublickey  *string
	XMongoPrivatekey *string
	InputModel       *ClusterInputModel
}

func (a *ClusterAPIService) CreateCluster(ctx context.Context, args *CreateClusterApiParams) CreateClusterApiRequest {
	return CreateClusterApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
		inputModel:       args.InputModel,
	}
}

// Public Key
func (r CreateClusterApiRequest) XMongoPublickey(xMongoPublickey string) CreateClusterApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r CreateClusterApiRequest) XMongoPrivatekey(xMongoPrivatekey string) CreateClusterApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r CreateClusterApiRequest) Execute() (*ClusterModel, *http.Response, error) {
	return r.ApiService.createClusterExecute(r)
}

/*
CreateCluster CreateCluster handles the POST requests to create the cluster with the provided TshirtSize

create the cluster with the provided TshirtSize

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return CreateClusterApiRequest
*/

// Execute executes the request
//
//	@return ClusterModel
func (a *ClusterAPIService) createClusterExecute(r CreateClusterApiRequest) (*ClusterModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterAPIService.CreateCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster"
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

type DeleteClusterApiRequest struct {
	ctx              context.Context
	ApiService       ClusterAPI
	projectId        string
	clusterName      string
	xMongoPublickey  *string
	xMongoPrivatekey *string
	retainBackup     *string
}

type DeleteClusterApiParams struct {
	ProjectId        string
	ClusterName      string
	XMongoPublickey  *string
	XMongoPrivatekey *string
	RetainBackup     *string
}

func (a *ClusterAPIService) DeleteCluster(ctx context.Context, args *DeleteClusterApiParams) DeleteClusterApiRequest {
	return DeleteClusterApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		clusterName:      args.ClusterName,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
		retainBackup:     args.RetainBackup,
	}
}

// Public Key
func (r DeleteClusterApiRequest) XMongoPublickey(xMongoPublickey string) DeleteClusterApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r DeleteClusterApiRequest) XMongoPrivatekey(xMongoPrivatekey string) DeleteClusterApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

// retainBackup
func (r DeleteClusterApiRequest) RetainBackup(retainBackup string) DeleteClusterApiRequest {
	r.retainBackup = &retainBackup
	return r
}

func (r DeleteClusterApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.deleteClusterExecute(r)
}

/*
DeleteCluster Delete a cluster

Delete a cluster by project ID and cluster name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param clusterName Cluster name
 @return DeleteClusterApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *ClusterAPIService) deleteClusterExecute(r DeleteClusterApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterAPIService.DeleteCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster/{ClusterName}"
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
	if r.retainBackup == nil {
		return localVarReturnValue, nil, reportError("retainBackup is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "RetainBackup", r.retainBackup, "")
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

type GetAllClustersApiRequest struct {
	ctx              context.Context
	ApiService       ClusterAPI
	projectId        string
	xMongoPublickey  *string
	xMongoPrivatekey *string
}

type GetAllClustersApiParams struct {
	ProjectId        string
	XMongoPublickey  *string
	XMongoPrivatekey *string
}

func (a *ClusterAPIService) GetAllClusters(ctx context.Context, args *GetAllClustersApiParams) GetAllClustersApiRequest {
	return GetAllClustersApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
	}
}

// Public Key
func (r GetAllClustersApiRequest) XMongoPublickey(xMongoPublickey string) GetAllClustersApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r GetAllClustersApiRequest) XMongoPrivatekey(xMongoPrivatekey string) GetAllClustersApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r GetAllClustersApiRequest) Execute() ([]ClusterModel, *http.Response, error) {
	return r.ApiService.getAllClustersExecute(r)
}

/*
GetAllClusters Get all clusters

Get all clusters along with their advanced configuration by project ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return GetAllClustersApiRequest
*/

// Execute executes the request
//
//	@return []ClusterModel
func (a *ClusterAPIService) getAllClustersExecute(r GetAllClustersApiRequest) ([]ClusterModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ClusterModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterAPIService.GetAllClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster"
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

type GetClusterApiRequest struct {
	ctx              context.Context
	ApiService       ClusterAPI
	projectId        string
	clusterName      string
	xMongoPublickey  *string
	xMongoPrivatekey *string
}

type GetClusterApiParams struct {
	ProjectId        string
	ClusterName      string
	XMongoPublickey  *string
	XMongoPrivatekey *string
}

func (a *ClusterAPIService) GetCluster(ctx context.Context, args *GetClusterApiParams) GetClusterApiRequest {
	return GetClusterApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		clusterName:      args.ClusterName,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
	}
}

// Public Key
func (r GetClusterApiRequest) XMongoPublickey(xMongoPublickey string) GetClusterApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r GetClusterApiRequest) XMongoPrivatekey(xMongoPrivatekey string) GetClusterApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r GetClusterApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.getClusterExecute(r)
}

/*
GetCluster Get the state of a cluster

Get the state of a cluster by project ID and cluster name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param clusterName Cluster name
 @return GetClusterApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *ClusterAPIService) getClusterExecute(r GetClusterApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterAPIService.GetCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster/{ClusterName}/status"
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

type UpdateClusterApiRequest struct {
	ctx              context.Context
	ApiService       ClusterAPI
	projectId        string
	xMongoPublickey  *string
	xMongoPrivatekey *string
	updateInputModel *ClusterUpdateInputModel
}

type UpdateClusterApiParams struct {
	ProjectId        string
	XMongoPublickey  *string
	XMongoPrivatekey *string
	UpdateInputModel *ClusterUpdateInputModel
}

func (a *ClusterAPIService) UpdateCluster(ctx context.Context, args *UpdateClusterApiParams) UpdateClusterApiRequest {
	return UpdateClusterApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
		updateInputModel: args.UpdateInputModel,
	}
}

// Public Key
func (r UpdateClusterApiRequest) XMongoPublickey(xMongoPublickey string) UpdateClusterApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r UpdateClusterApiRequest) XMongoPrivatekey(xMongoPrivatekey string) UpdateClusterApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r UpdateClusterApiRequest) Execute() (*ClusterModel, *http.Response, error) {
	return r.ApiService.updateClusterExecute(r)
}

/*
UpdateCluster UpdateCluster handles the Put requests to create the provided MongoDbVersion and MongoDBMajorVersion

Update the cluster with the provided MongoDbVersion and MongoDBMajorVersion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return UpdateClusterApiRequest
*/

// Execute executes the request
//
//	@return ClusterModel
func (a *ClusterAPIService) updateClusterExecute(r UpdateClusterApiRequest) (*ClusterModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterAPIService.UpdateCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster"
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
	if r.updateInputModel == nil {
		return localVarReturnValue, nil, reportError("updateInputModel is required and must be specified")
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
	localVarPostBody = r.updateInputModel
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
