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

	Id := "Id"
	AutoExportEnabled := false
	ClusterId := "m-AWS-22-09-23-10-48-28-5e8de3e1042f5b33ab81f33a"

	request := openapiclient.CloudBackupScheduleModel{
		Id:                                &Id,
		AutoExportEnabled:                 &AutoExportEnabled,
		ClusterId:                         &ClusterId,
		CopySettings:                      nil,
		DeleteCopiedBackups:               nil,
		Export:                            nil,
		Links:                             nil,
		NextSnapshot:                      nil,
		Policies:                          nil,
		Profile:                           nil,
		ReferenceHourOfDay:                nil,
		ReferenceMinuteOfHour:             nil,
		RestoreWindowDays:                 nil,
		UpdateSnapshots:                   nil,
		UseOrgAndGroupNamesInExportPrefix: nil,
	}
	println("Create Backup Snapshot", request)
	updateInputModel := openapiclient.UpdateClusterBackupPolicyApiParams{
		ProjectId:        projectId,
		ClusterName:      "m-AWS-22-09-23-10-48-28-5e8de3e1042f5b33ab81f33a",
		XMongoPublickey:  &PublicKey,
		XMongoPrivatekey: &PrivateKey,
		UpdateInputModel: &request,
	}

	model, _, _ := apiClient.CloudBackupScheduleAPI.UpdateClusterBackupPolicy(context.Background(), &updateInputModel).Execute()

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
