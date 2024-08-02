package handler

import "ssh-keyword/internal/config"

type Handler interface {
	Handle(args []string, connections []config.Connection)
}
