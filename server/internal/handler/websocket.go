package handler

import "github.com/gin-gonic/gin"

type WebSocketHandler interface {
	Upgrade(c *gin.Context)
}

type webSocketHandler struct {
}
