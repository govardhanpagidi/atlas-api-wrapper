/*
MMC Atlas API Helper

Testing DatabaseUserAPIService

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

func Test_mmc_atlas_api_client_DatabaseUserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatabaseUserAPIService CreateDatabaseUser", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var projectId string
		// TODO: create model
		var inputModel openapiclient.DatabaseUserInputModel = openapiclient.DatabaseUserInputModel{}

		resp, httpRes, err := apiClient.DatabaseUserAPI.CreateDatabaseUser(context.Background(), projectId, &inputModel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseUserAPIService DeleteDatabaseUser", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var projectId string
		var username string

		resp, httpRes, err := apiClient.DatabaseUserAPI.DeleteDatabaseUser(context.Background(), projectId, username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseUserAPIService GetAllDatabaseUser", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.DatabaseUserAPI.GetAllDatabaseUser(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseUserAPIService GetDatabaseUser", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var projectId string
		var username string

		httpRes, err := apiClient.DatabaseUserAPI.GetDatabaseUser(context.Background(), projectId, username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}