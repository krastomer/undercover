package di

import "github.com/gin-gonic/gin"

type Container struct {
	Server *gin.Engine
}

func (c *Container) Start() {
	c.Server.Run(":8080")
}
