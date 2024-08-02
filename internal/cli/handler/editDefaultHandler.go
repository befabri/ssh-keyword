package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
)

type EditDefaultHandler struct{}

func (h *EditDefaultHandler) Handle(args []string, connections []config.Connection) {
	if len(args) < 2 {
		fmt.Println("IP address is required for setting a default connection.")
		return
	}
	value := args[1]
	found := false
	for _, connection := range connections {
		if connection.IP == value {
			connection.Default = true
			found = true
		} else {
			connection.Default = false
		}
	}

	if !found {
		fmt.Println("Connection not found.")
		return
	}

	err := config.SaveConnections(connections)
	if err != nil {
		fmt.Println("Failed to set default connection:", err)
		return
	}

	fmt.Println("Default connection set successfully.")
}
