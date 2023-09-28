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
	clusterName := "m-AWS-12-09-23-17-52-09-5e8de3e1042f5b33ab81f33a"
	mongoVersion := "7.0"
	input := openapiclient.ClusterUpdateInputModel{

		ProjectId:           &projectId,
		ClusterName:         &clusterName,
		MongoDBMajorVersion: &mongoVersion,
	}

	request := openapiclient.UpdateClusterApiParams{
		ProjectId:        projectId,
		XMongoPublickey:  &PublicKey,
		XMongoPrivatekey: &PrivateKey,
		UpdateInputModel: &input,
	}

	model, _, _ := apiClient.ClusterAPI.UpdateCluster(context.Background(), &request).Execute()
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