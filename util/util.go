package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	envLogLevel = "LOG_LEVEL"
	debug       = "debug"
	warning     = "warningß"
)

var (
	defaultLogLevel = "warning"
)

// NewMongoDBSDKClient creates a new MongoDB Atlas API client with the given public key and private key
func NewMongoDBSDKClient(publicKey string, privateKey string) (*admin.APIClient, error) {
	// Set the base URL of the client to the MongoDB Atlas API base URL
	baseURL := admin.UseBaseURL(constants.MongoBaseUrl)

	// If the MONGODB_ATLAS_BASE_URL environment variable is set, use it as the base URL of the client instead
	if baseURLString := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURLString != "" {
		baseURL = admin.UseBaseURL(baseURLString)
	}

	// Log the creation of the MongoDB Atlas API client with the given public key
	log.Printf("CreateMongoDBClient--- publicKey:%s", publicKey)

	// Create a new MongoDB Atlas API client with the given public key, private key, and base URL
	client, err := admin.NewClient(admin.UseDigestAuth(publicKey, privateKey), baseURL)
	if err != nil {
		return nil, err
	}

	// Return the new MongoDB Atlas API client and any error
	return client, err
}

// MongoDriverClient creates a new MongoDB client with the given username, password, and hostname
func MongoDriverClient(userName string, password string, hostname string) (*mongo.Client, error) {
	// Create a connection string with the given username, password, and hostname
	connectionString := fmt.Sprintf(constants.MongoDbURI, userName, password, hostname)

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

	// Return the new MongoDB client and any error
	return client, nil
}

// SetupLogger is called by each resource handler to centrally set up the logger
func SetupLogger(loggerPrefix string) {
	// Create a new logger with the given logger prefix
	logr := logger.NewLogger(loggerPrefix)

	// Set the output of the logger to the logger writer
	logger.SetOutput(logr.Writer())

	// Set the log level of the logger to the log level specified in the environment variable or the default log level
	logger.SetLevel(getLogLevel())
}

// getLogLevel gets the log level specified in the environment variable or the default log level
func getLogLevel() logger.Level {
	// Get the log level from the environment variable or use the default log level
	levelString, exists := os.LookupEnv(envLogLevel)
	if !exists {
		_, _ = logger.Warnf("getLogLevel() Environment variable %s not found. Set it in template.yaml (defaultLogLevel=%s)", envLogLevel, defaultLogLevel)
		levelString = defaultLogLevel
	}

	// Return the log level corresponding to the log level string
	switch levelString {
	case debug:
		return logger.DebugLevel
	case warning:
		return logger.WarningLevel
	default:
		return logger.DebugLevel
	}
}

// ToString converts a pointer to a string to a string or returns an empty string if the pointer is nil
func ToString(s *string) string {
	if s != nil {
		return *s
	}
	return constants.EmptyString
}

// ToStringSlice converts a slice of *string values to a formatted string.
func ToStringSlice(slice []*string) string {
	var strSlice []string

	// Iterate over the slice of *string values and append the string value of each non-nil element to the strSlice
	for _, s := range slice {
		if s != nil {
			strSlice = append(strSlice, *s)
		}
	}

	// Return the formatted string with the elements of the strSlice joined by commas and enclosed in square brackets
	return "[" + strings.Join(strSlice, ", ") + "]"
}

// Debugf logs a debug message with the given format and arguments, including the trace ID from the context
func Debugf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value(constants.TraceID).(string)

	// Append the trace ID to the beginning of the arguments slice and log the debug message with the new arguments
	newArgs := append([]interface{}{traceId}, args...)
	_, _ = logger.Debugf("[%s]"+format, newArgs...)
}

// Warnf logs a warning message with the given format and arguments, including the trace ID from the context
func Warnf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value(constants.TraceID).(string)

	// Append the trace ID to the beginning of the arguments slice and log the warning message with the new arguments
	newArgs := append([]interface{}{traceId}, args...)
	_, _ = logger.Warnf("[%s]"+format, newArgs...)
}

// Fatalf logs a fatal message with the given format and arguments, including the trace ID from the context
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value(constants.TraceID).(string)

	// Append the trace ID to the beginning of the arguments slice and log the warning message with the new arguments
	newArgs := append([]interface{}{traceId}, args...)
	_, _ = logger.Warnf("[%s]"+format, newArgs...)
}

// Cast64 converts a pointer to an int to a pointer to an int64
func Cast64(i *int) *int64 {
	// Use the `cast.ToInt64` function from the `github.com/spf13/cast` package to convert the pointer to an int64 pointer
	x := cast.ToInt64(&i)
	return &x
}

// TraceIDMiddleware generates a trace ID and adds it to the request context before calling the next handler
func TraceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a trace ID
		traceID := fmt.Sprintf("TraceID-%d", time.Now().UnixNano())

		// Add the trace ID to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, constants.TraceID, traceID)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func LoadConfigFromFile(targetFileName string, obj interface{}) error {
	// Start searching from the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Try to find the file in the current directory
	foundFilePath, err := findFile(currentDir, targetFileName)
	if err != nil {
		return err
	}

	// If not found, try to find the file in the parent directory
	if foundFilePath == "" {
		parentDir := filepath.Dir(currentDir)
		foundFilePath, err = findFile(parentDir, targetFileName)
		if err != nil {
			return err
		}
	}

	if foundFilePath != "" {
		// Read the content of the file
		content, err := os.ReadFile(foundFilePath)
		if err != nil {
			return err
		}

		// Unmarshal the data into the provided object
		err = json.Unmarshal(content, obj)
		if err != nil {
			return err
		}
	} else {
		return nil // File not found
	}

	return nil
}

func findFile(dirPath, targetFileName string) (string, error) {
	filePath := filepath.Join(dirPath, targetFileName)
	_, err := os.Stat(filePath)
	if err == nil {
		return filePath, nil
	} else if os.IsNotExist(err) {
		return "", nil
	} else {
		return "", err
	}
}
