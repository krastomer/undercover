package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/undercover/server/internal/service/broadcast"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

func main() {
	router := gin.Default()

	broadcastSvc := broadcast.NewBroadcastService()

	router.GET("/ws", func(c *gin.Context) {
		userID := c.GetHeader("x-user-id")
		err := broadcastSvc.NewClient(c.Request.Context(), userID, c.Writer, c.Request)
		if err != nil {
			c.JSON(500, err)
			return
		}
	})

	router.GET("/trigger", func(c *gin.Context) {
		broadcastSvc.SendMessageAll(c.Request.Context(), "all message")
	})

	router.Run(":8080")
}
