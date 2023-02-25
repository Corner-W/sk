package kfk

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type KafkaConfig struct {
	Topic            string `json:"topic"`
	Topic2           string `json:"topic2"`
	GroupId          string `json:"group.id"`
	BootstrapServers string `json:"bootstrap.servers"`
	SecurityProtocol string `json:"security.protocol"`
	SaslMechanism    string `json:"sasl.mechanism"`
	SaslUsername     string `json:"sasl.username"`
	SaslPassword     string `json:"sasl.password"`
}

// config should be a pointer to structure, if not, panic
func LoadJsonConfig() *KafkaConfig {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(workPath, "conf")
	fullPath := filepath.Join(configPath, "kafka.json")
	file, err := os.Open(fullPath)
	if err != nil {
		msg := fmt.Sprintf("Can not load config at %s. Error: %v", fullPath, err)
		panic(msg)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var config = &KafkaConfig{}
	err = decoder.Decode(config)
	if err != nil {
		msg := fmt.Sprintf("Decode json fail for config file at %s. Error: %v", fullPath, err)
		panic(msg)
	}
	json.Marshal(config)
	return config
}
