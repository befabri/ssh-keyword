package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
)

type KeywordHandler struct{}

func (h *KeywordHandler) Handle(args []string, connections []config.Connection) {
	if len(args) < 1 {
		fmt.Println("Keyword is required.")
		return
	}
	command := args[0]
	connection, found := utils.FindConnectionByKeyword(connections, command)
	if !found {
		fmt.Println("No connection found with the given keyword.")
		fmt.Println("Use -h or --help for usage information.")
		return
	}
	utils.SshToIP(connection.IP, connection.User, connection.Port)
}
