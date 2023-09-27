package examples

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
	openapiclient "github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
)

func main1() {

	url := "http://localhost:4000"
	apiClient := openapiclient.NewMMCClient(url)

	PrivateKey := "437ac971-163e-445b-998b-0cfcb30e3bf8"
	projectId := "5e8de3e1042f5b33ab81f33a"
	PublicKey := "hlmhviho"
	Username := "dharma"
	request := openapiclient.DeleteDatabaseUserApiParams{
		ProjectId:        projectId,
		XMongoPublickey:  &PublicKey,
		XMongoPrivatekey: &PrivateKey,
		Username:         Username,
	}

	model, _, errs := apiClient.DatabaseUserAPI.DeleteDatabaseUser(context.Background(), &request).Execute()
	fmt.Printf("Error marshaling request to JSON: %v", errs)
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
