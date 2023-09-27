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

type CloudBackupScheduleAPI interface {

	/*
		GetCloudBackupSchedule Get the state of a cluster backup schedule

		Get the state of a cluster backup schedule by project ID and cluster name

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param clusterName Cluster name
		@return GetCloudBackupScheduleApiRequest
	*/
	/*
		GetCloudBackupSchedule Get the state of a cluster backup schedule


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCloudBackupScheduleApiParams - Parameters for the request
		@return GetCloudBackupScheduleApiRequest
	*/
	GetCloudBackupSchedule(ctx context.Context, args *GetCloudBackupScheduleApiParams) GetCloudBackupScheduleApiRequest

	// Interface only available internally
	getCloudBackupScheduleExecute(r GetCloudBackupScheduleApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		UpdateClusterBackupPolicy UpdateClusterBackupPolicy handles the PUT requests to update the cluster backup policy

		Update the cluster backup policy

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param projectId Project ID
		@param clusterName Cluster name
		@return UpdateClusterBackupPolicyApiRequest
	*/
	/*
		UpdateClusterBackupPolicy UpdateClusterBackupPolicy handles the PUT requests to update the cluster backup policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateClusterBackupPolicyApiParams - Parameters for the request
		@return UpdateClusterBackupPolicyApiRequest
	*/
	UpdateClusterBackupPolicy(ctx context.Context, args *UpdateClusterBackupPolicyApiParams) UpdateClusterBackupPolicyApiRequest

	// Interface only available internally
	updateClusterBackupPolicyExecute(r UpdateClusterBackupPolicyApiRequest) (*AtlasResponse, *http.Response, error)
}

// CloudBackupScheduleAPIService CloudBackupScheduleAPI service
type CloudBackupScheduleAPIService service

type GetCloudBackupScheduleApiRequest struct {
	ctx              context.Context
	ApiService       CloudBackupScheduleAPI
	projectId        string
	clusterName      string
	xMongoPublickey  *string
	xMongoPrivatekey *string
}

type GetCloudBackupScheduleApiParams struct {
	ProjectId        string
	ClusterName      string
	XMongoPublickey  *string
	XMongoPrivatekey *string
}

func (a *CloudBackupScheduleAPIService) GetCloudBackupSchedule(ctx context.Context, args *GetCloudBackupScheduleApiParams) GetCloudBackupScheduleApiRequest {
	return GetCloudBackupScheduleApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		clusterName:      args.ClusterName,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
	}
}

// Public Key
func (r GetCloudBackupScheduleApiRequest) XMongoPublickey(xMongoPublickey string) GetCloudBackupScheduleApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r GetCloudBackupScheduleApiRequest) XMongoPrivatekey(xMongoPrivatekey string) GetCloudBackupScheduleApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r GetCloudBackupScheduleApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.getCloudBackupScheduleExecute(r)
}

/*
GetCloudBackupSchedule Get the state of a cluster backup schedule

Get the state of a cluster backup schedule by project ID and cluster name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param clusterName Cluster name
 @return GetCloudBackupScheduleApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *CloudBackupScheduleAPIService) getCloudBackupScheduleExecute(r GetCloudBackupScheduleApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupScheduleAPIService.GetCloudBackupSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/clusters/{ClusterName}/backup/schedule"
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

type UpdateClusterBackupPolicyApiRequest struct {
	ctx              context.Context
	ApiService       CloudBackupScheduleAPI
	projectId        string
	clusterName      string
	xMongoPublickey  *string
	xMongoPrivatekey *string
	updateInputModel *CloudBackupScheduleModel
}

type UpdateClusterBackupPolicyApiParams struct {
	ProjectId        string
	ClusterName      string
	XMongoPublickey  *string
	XMongoPrivatekey *string
	UpdateInputModel *CloudBackupScheduleModel
}

func (a *CloudBackupScheduleAPIService) UpdateClusterBackupPolicy(ctx context.Context, args *UpdateClusterBackupPolicyApiParams) UpdateClusterBackupPolicyApiRequest {
	return UpdateClusterBackupPolicyApiRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        args.ProjectId,
		clusterName:      args.ClusterName,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
		updateInputModel: args.UpdateInputModel,
	}
}

// Public Key
func (r UpdateClusterBackupPolicyApiRequest) XMongoPublickey(xMongoPublickey string) UpdateClusterBackupPolicyApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r UpdateClusterBackupPolicyApiRequest) XMongoPrivatekey(xMongoPrivatekey string) UpdateClusterBackupPolicyApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

func (r UpdateClusterBackupPolicyApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.updateClusterBackupPolicyExecute(r)
}

/*
UpdateClusterBackupPolicy UpdateClusterBackupPolicy handles the PUT requests to update the cluster backup policy

Update the cluster backup policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param clusterName Cluster name
 @return UpdateClusterBackupPolicyApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *CloudBackupScheduleAPIService) updateClusterBackupPolicyExecute(r UpdateClusterBackupPolicyApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupScheduleAPIService.UpdateClusterBackupPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/clusters/{ClusterName}/backup/schedule"
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
