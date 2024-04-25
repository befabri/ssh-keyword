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

func PromptInput(question string) (string, error) {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if strings.ToLower(input) == "quit" {
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
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", user, ip), "-p", port)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to connect:", err)
	}
}
