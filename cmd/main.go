package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yu-nakagawa/go_devcontainer/db"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Wod")

	})
	db.Init()
	r.Run()
	db.Close()
}
