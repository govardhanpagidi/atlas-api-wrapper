package configuration

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// ErrorMessageConfig Singleton represents the singleton instance.
type MessageConfig struct {
	errorConfig map[string]ErrorConfig
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
func (s *MessageConfig) GetData() map[string]ErrorConfig {
	return s.errorConfig
}

func GetConfig() map[string]ErrorConfig {
	return GetInstance().errorConfig
}

func loadGlobalErrorConfig() map[string]ErrorConfig {
	var errorConfig map[string]ErrorConfig
	content, err := os.ReadFile("./messageConfig.json")
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
