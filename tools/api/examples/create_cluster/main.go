package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
	openapiclient "github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
)

func main() {

	url := "http://localhost:4000"
	apiClient := openapiclient.NewMMCClient(url)

	CloudProvider := "AWS"
	MongoDBVersion := "6.0.9"
	PrivateKey := "437ac971-163e-445b-998b-0cfcb30e3bf8"
	projectId := "5e8de3e1042f5b33ab81f33a"
	PublicKey := "hlmhviho"
	TshirtSize := "m"

	input := openapiclient.ClusterInputModel{
		CloudProvider:  &CloudProvider,
		MongoDBVersion: &MongoDBVersion,
		ProjectId:      &projectId,
		TshirtSize:     &TshirtSize,
	}
	request := openapiclient.CreateClusterApiParams{
		ProjectId:        projectId,
		InputModel:       &input,
		XMongoPublickey:  &PublicKey,
		XMongoPrivatekey: &PrivateKey,
	}

	model, _, _ := apiClient.ClusterAPI.CreateCluster(context.Background(), &request).Execute()
	// Marshal the request to JSON

	// Convert the struct to JSON
	// Print the JSON string
	printModel(model)

}

func printModel(model *mmc_atlas_api_client.ClusterModel) bool {
	jsonData, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Error:", err)
		return true
	}

	fmt.Println(string(jsonData))
	return false
}
