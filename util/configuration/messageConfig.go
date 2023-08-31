package configuration

import (
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	"log"
	"sync"
)

// MessageConfig Singleton represents the singleton instance.
type MessageConfig struct {
	msgConfig map[string]Messages
}

var (
	instance *MessageConfig
	once     sync.Once
)

// GetInstance returns the singleton instance.
func GetInstance() *MessageConfig {
	once.Do(func() {
		instance = &MessageConfig{}
		instance.msgConfig = loadGlobalMessageConfig()
	})
	return instance
}

// GetData returns data from the singleton instance.
func (s *MessageConfig) GetData() map[string]Messages {
	return s.msgConfig
}

// GetConfig returns the global error configuration
func GetConfig() map[string]Messages {
	// Return the msgConfig field of the global instance of the MessageConfig struct
	return GetInstance().msgConfig
}

func loadGlobalMessageConfig() map[string]Messages {
	var errorConfig map[string]Messages

	err := util.LoadConfigFromFile(constants.MessageConfigLocation, &errorConfig)
	if err != nil {
		log.Fatalf("Failed to load Message Config.")
	}
	if errorConfig == nil {
		log.Fatalf("Unable to load ConfigurationFile : %s", constants.MessageConfigLocation)
	}
	return errorConfig
}
