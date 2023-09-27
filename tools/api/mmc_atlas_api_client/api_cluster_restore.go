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

type ClusterRestoreAPI interface {

	/*
		CreateRestoreJob Method for CreateRestoreJob

		Creates a restore job for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clusterName The name of the cluster to restore.
		@param projectId The ID of the project that contains the cluster to restore.
		@return CreateRestoreJobApiRequest
	*/
	/*
		CreateRestoreJob Method for CreateRestoreJob


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateRestoreJobApiParams - Parameters for the request
		@return CreateRestoreJobApiRequest
	*/
	CreateRestoreJob(ctx context.Context, args *CreateRestoreJobApiParams) CreateRestoreJobApiRequest

	// Interface only available internally
	createRestoreJobExecute(r CreateRestoreJobApiRequest) (*AtlasResponse, *http.Response, error)

	/*
		GetRestoreJob Get Restore Job

		Returns the restore job for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clusterName The name of the cluster to restore.
		@param projectId The ID of the project that contains the cluster to restore.
		@return GetRestoreJobApiRequest
	*/
	/*
		GetRestoreJob Get Restore Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetRestoreJobApiParams - Parameters for the request
		@return GetRestoreJobApiRequest
	*/
	GetRestoreJob(ctx context.Context, args *GetRestoreJobApiParams) GetRestoreJobApiRequest

	// Interface only available internally
	getRestoreJobExecute(r GetRestoreJobApiRequest) (*AtlasResponse, *http.Response, error)
}

// ClusterRestoreAPIService ClusterRestoreAPI service
type ClusterRestoreAPIService service

type CreateRestoreJobApiRequest struct {
	ctx                   context.Context
	ApiService            ClusterRestoreAPI
	clusterName           string
	projectId             string
	xMongoPublickey       *string
	xMongoPrivatekey      *string
	snapshotId            *string
	opLogTs               *int
	opLogInc              *int
	pointInTimeUtcSeconds *int
	targetClusterName     *string
	targetProjectId       *string
	deliveryType          *string
}

type CreateRestoreJobApiParams struct {
	ClusterName           string
	ProjectId             string
	XMongoPublickey       *string
	XMongoPrivatekey      *string
	SnapshotId            *string
	OpLogTs               *int
	OpLogInc              *int
	PointInTimeUtcSeconds *int
	TargetClusterName     *string
	TargetProjectId       *string
	DeliveryType          *string
}

func (a *ClusterRestoreAPIService) CreateRestoreJob(ctx context.Context, args *CreateRestoreJobApiParams) CreateRestoreJobApiRequest {
	return CreateRestoreJobApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		clusterName:           args.ClusterName,
		projectId:             args.ProjectId,
		xMongoPublickey:       args.XMongoPublickey,
		xMongoPrivatekey:      args.XMongoPrivatekey,
		snapshotId:            args.SnapshotId,
		opLogTs:               args.OpLogTs,
		opLogInc:              args.OpLogInc,
		pointInTimeUtcSeconds: args.PointInTimeUtcSeconds,
		targetClusterName:     args.TargetClusterName,
		targetProjectId:       args.TargetProjectId,
		deliveryType:          args.DeliveryType,
	}
}

// Public Key
func (r CreateRestoreJobApiRequest) XMongoPublickey(xMongoPublickey string) CreateRestoreJobApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r CreateRestoreJobApiRequest) XMongoPrivatekey(xMongoPrivatekey string) CreateRestoreJobApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

// The ID of the snapshot to restore.
func (r CreateRestoreJobApiRequest) SnapshotId(snapshotId string) CreateRestoreJobApiRequest {
	r.snapshotId = &snapshotId
	return r
}

// The timestamp of the oplog to restore.
func (r CreateRestoreJobApiRequest) OpLogTs(opLogTs int) CreateRestoreJobApiRequest {
	r.opLogTs = &opLogTs
	return r
}

// The increment of the oplog to restore.
func (r CreateRestoreJobApiRequest) OpLogInc(opLogInc int) CreateRestoreJobApiRequest {
	r.opLogInc = &opLogInc
	return r
}

// The point-in-time to restore, specified as a Unix timestamp in seconds.
func (r CreateRestoreJobApiRequest) PointInTimeUtcSeconds(pointInTimeUtcSeconds int) CreateRestoreJobApiRequest {
	r.pointInTimeUtcSeconds = &pointInTimeUtcSeconds
	return r
}

