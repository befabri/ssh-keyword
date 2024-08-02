package cli

import (
	"fmt"
	"ssh-keyword/internal/cli/handler"
	"ssh-keyword/internal/config"
	"ssh-keyword/internal/utils"
)

type CLI struct {
	handlers map[string]handler.Handler
}

func NewCLI() *CLI {
	cli := &CLI{handlers: make(map[string]handler.Handler)}
	cli.handlers["-a"] = &handler.AddHandler{}
	cli.handlers["--add"] = &handler.AddHandler{}
	cli.handlers["-d"] = &handler.EditDefaultHandler{}
	cli.handlers["--default"] = &handler.EditDefaultHandler{}
	cli.handlers["-rm"] = &handler.RemoveHandler{}
	cli.handlers["--remove"] = &handler.RemoveHandler{}
	cli.handlers["-ls"] = &handler.ListHandler{}
	cli.handlers["--list"] = &handler.ListHandler{}
	cli.handlers["-e"] = &handler.EditHandler{}
	cli.handlers["--edit"] = &handler.EditHandler{}
	cli.handlers["-h"] = &handler.HelpHandler{}
	cli.handlers["--help"] = &handler.HelpHandler{}

	cli.handlers["ip"] = &handler.IpHandler{}
	cli.handlers["keyword"] = &handler.KeywordHandler{}
	cli.handlers["default"] = &handler.KeywordHandler{}
	return cli
}

func (cli *CLI) HandleArgs(args []string, connections []config.Connection) {
	if len(connections) == 0 {
		if len(args) == 0 || !utils.Contains([]string{"-a", "--add", "-h", "--help"}, args[0]) {
			fmt.Println("No server connections available. Please add a server.")
			return
		}
	}

	if len(args) == 0 {
		cli.handlers["default"].Handle(args, connections)
		return
	}

	command := args[0]

	if len(args) == 1 && utils.IsIP(command) {
		cli.handlers["ip"].Handle(args, connections)
		return
	}

	if handler, exists := cli.handlers[command]; exists {
		handler.Handle(args, connections)
	} else {
		cli.handlers["keyword"].Handle(args, connections)
	}
}
