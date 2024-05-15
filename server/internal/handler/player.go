package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krastomer/undercover/server/internal/service/game"
)

type PlayerHandler interface {
	CreateGame(c *gin.Context)
}

type playerHandler struct {
	gameService game.GameService
}

func NewPlayerHandler(gameService game.GameService) PlayerHandler {
	return &playerHandler{gameService: gameService}
}

func (p playerHandler) CreateGame(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello, world",
	})
}
