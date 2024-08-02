package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"ssh-keyword/internal/config"
	"strings"
)

func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func PromptInput(question string, opts ...bool) (string, error) {
	quit := true
	if len(opts) > 0 {
		quit = opts[0]
	}

	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if quit && strings.ToLower(input) == "quit" {
		return "", fmt.Errorf("operation cancelled")
	}

	return input, nil
}

func ConfirmAction(question string) (bool, error) {
	input, err := PromptInput(question + " (Y/n)")
	if err != nil {
		return false, err
	}

	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "", nil
}

func FindConnectionByIP(connections []config.Connection, ip string) (config.Connection, bool) {
	for _, connection := range connections {
		if connection.IP == ip {
			return connection, true
		}
	}
	return config.Connection{}, false
}

func FindConnectionIndex(connections []config.Connection, ip string) int {
	for i, connection := range connections {
		if connection.IP == ip {
			return i
		}
	}
	return -1
}

func FindConnectionByKeyword(connections []config.Connection, keyword string) (config.Connection, bool) {
	for _, connection := range connections {
		for _, k := range connection.Keywords {
			if strings.EqualFold(k, keyword) {
				return connection, true
			}
		}
	}
	return config.Connection{}, false
}

func FindConnectionDefault(connections []config.Connection) (config.Connection, bool) {
	for _, connection := range connections {
		if connection.Default {
			return connection, true
		}
	}
	return config.Connection{}, false
}

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func SshToIP(ip string, user string, port string) {
	fmt.Printf("Connecting to server: %s@%s\n", user, ip)
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", user, ip), "-p", port)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to connect:", err)
	}
}

func Display(connection config.Connection) {
	keywords := strings.Join(connection.Keywords, ", ")
	fmt.Printf("ip: %s, user: %s, port: %s, keywords: [%s], default: %t\n", connection.IP, connection.User, connection.Port, keywords, connection.Default)
}
