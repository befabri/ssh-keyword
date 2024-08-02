package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
)

type DefaultHandler struct{}

func (h *DefaultHandler) Handle(args []string, connections []config.Connection) {
	connection, found := utils.FindConnectionDefault(connections)
	if !found {
		fmt.Println("No default server.")
		return
	}
	utils.SshToIP(connection.IP, connection.User, connection.Port)
}
