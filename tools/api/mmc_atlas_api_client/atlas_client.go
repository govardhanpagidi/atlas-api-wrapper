package mmc_atlas_api_client // import "go.mongodb.org/atlas-sdk/v20230201008/admin"

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

const (
	// DefaultCloudURL is default base URL for the services.
	DefaultCloudURL = "http://localhost:8080"
	// Version the version of the current API client inherited from.
	Version = ""
	// ClientName of the v2 API client.
	ClientName = "go-mmc-sdk"
)

// NewClient returns a new API Client.
func NewClient(modifiers ...ClientModifier) (*APIClient, error) {
	userAgent := fmt.Sprintf("%s/%s (%s;%s)", ClientName, Version, runtime.GOOS, runtime.GOARCH)
	defaultConfig := &Configuration{
		HTTPClient: http.DefaultClient,
		Servers: ServerConfigurations{ServerConfiguration{
			URL: DefaultCloudURL,
		},
		},
		UserAgent: userAgent,
	}
	for _, modifierFunction := range modifiers {
		err := modifierFunction(defaultConfig)
		if err != nil {
			return nil, err
		}
	}

	return NewAPIClient(defaultConfig), nil
}

// ClientModifiers lets you create function that controls configuration before creating client.
type ClientModifier func(*Configuration) error

func UseDigestAuth(username, password string) ClientModifier {
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	return func(c *Configuration) error {

		transport := &http.Transport{}
		// Create a new http.RoundTripper object that wraps the transport and adds the Basic Authentication header
		authTransport := &Transport{
			Transport: transport,
			BasicAuth: basicAuth,
		}

		// Use the authTransport object to make HTTP requests
		client := &http.Client{
			Transport: authTransport,
		}
		c.HTTPClient = client
		return nil
	}
}

// Advanced modifiers.

// UseHTTPClient set custom http client implementation.
func UseHTTPClient(client *http.Client) ClientModifier {
	return func(c *Configuration) error {
		c.HTTPClient = client
		return nil
	}
}

// UseDebug enable debug level logging.
func UseDebug(debug bool) ClientModifier {
	return func(c *Configuration) error {
		c.Debug = debug
		return nil
	}
}

// UseBaseURL set custom base url. If empty, default is used.
func UseBaseURL(baseURL string) ClientModifier {
	return func(c *Configuration) error {
		if baseURL == "" {
			baseURL = DefaultCloudURL
		}
		urlWithoutSuffix := strings.TrimSuffix(baseURL, "/")
		c.Servers = ServerConfigurations{ServerConfiguration{
			URL: urlWithoutSuffix,
		}}
		return nil
	}
}

// UseUserAgent set custom UserAgent header.
func UseUserAgent(userAgent string) ClientModifier {
	return func(c *Configuration) error {
		c.UserAgent = userAgent
		return nil
	}
}

// AsError checks if API returned known error type.
func AsError(err error) (*AtlasResponse, bool) {
	var openapiError *GenericOpenAPIError
	if ok := errors.As(err, &openapiError); !ok {
		return nil, false
	}
	errModel := openapiError.Model()
	return &errModel, true
}

// IsErrorCode returns true if the error contains the specific code.
func IsErrorCode(err error, code string) bool {
	mappedErr, _ := AsError(err)
	if mappedErr == nil {
		return false
	}
	return mappedErr.GetErrorCode() == code
}

type Transport struct {
	Transport *http.Transport
	BasicAuth string
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.BasicAuth)
	return t.Transport.RoundTrip(req)
}
