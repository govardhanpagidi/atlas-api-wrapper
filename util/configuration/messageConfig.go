package configuration

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/atlas-api-helper/util/constants"
)

// MessageConfig Singleton represents the singleton instance.
type MessageConfig struct {
	errorConfig map[string]Messages
}

var (
	instance *MessageConfig
	once     sync.Once
)

// GetInstance returns the singleton instance.
func GetInstance() *MessageConfig {
	once.Do(func() {
		instance = &MessageConfig{}
		instance.errorConfig = loadGlobalErrorConfig()
	})
	return instance
}

// GetData returns data from the singleton instance.
func (s *MessageConfig) GetData() map[string]Messages {
	return s.errorConfig
}

// GetConfig returns the global error configuration
func GetConfig() map[string]Messages {
	// Return the errorConfig field of the global instance of the MessageConfig struct
	return GetInstance().errorConfig
}

// loadGlobalErrorConfig loads the global error configuration from the constants.MessageConfigLocation file
func loadGlobalErrorConfig() map[string]Messages {
	// Declare the errorConfig variable as a map[string]Messages
	var errorConfig map[string]Messages

	// Read the contents of the constants.MessageConfigLocation file
	content, err := os.ReadFile(constants.MessageConfigLocation)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshal the contents of the file into the errorConfig variable
	err = json.Unmarshal(content, &errorConfig)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Return the errorConfig variable
	return errorConfig
}
