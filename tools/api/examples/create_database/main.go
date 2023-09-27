package examples

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
	openapiclient "github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
)

func main() {

	url := "http://localhost:4000"

	apiClient, _ := openapiclient.NewClient(
		openapiclient.UseDigestAuth("mmc_user", "dharma"),
		openapiclient.UseBaseURL(url),
		openapiclient.UseDebug(true))

	DatabaseName := "demo"
	CollectionName := "test"
	HostName := "m-aws-22-09-23-10-48-28.iijwc.mongodb.net"

	input := openapiclient.DatabaseInputModel{
		CollectionName: &CollectionName,
		DatabaseName:   &DatabaseName,
		HostName:       &HostName,
	}
	request := openapiclient.CreateDatabaseApiParams{
		InputModel: &input,
	}

	model, _, err := apiClient.DatabaseAPI.CreateDatabase(context.Background(), &request).Execute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
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
