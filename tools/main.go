package main

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	api  = "/api"
	port = "8080"
)

func main() {
	cfg := client.NewConfiguration()
	cfg.Host = "localhost:4000"
	clientSession := client.NewAPIClient(cfg)
	CloudProvider := "AWS"

	MongoDBVersion := "6.0.9"
	PrivateKey := "437ac971-163e-445b-998b-0cfcb30e3bf8"
	projectId := "5e8de3e1042f5b33ab81f33a"
	PublicKey := "hlmhviho"
	TshirtSize := "s"
	input := client.ClusterInputModel{
		CloudProvider: &CloudProvider,

		MongoDBVersion: &MongoDBVersion,
		PrivateKey:     &PrivateKey,
		ProjectId:      &projectId,
		PublicKey:      &PublicKey,
		TshirtSize:     &TshirtSize,
	}

	request := clientSession.ClusterAPI.CreateCluster(context.Background(), projectId)
	request.InputModel(input)
	// Marshal the request to JSON
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling request to JSON:", err)

	}

	// Print the JSON string
	jsonString := string(jsonBytes)
	fmt.Println("bbbbbb %s", jsonString)
	_, resp, err := request.Execute()
	if err != nil {
		println(resp)
		panic(err)
	}
}
