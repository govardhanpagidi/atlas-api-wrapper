package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/atlas-api-helper/handlers"
	database_user "github.com/atlas-api-helper/resources/databaseUser"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
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

func TestCreateCluster(t *testing.T) {
	// Set up mock input values
	requestBody := []byte(`{"publicKey": "nlbcisuz","privateKey": "b37ea498-3950-4b8c-bfed-7779987d6195", "tshirtSize": "S","CloudProvider":"AWS"}`)
	// Create a new request with the mock input values
	uri := "/project/" + projectId + "/cluster"
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
	router.HandleFunc("/project/{projectId}/cluster", handlers.CreateCluster).Methods(http.MethodPost)

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
	cluster, _, err := client.MultiCloudClustersApi.GetCluster(context.Background(), projectId, clusterName).Execute()
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
	if cluster.ConnectionStrings.StandardSrv != nil {
		parts := strings.SplitN(*cluster.ConnectionStrings.StandardSrv, "//", 2)
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
