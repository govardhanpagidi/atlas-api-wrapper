package main

import (
	"context"
	"encoding/json"
	"fmt"

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
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling request to JSON:", err)

	}

	// Print the JSON string
	jsonString := string(jsonBytes)
	fmt.Printf("bbbbbb %s", jsonString)
	model, _, errs := apiClient.ClusterAPI.CreateClusterWithParams(context.Background(), &request).Execute()
	fmt.Printf("Error marshaling request to JSON: %v", errs)
	// Marshal the request to JSON

	fmt.Println("bbeeebbbb %v", model)

}
