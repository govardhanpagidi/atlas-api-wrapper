/*
MMC Atlas API Helper

Testing CloudBackupSnapshotAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mmc_atlas_api_client

import (
	"context"
	openapiclient "github.com/mongodb/atlas-api-wrapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_mmc_atlas_api_client_CloudBackupSnapshotAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudBackupSnapshotAPIService CreateBackupSnapshot", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var projectId string
		var clusterName string
		// var description string
		// var retentionInDays string

		resp, httpRes, err := apiClient.CloudBackupSnapshotAPI.CreateBackupSnapshot(context.Background(), projectId, clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudBackupSnapshotAPIService GetAllBackupSnapshot", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var projectId string
		var clusterName string

		resp, httpRes, err := apiClient.CloudBackupSnapshotAPI.GetAllBackupSnapshot(context.Background(), projectId, clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}