package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Connection struct {
	Default  bool     `json:"default"`
	IP       string   `json:"ip"`
	Keywords []string `json:"keywords"`
	User     string   `json:"user"`
	Port     string   `json:"port"`
}

var jsonFileName = "connections.json"

func LoadConnections() ([]Connection, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var connections []Connection
	err = json.Unmarshal(file, &connections)
	if err != nil {
		return nil, err
	}
	return connections, nil
}

func SaveConnections(connections []Connection) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(connections, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0644)
	return err
}

func GetConfigPath() (string, error) {
	exePath, err := os.Executable() // Get the executable path
	if err != nil {
		return "", fmt.Errorf("error getting executable path: %v", err)
	}
	dir := filepath.Dir(exePath)
	jsonFilePath := filepath.Join(dir, jsonFileName)
	return jsonFilePath, nil
}
