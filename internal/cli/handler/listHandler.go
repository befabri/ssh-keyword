package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
)

type ListHandler struct{}

func (h *ListHandler) Handle(args []string, connections []config.Connection) {
	fmt.Println("Listing all connections:")
	for _, connection := range connections {
		utils.Display(connection)
	}
}
