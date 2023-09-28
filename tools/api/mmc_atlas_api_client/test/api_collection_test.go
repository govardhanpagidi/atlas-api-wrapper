/*
MMC Atlas API Helper

Testing CollectionAPIService

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

func Test_mmc_atlas_api_client_CollectionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CollectionAPIService CreateCollection", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var databaseName string
		// TODO: create model
		var inputModel openapiclient.CollectionInputModel = openapiclient.CollectionInputModel{}

		resp, httpRes, err := apiClient.CollectionAPI.CreateCollection(context.Background(), databaseName, &inputModel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionAPIService DeleteCollection", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var databaseName string
		var collectionName string
		// var hostName string

		resp, httpRes, err := apiClient.CollectionAPI.DeleteCollection(context.Background(), databaseName, collectionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionAPIService ListCollection", func(t *testing.T) {
		// TODO this test is incomplete. You should add your own test code here.
		t.Skip("skip test") // remove to run test

		var databaseName string
		// var hostName string

		resp, httpRes, err := apiClient.CollectionAPI.ListCollection(context.Background(), databaseName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}