package main

import (
	"fmt"
	"os"
	"ssh-keyword/internal/cli"
	"ssh-keyword/internal/config"
)

func main() {
	connections, err := config.LoadConnections()
	if err != nil {
		if os.IsNotExist(err) {
			connections = []config.Connection{}
		} else {
			fmt.Println("Error loading connections:", err)
			return
		}
	}
	cli.HandleArgs(os.Args[1:], connections)
}
