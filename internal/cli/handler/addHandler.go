package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
	"strconv"
	"strings"
)

type AddHandler struct{}

func (h *AddHandler) Handle(args []string, connections []config.Connection) {
	if len(args) < 2 {
		fmt.Println("IP address is required for adding a connection.")
		return
	}
	ip := args[1]
	fmt.Println("Entry for new connection")
	if !utils.IsIP(ip) {
		fmt.Println("Invalid IP address.")
		return
	}
	fmt.Println("Enter 'Quit' for exit")
	fmt.Println()

	for _, connection := range connections {
		if connection.IP == ip {
			fmt.Println("Connection already exists.")
			return
		}
	}

	user, err := utils.PromptInput("Enter a user: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	port, err := utils.PromptInput("Port: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := strconv.Atoi(port); err != nil {
		fmt.Println("Invalid port.")
		return
	}

	keywordsInput, err := utils.PromptInput("Enter a list of keywords separated by ',': ")
	if err != nil {
		fmt.Println(err)
		return
	}
	keywords := strings.Split(keywordsInput, ",")
	for i, keyword := range keywords {
		keywords[i] = strings.TrimSpace(keyword)
	}

	defaultInput, err := utils.PromptInput("Default server ([Y]es | [N]o): ")
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.ToLower(defaultInput) == "y" || strings.ToLower(defaultInput) == "yes" {
		for i := range connections {
			connections[i].Default = false
		}
	}
	isDefault := strings.ToLower(defaultInput) == "y" || strings.ToLower(defaultInput) == "yes"

	newConn := config.Connection{
		IP:       ip,
		User:     user,
		Port:     port,
		Keywords: keywords,
		Default:  isDefault,
	}

	connections = append(connections, newConn)
	err = config.SaveConnections(connections)
	if err != nil {
		fmt.Println("Failed to save new connection:", err)
		return
	}

	fmt.Println("Connection added successfully.")
}