// The name of the target cluster.
func (r CreateRestoreJobApiRequest) TargetClusterName(targetClusterName string) CreateRestoreJobApiRequest {
	r.targetClusterName = &targetClusterName
	return r
}

// The ID of the project that contains the target cluster.
func (r CreateRestoreJobApiRequest) TargetProjectId(targetProjectId string) CreateRestoreJobApiRequest {
	r.targetProjectId = &targetProjectId
	return r
}

// The delivery type of the restore job.
func (r CreateRestoreJobApiRequest) DeliveryType(deliveryType string) CreateRestoreJobApiRequest {
	r.deliveryType = &deliveryType
	return r
}

func (r CreateRestoreJobApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.createRestoreJobExecute(r)
}

/*
CreateRestoreJob Method for CreateRestoreJob

Creates a restore job for the specified cluster.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterName The name of the cluster to restore.
 @param projectId The ID of the project that contains the cluster to restore.
 @return CreateRestoreJobApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *ClusterRestoreAPIService) createRestoreJobExecute(r CreateRestoreJobApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterRestoreAPIService.CreateRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster/{ClusterName}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"ClusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
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

	if r.snapshotId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SnapshotId", r.snapshotId, "")
	}
	if r.opLogTs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OpLogTs", r.opLogTs, "")
	}
	if r.opLogInc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OpLogInc", r.opLogInc, "")
	}
	if r.pointInTimeUtcSeconds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PointInTimeUtcSeconds", r.pointInTimeUtcSeconds, "")
	}
	if r.targetClusterName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "TargetClusterName", r.targetClusterName, "")
	}
	if r.targetProjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "TargetProjectId", r.targetProjectId, "")
	}
	if r.deliveryType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "DeliveryType", r.deliveryType, "")
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

type GetRestoreJobApiRequest struct {
	ctx              context.Context
	ApiService       ClusterRestoreAPI
	clusterName      string
	projectId        string
	xMongoPublickey  *string
	xMongoPrivatekey *string
	jobId            *string
}

type GetRestoreJobApiParams struct {
	ClusterName      string
	ProjectId        string
	XMongoPublickey  *string
	XMongoPrivatekey *string
	JobId            *string
}

func (a *ClusterRestoreAPIService) GetRestoreJob(ctx context.Context, args *GetRestoreJobApiParams) GetRestoreJobApiRequest {
	return GetRestoreJobApiRequest{
		ApiService:       a,
		ctx:              ctx,
		clusterName:      args.ClusterName,
		projectId:        args.ProjectId,
		xMongoPublickey:  args.XMongoPublickey,
		xMongoPrivatekey: args.XMongoPrivatekey,
		jobId:            args.JobId,
	}
}

// Public Key
func (r GetRestoreJobApiRequest) XMongoPublickey(xMongoPublickey string) GetRestoreJobApiRequest {
	r.xMongoPublickey = &xMongoPublickey
	return r
}

// Private Key
func (r GetRestoreJobApiRequest) XMongoPrivatekey(xMongoPrivatekey string) GetRestoreJobApiRequest {
	r.xMongoPrivatekey = &xMongoPrivatekey
	return r
}

// The JobID of the cluster to restore.
func (r GetRestoreJobApiRequest) JobId(jobId string) GetRestoreJobApiRequest {
	r.jobId = &jobId
	return r
}

func (r GetRestoreJobApiRequest) Execute() (*AtlasResponse, *http.Response, error) {
	return r.ApiService.getRestoreJobExecute(r)
}

/*
GetRestoreJob Get Restore Job

Returns the restore job for the specified cluster.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterName The name of the cluster to restore.
 @param projectId The ID of the project that contains the cluster to restore.
 @return GetRestoreJobApiRequest
*/

// Execute executes the request
//
//	@return AtlasResponse
func (a *ClusterRestoreAPIService) getRestoreJobExecute(r GetRestoreJobApiRequest) (*AtlasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AtlasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterRestoreAPIService.GetRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project/{ProjectId}/cluster/{ClusterName}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"ClusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
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
	if r.jobId == nil {
		return localVarReturnValue, nil, reportError("jobId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "JobId", r.jobId, "")
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
