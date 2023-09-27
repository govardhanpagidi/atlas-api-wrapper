package examples

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
	openapiclient "github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
)

func main() {
	url := "http://localhost:8080"

	apiClient, _ := openapiclient.NewClient(
		openapiclient.UseDigestAuth("mmc_user", "dharma"),
		openapiclient.UseBaseURL(url),
		openapiclient.UseDebug(true))

	DatabaseName := "437ac971-163e-445b-998b-0cfcb30e3bf8"
	HostName := "m-aws-22-09-23-16-15-36.iijwc.mongodb.net"

	request := openapiclient.DeleteDatabaseApiParams{
		DatabaseName: DatabaseName,
		HostName:     &HostName,
	}

	model, _, err := apiClient.DatabaseAPI.DeleteDatabase(context.Background(), &request).Execute()
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
