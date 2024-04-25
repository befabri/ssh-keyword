package cli

import (
	"fmt"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
	"strconv"
	"strings"
)

func HandleArgs(args []string, connections []config.Connection) {
	if len(connections) == 0 {
		if len(args) == 0 || !utils.Contains([]string{"-a", "--add", "-h", "--help"}, args[0]) {
			fmt.Println("No server connections available. Please add a server.")
			return
		}
	}

	if len(args) == 0 {
		handleDefault(connections)
		return
	}

	command := args[0]

	if len(args) == 1 && utils.IsIP(command) {
		handleIp(connections, command)
		return
	}

	var value string
	if len(args) > 1 {
		value = args[1]
	}

	switch command {
	case "-a", "--add":
		handleAdd(connections, value)
		return
	case "-d", "--default":
		handleEditDefault(connections, value)
		return
	case "-rm", "--remove":
		handleRemove(connections, value)
		return
	case "-ls", "--list":
		handleList(connections, value)
		return
	case "-e", "--edit":
		handleEdit(connections, value)
		return
	case "-h", "--help":
		handleHelp()
		return
	default:
		handleKeyword(connections, command)
	}
}

func handleAdd(connections []config.Connection, ip string) {
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

func handleRemove(connections []config.Connection, value string) {
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

func handleList(connections []config.Connection, value string) {
	if value == "" {
		fmt.Println("Listing all connections:")
		for _, connection := range connections {
			utils.Display(connection)
		}
	}
	// TODO: Implement filtering by IP or other criteria if needed.
}

func handleEdit(connections []config.Connection, value string) {
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
			connections[indexToEdit].Default = strings.ToLower(newDefault) == "yes" || strings.ToLower(newDefault) == "y"
		case "quit":
			err := config.SaveConnections(connections)
			if err != nil {
				fmt.Println("Failed to save changes:", err)
				return
			}
			fmt.Println("Changes saved successfully.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}

	}
}

func handleEditDefault(connections []config.Connection, value string) {
	found := false
	for i, connection := range connections {
		if connection.IP == value {
			connections[i].Default = true
			found = true
		} else {
			connections[i].Default = false
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

func sshToConnection(connection config.Connection) {
	fmt.Printf("Connecting to server: %s@%s\n", connection.User, connection.IP)
	utils.SshToIP(connection.IP, connection.User, connection.Port)
}

func handleKeyword(connections []config.Connection, command string) {
	connection, found := utils.FindConnectionByKeyword(connections, command)
	if !found {
		fmt.Println("No connection found with the given keyword.")
		fmt.Println("Use -h or --help for usage information.")
		return
	}
	sshToConnection(connection)
}

func handleDefault(connections []config.Connection) {
	connection, found := utils.FindConnectionDefault(connections)
	if !found {
		fmt.Println("No default server.")
		return
	}
	sshToConnection(connection)
}

func handleIp(connections []config.Connection, command string) {
	connection, found := utils.FindConnectionByIP(connections, command)
	if !found {
		fmt.Println("Connection details for IP not found.")
		return
	}
	sshToConnection(connection)
}

func handleHelp() {
	fmt.Println("Usage: ssh-keyword [keyword]")
	fmt.Println("       ssh-keyword [options] [command]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -a,  --add [IP]            Add a connection using the specified IP address.")
	fmt.Println("  -d,  --default [IP]        Set the specified IP as the default connection.")
	fmt.Println("  -rm, --remove [IP|Index]   Remove the connection with the specified IP or at the given index.")
	fmt.Println("  -ls, --list [IP]           List all connections or a specific connection by IP.")
	fmt.Println("  -e,  --edit [IP|Index]     Edit the connection with the specified IP or at the given index.")
	fmt.Println("  -h,  --help                Show this help message and exit.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ssh-keyword server                 Connects directly to the connection labeled 'server'.")
	fmt.Println("  ssh-keyword --add 192.168.1.1      Add a connection for 192.168.1.1.")
	fmt.Println("  ssh-keyword --remove 192.168.1.1   Remove the connection for 192.168.1.1.")
	fmt.Println("  ssh-keyword --list                 List all connections.")
	fmt.Println("  ssh-keyword --edit 2               Edit the connection at index 2.")
	fmt.Println("  ssh-keyword --help                 Show the help message.")
	fmt.Println()
	fmt.Println("Note: For removing or editing a connection, you can specify either the IP address or the index of the connection in the list.")
}
