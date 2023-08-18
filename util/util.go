package util

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util/logger"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const (
	envLogLevel = "LOG_LEVEL"
	debug       = "debug"
	warning     = "warningß"
)

var (
	defaultLogLevel = "warning"
)

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

func ToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value("traceID").(string)
	newArgs := append([]interface{}{traceId}, args...)
	logger.Debugf("[%s]"+format, newArgs...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value("traceID").(string)
	newArgs := append([]interface{}{traceId}, args...)
	logger.Warnf("[%s]"+format, newArgs...)
}

func Cast64(i *int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}
