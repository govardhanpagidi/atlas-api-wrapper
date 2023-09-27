package examples

import (
	"context"
	"encoding/json"
	"fmt"

	openapiclient "github.com/govardhanpagidi/atlas-api-wrapper/mmc_atlas_api_client"
)

func main() {

	url := "http://localhost:4000"
	apiClient := openapiclient.NewMMCClient(url)

	PrivateKey := "437ac971-163e-445b-998b-0cfcb30e3bf8"
	projectId := "5e8de3e1042f5b33ab81f33a"
	PublicKey := "hlmhviho"

	request := openapiclient.GetAllClustersApiParams{
		ProjectId:        projectId,
		XMongoPublickey:  &PublicKey,
		XMongoPrivatekey: &PrivateKey,
	}

	model, _, errs := apiClient.ClusterAPI.GetAllClusters(context.Background(), &request).Execute()
	fmt.Printf("Error marshaling request to JSON: %v", errs)
	// Marshal the request to JSON

	printModel(model)

}

func printModel(model []openapiclient.ClusterModel) bool {
	jsonData, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Error:", err)
		return true
	}

	fmt.Println(string(jsonData))
	return false
}
