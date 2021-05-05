package main

import (
	"github.com/yu-nakagawa/go_devcontainer/db"
	"github.com/yu-nakagawa/go_devcontainer/server"
)

func main() {
	db.Init()
	server.Init()
}
