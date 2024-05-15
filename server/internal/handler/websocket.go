package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krastomer/undercover/server/internal/service/broadcast"
)

type WebSocketHandler interface {
	Upgrade(c *gin.Context)
}

type webSocketHandler struct {
	broadcastSvc broadcast.BroadcastService
}

func NewWebSocketHandler(broadcastSvc broadcast.BroadcastService) WebSocketHandler {
	return &webSocketHandler{broadcastSvc: broadcastSvc}
}

func (h webSocketHandler) Upgrade(c *gin.Context) {
	// TODO: check user id
	token := c.Query("token")
	fmt.Println(token)
	err := h.broadcastSvc.NewClient(c.Request.Context(), token, c.Writer, c.Request)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}
