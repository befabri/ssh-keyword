package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
)

type IpHandler struct{}

func (h *IpHandler) Handle(args []string, connections []config.Connection) {
	if len(args) < 1 {
		fmt.Println("Ip is required.")
		return
	}
	command := args[0]
	connection, found := utils.FindConnectionByIP(connections, command)
	if !found {
		fmt.Println("Connection details for IP not found.")
		return
	}
	utils.SshToIP(connection.IP, connection.User, connection.Port)
}
