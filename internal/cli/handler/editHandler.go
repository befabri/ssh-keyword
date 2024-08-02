package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
	"strconv"
	"strings"
)

type EditHandler struct{}

func (h *EditHandler) Handle(args []string, connections []config.Connection) {
	if len(args) < 2 {
		fmt.Println("IP address is required for editing a connection.")
		return
	}
	value := args[1]
	indexToEdit := utils.FindConnectionIndex(connections, value)
	if indexToEdit == -1 {
		fmt.Println("Connection not found.")
		return
	}

	fmt.Printf("Editing connection for IP: %s\n", connections[indexToEdit].IP)
	for {
		fmt.Println()
		utils.Display(connections[indexToEdit])
		fieldToEdit, err := utils.PromptInput("Enter the field to edit (ip, user, port, keywords, default): ", false)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if strings.ToLower(fieldToEdit) == "exit" || strings.ToLower(fieldToEdit) == "q" || strings.ToLower(fieldToEdit) == "quit" {
			err := config.SaveConnections(connections)
			if err != nil {
				fmt.Println("Failed to save changes:", err)
			} else {
				fmt.Println("Changes saved successfully.")
			}
			return
		}

		switch strings.ToLower(fieldToEdit) {
		case "ip":
			newIP, _ := utils.PromptInput("Enter new IP: ")
			if utils.IsIP(newIP) {
				connections[indexToEdit].IP = newIP
			} else {
				fmt.Println("Invalid IP address.")
				continue
			}
		case "user":
			newUser, _ := utils.PromptInput("Enter new user: ")
			connections[indexToEdit].User = newUser
		case "port":
			newPort, _ := utils.PromptInput("Enter new port: ")
			if _, err := strconv.Atoi(newPort); err == nil {
				connections[indexToEdit].Port = newPort
			} else {
				fmt.Println("Invalid port.")
				continue
			}
		case "keywords":
			newKeywords, _ := utils.PromptInput("Enter new keywords (comma-separated): ")
			keywordList := strings.Split(newKeywords, ",")
			for i, keyword := range keywordList {
				keywordList[i] = strings.TrimSpace(keyword)
			}
			connections[indexToEdit].Keywords = keywordList
		case "default":
			newDefault, _ := utils.PromptInput("Set as default? (yes/no): ")
			if strings.ToLower(newDefault) == "y" || strings.ToLower(newDefault) == "yes" {
				for i := range connections {
					connections[i].Default = false
				}
			}
			connections[indexToEdit].Default = strings.ToLower(newDefault) == "yes" || strings.ToLower(newDefault) == "y"
		default:
			fmt.Println("Invalid choice. Try again.")
		}

	}
}
