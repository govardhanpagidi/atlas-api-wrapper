package configuration

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Singleton represents the singleton instance.
type ErrorMessageConfig struct {
	errorConfig map[string]ErrorConfig
}

var (
	instance *ErrorMessageConfig
	once     sync.Once
)

// GetInstance returns the singleton instance.
func GetInstance() *ErrorMessageConfig {
	once.Do(func() {
		instance = &ErrorMessageConfig{}
		instance.errorConfig = loadGlobalErrorConfig()
	})
	return instance
}

// GetData returns data from the singleton instance.
func (s *ErrorMessageConfig) GetData() map[string]ErrorConfig {
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
