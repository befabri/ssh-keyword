package handler

import (
	"fmt"
	"ssh-keyword/internal/config"
)

type HelpHandler struct{}

func (h *HelpHandler) Handle(args []string, connections []config.Connection) {
	fmt.Println("Usage: ssh-keyword [keyword]")
	fmt.Println("       ssh-keyword [options] [command]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -a,  --add [IP]                     Add a connection using the specified IP address.")
	fmt.Println("  -d,  --default [IP]                 Set the specified IP as the default connection.")
	fmt.Println("  -rm, --remove [IP]                  Remove the connection with the specified IP.")
	fmt.Println("  -ls, --list [IP]                    List all connections or a specific connection by IP.")
	fmt.Println("  -e,  --edit [IP]                    Edit the connection with the specified IP.")
	fmt.Println("  -h,  --help                         Show this help message and exit.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ssh-keyword                         Connects to the default connection.")
	fmt.Println("  ssh-keyword server                  Connects to the connection labeled 'server'.")
	fmt.Println("  ssh-keyword 192.168.1.1             Connects to 192.168.1.1.")
	fmt.Println("  ssh-keyword --default 192.168.1.1   Sets 192.168.1.1 as the default connection.")
	fmt.Println("  ssh-keyword --add 192.168.1.1       Add a connection for 192.168.1.1.")
	fmt.Println("  ssh-keyword --remove 192.168.1.1    Remove the connection for 192.168.1.1.")
	fmt.Println("  ssh-keyword --list                  List all connections.")
	fmt.Println("  ssh-keyword --edit 192.168.1.1      Edit the connection for 192.168.1.1.")
	fmt.Println("  ssh-keyword --help                  Show the help message.")
	fmt.Println()
}
