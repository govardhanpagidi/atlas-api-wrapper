// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

const (
	DefaultProfile    = "default"
	profileNamePrefix = "cfn/atlas/profile"
)

type Profile struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
	BaseURL    string `json:"BaseUrl,omitempty"`
}

// NewProfile creates a new Profile object from the given request and profile name
func NewProfile(req *http.Request, profileName *string) (*Profile, error) {
	// If the profileName is nil or empty, set it to the DefaultProfile constant
	if profileName == nil || *profileName == "" {
		profileName = aws.String(DefaultProfile)
	}

	// Create a new AWS session with the us-west-2 region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	// Create a new Secrets Manager client
	secretsManagerClient := secretsmanager.New(sess)

	// Get the secret value for the given profile name
	resp, err := secretsManagerClient.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: aws.String(fmt.Sprintf("%s/%s", profileNamePrefix, *profileName))})
	if err != nil {
		return nil, err
	}

	// Create a new Profile object
	profile := new(Profile)

	// Unmarshal the secret string into the Profile object
	err = json.Unmarshal([]byte(*resp.SecretString), &profile)
	if err != nil {
		return nil, err
	}

	// Return the Profile object
	return profile, nil
}

// NewBaseURL returns the MongoDB Atlas base URL from the environment variable or the Profile object
func (p *Profile) NewBaseURL() string {
	// If the MONGODB_ATLAS_BASE_URL environment variable is set, return its value
	if baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURL != "" {
		return baseURL
	}

	// Otherwise, return the BaseURL field of the Profile object
	return p.BaseURL
}

// NewPublicKey returns the MongoDB Atlas public key from the environment variable or the Profile object
func (p *Profile) NewPublicKey() string {
	// If the MONGODB_ATLAS_PUBLIC_KEY environment variable is set, return its value
	if publicKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY"); publicKey != "" {
		return publicKey
	}

	// Otherwise, return the PublicKey field of the Profile object
	return p.PublicKey
}

// NewPrivateKey returns the MongoDB Atlas private key from the environment variable or the Profile object
func (p *Profile) NewPrivateKey() string {
	// If the MONGODB_ATLAS_PRIVATE_KEY environment variable is set, return its value
	if privateKey := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY"); privateKey != "" {
		return privateKey
	}

	// Otherwise, return the PrivateKey field of the Profile object
	return p.PrivateKey
}

// AreKeysAvailable returns true if the MongoDB Atlas public and private keys are available in the Profile object or environment variables
func (p *Profile) AreKeysAvailable() bool {
	// If the MongoDB Atlas public key or private key is not available, return false
	if p.NewPublicKey() == "" || p.PrivateKey == "" {
		return false
	}

	// Otherwise, return true
	return true
}
