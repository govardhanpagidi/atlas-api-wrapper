package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
	openapiclient "github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
)

func main() {

	url := "http://localhost:8080"
	apiClient := openapiclient.NewMMCClient(url)

	PrivateKey := "437ac971-163e-445b-998b-0cfcb30e3bf8"
	projectId := "5e8de3e1042f5b33ab81f33a"
	PublicKey := "hlmhviho"
	Description := "test"
	RetentionInDays := "3"

	request := openapiclient.CreateBackupSnapshotApiParams{
		Description:      &Description,
		ClusterName:      "m-AWS-22-09-23-10-48-28-5e8de3e1042f5b33ab81f33a",
		ProjectId:        projectId,
		RetentionInDays:  &RetentionInDays,
		XMongoPublickey:  &PublicKey,
		XMongoPrivatekey: &PrivateKey,
	}

	model, _, _ := apiClient.CloudBackupSnapshotAPI.CreateBackupSnapshot(context.Background(), &request).Execute()
	// Marshal the request to JSON

	// Convert the struct to JSON
	// Print the JSON string
	printModel(model)

}

func printModel(model *mmc_atlas_api_client.AtlasResponse) bool {
	jsonData, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Error:", err)
		return true
	}

	fmt.Println(string(jsonData))
	return false
}
