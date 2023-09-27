package examples

import (
	"context"
	"fmt"

	openapiclient "github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
)

func main1() {

	url := "http://localhost:4000"
	apiClient := openapiclient.NewMMCClient(url)

	PrivateKey := "437ac971-163e-445b-998b-0cfcb30e3bf8"
	projectId := "5e8de3e1042f5b33ab81f33a"
	PublicKey := "hlmhviho"
	Username := "dharma"

	request := openapiclient.GetDatabaseUserApiParams{
		ProjectId:        projectId,
		XMongoPublickey:  &PublicKey,
		XMongoPrivatekey: &PrivateKey,
		Username:         Username,
	}

	model, errs := apiClient.DatabaseUserAPI.GetDatabaseUser(context.Background(), &request).Execute()
	fmt.Printf("Error marshaling request to JSON: %v", errs)
	// Marshal the request to JSON

	fmt.Println("bbeeebbbb %v", model)

}
