package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
)

type RemoveHandler struct{}

func (h *RemoveHandler) Handle(args []string, connections []config.Connection) {
	if len(args) < 2 {
		fmt.Println("IP address is required for removing a connection.")
		return
	}
	value := args[1]
	indexToRemove := utils.FindConnectionIndex(connections, value)
	if indexToRemove == -1 {
		fmt.Println("Connection not found.")
		return
	}

	connections = append(connections[:indexToRemove], connections[indexToRemove+1:]...)
	err := config.SaveConnections(connections)
	if err != nil {
		fmt.Println("Failed to remove connection:", err)
		return
	}

	fmt.Println("Connection removed successfully.")
}
