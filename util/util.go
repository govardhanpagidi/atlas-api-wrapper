package util

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	wrapperApi  = "mongodb-atlas-wrapper-api"
	envLogLevel = "LOG_LEVEL"
	debug       = "debug"
	warning     = "warningß"
)

var (
	toolName        = wrapperApi
	defaultLogLevel = "warning"
	userAgent       = fmt.Sprintf("%s/%s (%s;%s)", toolName, "version.Version", runtime.GOOS, runtime.GOARCH)
)

// EnsureAtlasRegion This takes either "us-east-1" or "US_EAST_1"
// and returns "US_EAST_1" -- i.e. a valid Atlas region
func EnsureAtlasRegion(region string) string {
	r := strings.ToUpper(strings.ReplaceAll(region, "-", "_"))
	log.Printf("EnsureAtlasRegion--- region:%s r:%s", region, r)
	return r
}

// NewMongoDBClient creates a new Client using apikeys
func NewMongoDBClient(ctx context.Context) (*mongodbatlas.Client, error) {
	publicKey := ctx.Value(constants.PubKey).(string)
	privateKey := ctx.Value(constants.PvtKey).(string)
	// setup a transport to handle digest
	log.Printf("CreateMongoDBClient--- publicKey:%s", publicKey)
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	opts := []mongodbatlas.ClientOpt{mongodbatlas.SetUserAgent(userAgent)}
	if baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURL != "" {
		opts = append(opts, mongodbatlas.SetBaseURL(baseURL))
	}

	return mongodbatlas.New(client, opts...)
}

func NewMongoDBSDKClient(publicKey string, privateKey string) (*admin.APIClient, error) {

	baseURL := admin.UseBaseURL("https://cloud.mongodb.com/")
	if baseURLString := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURLString != "" {
		baseURL = admin.UseBaseURL(baseURLString)
	}

	log.Printf("CreateMongoDBClient--- publicKey:%s", publicKey)
	client, err := admin.NewClient(admin.UseDigestAuth(publicKey, privateKey), baseURL)
	if err != nil {
		return nil, err
	}

	return client, err
}

func MongoDriverClient(userName string, password string, hostname string) (*mongo.Client, error) {
	connectionString := "mongodb+srv://" + userName + ":" + password + "@" + hostname

	// Set up a context and connect to the database
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return client, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
		return client, err
	}
	return client, nil
}

// SetupLogger is called by each resource handler to centrally
func SetupLogger(loggerPrefix string) {
	logr := logger.NewLogger(loggerPrefix)
	logger.SetOutput(logr.Writer())
	logger.SetLevel(getLogLevel())
}

func getLogLevel() logger.Level {
	levelString, exists := os.LookupEnv(envLogLevel)
	if !exists {
		_, _ = logger.Warnf("getLogLevel() Environment variable %s not found. Set it in template.yaml (defaultLogLevel=%s)", envLogLevel, defaultLogLevel)
		levelString = defaultLogLevel
	}
	switch levelString {
	case debug:
		return logger.DebugLevel
	case warning:
		return logger.WarningLevel

	default:
		return logger.DebugLevel
	}
}
