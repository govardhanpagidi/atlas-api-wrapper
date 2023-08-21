package configuration

import (
	"encoding/json"
	"github.com/atlas-api-helper/util/constants"
	"log"
	"os"
	"sync"
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

func GetConfig() map[string]Messages {
	return GetInstance().errorConfig
}

func loadGlobalErrorConfig() map[string]Messages {
	var errorConfig map[string]Messages
	content, err := os.ReadFile(constants.MessageConfigLocation)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &errorConfig)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return errorConfig
}
