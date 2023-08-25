package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/atlas-api-helper/handlers"
	"github.com/atlas-api-helper/resources/cluster"
	"github.com/atlas-api-helper/resources/collection"
	"github.com/atlas-api-helper/resources/database"
	database_user "github.com/atlas-api-helper/resources/databaseUser"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

var clusterName string = ""
var publicKey string = "nlbcisuz"
var privateKey string = "b37ea498-3950-4b8c-bfed-7779987d6195"
var projectId string = "64b6db1fe471e8514b8a59a6"
var connectionString string = ""
var databaseName string = "test"
var username string = "testUser"
var password string = "testPass"

type Config struct {
	ClusterName      string `json:"clusterName,omitempty"`
	PublicKey        string `json:"publicKey,omitempty"`
	PrivateKey       string `json:"privateKey,omitempty"`
	ProjectId        string `json:"projectId,omitempty"`
	ConnectionString string `json:"connectionString,omitempty"`
	DatabaseName     string `json:"databaseName,omitempty"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
}

func TestMain(m *testing.M) {
	var config Config

	// Read values from environment variables
	config.ClusterName = os.Getenv("CLUSTER_NAME")
	config.PublicKey = os.Getenv("PUBLIC_KEY")
	config.PrivateKey = os.Getenv("PRIVATE_KEY")
	config.ProjectId = os.Getenv("PROJECT_ID")
	config.ConnectionString = os.Getenv("CONNECTION_STRING")
	config.DatabaseName = os.Getenv("DATABASE_NAME")
	config.Username = os.Getenv("USERNAME")
	config.Password = os.Getenv("PASSWORD")

	// Load values from JSON file into tempConfig
	var tempConfig Config
	loadValuesFromFile("test_config.json", &tempConfig)

	// Assign values from tempConfig to config if they are empty
	if config.ClusterName == "" {
		config.ClusterName = tempConfig.ClusterName
	}
	if config.PublicKey == "" {
		config.PublicKey = tempConfig.PublicKey
	}
	if config.PrivateKey == "" {
		config.PrivateKey = tempConfig.PrivateKey
	}
	if config.ProjectId == "" {
		config.ProjectId = tempConfig.ProjectId
	}
	if config.ConnectionString == "" {
		config.ConnectionString = tempConfig.ConnectionString
	}
	if config.DatabaseName == "" {
		config.DatabaseName = tempConfig.DatabaseName
	}
	if config.Username == "" {
		config.Username = tempConfig.Username
	}
	if config.Password == "" {
		config.Password = tempConfig.Password
	}

	clusterName = config.ClusterName
	publicKey = config.PublicKey
	privateKey = config.PrivateKey
	projectId = config.ProjectId
	connectionString = config.ConnectionString
	databaseName = config.DatabaseName
	username = config.Username
	password = config.Password

	// Run the tests
	exitCode := m.Run()

	// Clean up your resources here if needed

	// Exit with the test exit code
	os.Exit(exitCode)
}

func loadValuesFromFile(filename string, config *Config) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Printf("Error decoding JSON from file %s: %v\n", filename, err)
		return
	}
}

func TestCreateCluster(t *testing.T) {
	// Set up mock input values
	requestBody := []byte(`{"publicKey": "nlbcisuz","privateKey": "b37ea498-3950-4b8c-bfed-7779987d6195", "tshirtSize": "S","CloudProvider":"AWS"}`)
	// Create a new request with the mock input values
	uri := "/project/" + projectId + "/clusterObj"
	println("*************************************************************************************************")

	println(uri)
	println("*************************************************************************************************")
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the CreateCluster handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/project/{projectId}/clusterObj", handlers.CreateCluster).Methods(http.MethodPost)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}

	var jsonObject map[string]interface{}

	// Unmarshal the JSON string into the jsonObject variable
	err = json.Unmarshal([]byte(responseBody), &jsonObject)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	clusterName = jsonObject["response"].(map[string]interface{})["name"].(string)

	println(clusterName)

	time.Sleep(45 * time.Second)

	client, err := util.NewMongoDBSDKClient(publicKey, privateKey)
	if err != nil {
		t.FailNow()
	}
	clusterObj, _, err := client.MultiCloudClustersApi.GetCluster(context.Background(), projectId, clusterName).Execute()
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
	if clusterObj.ConnectionStrings.StandardSrv != nil {
		parts := strings.SplitN(*clusterObj.ConnectionStrings.StandardSrv, "//", 2)
		connectionString = parts[1]
		println(connectionString)
	} else {
		t.Errorf("Failed to fetch clusterName")
		t.FailNow()
	}
}
func TestGetCluster(t *testing.T) {
	time.Sleep(30 * time.Second)
	// Create a new request with the mock input values
	req, err := http.NewRequest(http.MethodGet, "/project/"+projectId+"/cluster/"+clusterName+"/status?publicKey="+publicKey+"&privateKey="+privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the GetCluster handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/project/{projectId}/cluster/{clusterName}/status", handlers.GetCluster).Methods(http.MethodGet)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}

	responseBody := rr.Body.String()

	var response atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &response)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	prefix := "The Cluster status is:"
	if response.HttpStatusCode != 200 && !strings.HasPrefix(responseBody, prefix) {
		t.Errorf("Output dosent match expectation:%s", response.Response)
		t.FailNow()
	}
}
func TestGetAllCluster(t *testing.T) {
	// Set up mock input values
	//clusterName := "Cluster0"

	// Create a new request with the mock input values
	req, err := http.NewRequest("GET", "/project/"+projectId+"?publicKey="+publicKey+"&privateKey="+privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the GetAllCluster handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/project/{projectId}", handlers.GetAllClusters)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}

}

func TestCreateDatabaseUser(t *testing.T) {

	// Create a new request with the mock input values
	requestBody := database_user.InputModel{
		Username:   &username,
		Password:   &password,
		PublicKey:  &publicKey,
		PrivateKey: &privateKey,
	}
	requestBodyJson, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/databaseUsers/"+projectId, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the CreateDatabaseUser handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/databaseUsers/{projectId}", handlers.CreateDatabaseUser).Methods(http.MethodPost)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}

}

func TestGetDatabaseUser(t *testing.T) {
	// Set up mock input values
	time.Sleep(5 * time.Second)
	// Create a new request with the mock input values
	req, err := http.NewRequest("GET", "/databaseUsers/"+projectId+"/"+databaseName+"/"+username+"?publicKey="+publicKey+"&privateKey="+privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the GetDatabaseUser handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/databaseUsers/{projectId}/{databaseName}/{username}", handlers.GetDatabaseUser)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}

}

func TestGetAllDatabaseUser(t *testing.T) {
	// Set up mock input values
	time.Sleep(5 * time.Second)
	// Create a new request with the mock input values
	req, err := http.NewRequest("GET", "/databaseUsers/"+projectId+"?publicKey="+publicKey+"&privateKey="+privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the GetDatabaseUser handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/databaseUsers/{projectId}", handlers.GetAllDatabaseUser)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}

}

func TestCreateDatabase(t *testing.T) {
	time.Sleep(10 * time.Second)
	data := map[string]interface{}{
		"collectionName": "default",
		"databaseName":   databaseName,
		"hostName":       connectionString,
		"username":       username,
		"password":       password,
	}

	jsonBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	requestBody := string(jsonBody)

	// Create a new request with the mock input values
	req, err := http.NewRequest("POST", "/database", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the CreateDatabase handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/database", handlers.CreateDatabase).Methods(http.MethodPost)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
	}

	// Check that the response body is as expected
	expectedBody := fmt.Sprintf(configuration.GetConfig()[constants.DatabaseSuccess].Message, databaseName)

	responseBody := rr.Body.String()
	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if expectedBody != msg.Message {
		t.Errorf("Unexpected response body: got %v want %v", expectedBody, msg.Message)
		t.FailNow()
	}
}

func TestDeleteDatabase(t *testing.T) {
	time.Sleep(5 * time.Second)
	// Create a new request with the mock input values
	req, err := http.NewRequest("DELETE", "/database/"+databaseName+"?hostName="+connectionString+"&username="+username+"&password="+password, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the DeleteDatabase handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/database/{databaseName}", handlers.DeleteDatabase).Methods(http.MethodDelete)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}

	// Check that the response body is as expected
	expectedBody := fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteSuccess].Message, databaseName)

	var msg atlasresponse.AtlasRespone

	responseBody := rr.Body.String()
	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if expectedBody != msg.Message {
		t.Errorf("Unexpected response body: got %v want %v", expectedBody, msg.Message)
		t.FailNow()
	}
}

func TestCreateCollection(t *testing.T) {
	time.Sleep(5 * time.Second)
	// Set up mock input values
	collectionName := "test"
	collectionNames := []*string{&collectionName}
	data := map[string]interface{}{
		"collectionNames": collectionNames,
		"databaseName":    databaseName,
		"hostName":        connectionString,
		"username":        username,
		"password":        password,
	}

	jsonBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	requestBody := string(jsonBody)

	// Create a new request with the mock input values
	req, err := http.NewRequest("POST", "/collections/"+databaseName, strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the CreateCollection handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/collections/{databaseName}", handlers.CreateCollection)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}

}
func TestDeleteCollection(t *testing.T) {
	time.Sleep(5 * time.Second)
	// Create a new request with the mock input values
	req, err := http.NewRequest("DELETE", "/collections/"+databaseName+"/test?hostName="+connectionString+"&username="+username+"&password="+password, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the DeleteCollection handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/collections/{databaseName}/{collectionName}", handlers.DeleteCollection)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}
}

func TestDeleteDatabaseUser(t *testing.T) {
	time.Sleep(5 * time.Second)
	// Create a new request with the mock input values
	req, err := http.NewRequest("DELETE", "/databaseUsers/"+projectId+"/"+databaseName+"/"+username+"?publicKey="+publicKey+"&privateKey="+privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the DeleteDatabaseUser handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/databaseUsers/{projectId}/{databaseName}/{username}", handlers.DeleteDatabaseUser)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
	}

	responseBody := rr.Body.String()

	var msg atlasresponse.AtlasRespone

	// Unmarshal the response body into the map
	jsonErr := json.Unmarshal([]byte(responseBody), &msg)
	if jsonErr != nil {
		t.Errorf("Error unmarshaling JSON: %v", jsonErr)
		t.FailNow()
	}

	if msg.HttpStatusCode != 200 {
		t.Errorf("Output dosent match expectation:%s", msg.Message)
		t.FailNow()
	}

}

func TestDeleteCluster(t *testing.T) {
	time.Sleep(5 * time.Second)
	// Create a new request with the mock input values
	req, err := http.NewRequest("DELETE", "/project/"+projectId+"/cluster/"+clusterName+"?publicKey="+publicKey+"&privateKey="+privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()
	rr.Header().Set("Content-Type", "application/json")
	rr.Header().Set("Accept", "*/*")
	// Create a new router and register the DeleteCluster handler
	router := mux.NewRouter()
	router.Use(util.TraceIDMiddleware)
	router.HandleFunc("/project/{projectId}/cluster/{clusterName}", handlers.DeleteCluster).Methods(http.MethodDelete)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check that the response status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
		t.FailNow()
	}
}

//unit tests

var tshirtSize string = "s"
var cloudProvider string = "test"
var mongodbVersion string = "6.0.8"

func TestClusterCreateInputValidationUnit(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      nil,
		ClusterName:    nil,
		PrivateKey:     nil,
		PublicKey:      nil,
		TshirtSize:     nil,
		CloudProvider:  nil,
		MongoDBVersion: nil,
	}

	// Call the Read method with the mock client
	response := cluster.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestClusterCreateInputGetInvalidproject(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &clusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &privateKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &cloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestClusterCreateInputClusterCreateError(t *testing.T) {

	validCloudProvider := "aws"
	invalidClusterName := "@312321#21"
	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &invalidClusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &publicKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &validCloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 500 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestClusterCreateInputGetLoadConfig(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &clusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &publicKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &cloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestClusterReadInputValidationUnit(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      nil,
		ClusterName:    nil,
		PrivateKey:     nil,
		PublicKey:      nil,
		TshirtSize:     nil,
		CloudProvider:  nil,
		MongoDBVersion: nil,
	}

	// Call the Read method with the mock client
	response := cluster.Read(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestClusterReadInputGetInvalidproject(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &clusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &privateKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &cloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.Read(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestClusterReadInputClusterCreateError(t *testing.T) {

	validCloudProvider := "aws"
	invalidClusterName := "@312321#21"
	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &invalidClusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &publicKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &validCloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.Read(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestClusterDeleteInputValidationUnit(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      nil,
		ClusterName:    nil,
		PrivateKey:     nil,
		PublicKey:      nil,
		TshirtSize:     nil,
		CloudProvider:  nil,
		MongoDBVersion: nil,
	}

	// Call the Read method with the mock client
	response := cluster.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestClusterDeleteInputGetInvalidproject(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &clusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &privateKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &cloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestClusterDeleteInputClusterCreateError(t *testing.T) {

	validCloudProvider := "aws"
	invalidClusterName := "@312321#21"
	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &invalidClusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &publicKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &validCloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 500 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestClusterListInputValidationUnit(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      nil,
		ClusterName:    nil,
		PrivateKey:     nil,
		PublicKey:      nil,
		TshirtSize:     nil,
		CloudProvider:  nil,
		MongoDBVersion: nil,
	}

	// Call the Read method with the mock client
	response := cluster.List(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestClusterListInputGetInvalidproject(t *testing.T) {

	// Create the input model for testing
	inputModel := cluster.InputModel{
		ProjectId:      &projectId,
		ClusterName:    &clusterName,
		PrivateKey:     &privateKey,
		PublicKey:      &privateKey,
		TshirtSize:     &tshirtSize,
		CloudProvider:  &cloudProvider,
		MongoDBVersion: &mongodbVersion,
	}

	// Call the Read method with the mock client
	response := cluster.List(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("get non existing project passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseUserInputValidationUnit(t *testing.T) {

	// Create the input model for testing
	inputModel := database_user.InputModel{
		Username:   nil,
		Password:   nil,
		PublicKey:  nil,
		PrivateKey: nil,
		ProjectId:  nil,
	}

	// Call the Read method with the mock client
	response := database_user.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseUserInvalidUserNameUnit(t *testing.T) {

	invalidUser := "012ej21@3213"
	// Create the input model for testing
	inputModel := database_user.InputModel{
		Username:   &invalidUser,
		Password:   &password,
		PublicKey:  &publicKey,
		PrivateKey: &privateKey,
		ProjectId:  &projectId,
	}

	// Call the Read method with the mock client
	response := database_user.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 500 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseUserReadInputValidationUnit(t *testing.T) {

	// Create the input model for testing
	inputModel := database_user.InputModel{
		Username:   nil,
		Password:   nil,
		PublicKey:  nil,
		PrivateKey: nil,
		ProjectId:  nil,
	}

	// Call the Read method with the mock client
	response := database_user.Read(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseUserReadInvalidUserNameUnit(t *testing.T) {

	invalidUser := "012ej21@3213"
	// Create the input model for testing
	inputModel := database_user.InputModel{
		Username:   &invalidUser,
		Password:   &password,
		PublicKey:  &publicKey,
		PrivateKey: &privateKey,
		ProjectId:  &projectId,
	}

	// Call the Read method with the mock client
	response := database_user.Read(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 404 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseUserDeleteInputValidationUnit(t *testing.T) {

	// Create the input model for testing
	inputModel := database_user.InputModel{
		Username:   nil,
		Password:   nil,
		PublicKey:  nil,
		PrivateKey: nil,
		ProjectId:  nil,
	}

	// Call the Read method with the mock client
	response := database_user.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseUserDeleteInvalidUserNameUnit(t *testing.T) {

	invalidUser := "012ej21@3213"
	// Create the input model for testing
	inputModel := database_user.InputModel{
		Username:   &invalidUser,
		Password:   &password,
		PublicKey:  &publicKey,
		PrivateKey: &privateKey,
		ProjectId:  &projectId,
	}

	// Call the Read method with the mock client
	response := database_user.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 500 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestCollectionCreateWithInvalidInput(t *testing.T) {

	// Call the Read method with the mock client
	inputModel := collection.InputModel{
		CollectionNames: nil,
		DatabaseName:    nil,
		HostName:        nil,
		Username:        nil,
		Password:        nil,
	}

	response := collection.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestCollectionCreateWithInvalidUserNameAndPass(t *testing.T) {

	collectionNames := []*string{
		// Pointer to string values
		stringPtr("collection1"),
		stringPtr("collection2"),
		stringPtr("collection3"),
	}
	invalidHostName := "123@mongo.com"
	invalidUserName := "userName"
	invalidPass := "pass"

	// Call the Read method with the mock client
	inputModel := collection.InputModel{
		CollectionNames: collectionNames,
		DatabaseName:    &databaseName,
		HostName:        &invalidHostName,
		Username:        &invalidUserName,
		Password:        &invalidPass,
	}

	response := collection.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestCollectionDeleteWithInvalidInput(t *testing.T) {

	// Call the Read method with the mock client
	inputModel := collection.DeleteInputModel{
		CollectionName: nil,
		DatabaseName:   nil,
		HostName:       nil,
		Username:       nil,
		Password:       nil,
	}

	response := collection.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestCollectionDeleteWithInvalidUserNameAndPass(t *testing.T) {

	collectionName := "test"
	invalidHostName := "123@mongo.com"
	invalidUserName := "userName"
	invalidPass := "pass"

	// Call the Read method with the mock client
	inputModel := collection.DeleteInputModel{
		CollectionName: &collectionName,
		DatabaseName:   &databaseName,
		HostName:       &invalidHostName,
		Username:       &invalidUserName,
		Password:       &invalidPass,
	}

	response := collection.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseDeleteWithInvalidInput(t *testing.T) {

	// Call the Read method with the mock client
	inputModel := database.InputModel{
		CollectionName: nil,
		DatabaseName:   nil,
		HostName:       nil,
		Username:       nil,
		Password:       nil,
	}

	response := database.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseDeleteWithInvalidUserNameAndPass(t *testing.T) {

	collectionName := "test"
	invalidHostName := "123@mongo.com"
	invalidUserName := "userName"
	invalidPass := "pass"

	// Call the Read method with the mock client
	inputModel := database.InputModel{
		CollectionName: &collectionName,
		DatabaseName:   &databaseName,
		HostName:       &invalidHostName,
		Username:       &invalidUserName,
		Password:       &invalidPass,
	}

	response := database.Delete(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseCreateWithInvalidInput(t *testing.T) {

	// Call the Read method with the mock client
	inputModel := database.InputModel{
		CollectionName: nil,
		DatabaseName:   nil,
		HostName:       nil,
		Username:       nil,
		Password:       nil,
	}

	response := database.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func TestDatabaseCreateWithInvalidUserNameAndPass(t *testing.T) {

	collectionName := "test"
	invalidHostName := "123@mongo.com"
	invalidUserName := "userName"
	invalidPass := "pass"

	// Call the Read method with the mock client
	inputModel := database.InputModel{
		CollectionName: &collectionName,
		DatabaseName:   &databaseName,
		HostName:       &invalidHostName,
		Username:       &invalidUserName,
		Password:       &invalidPass,
	}

	response := database.Create(
		getContextWithTraceId(),
		&inputModel,
	)

	if response.HttpStatusCode != 400 {
		t.Error("Input validation passed instead of failing")
		t.FailNow()
	}
}

func getContextWithTraceId() context.Context {
	traceID := fmt.Sprintf("TraceID-%d", time.Now().UnixNano())
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.TraceID, traceID)
	return ctx
}

func stringPtr(s string) *string {
	return &s
}
