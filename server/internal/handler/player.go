package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gamerepo "github.com/krastomer/undercover/server/internal/repository/game"
	gamesvc "github.com/krastomer/undercover/server/internal/service/game"
)

type PlayerHandler interface {
	CreateGame(c *gin.Context)
}

type playerHandler struct {
	gameService gamesvc.GameService
}

func NewPlayerHandler(gameService gamesvc.GameService) PlayerHandler {
	return &playerHandler{gameService: gameService}
}

func (p playerHandler) CreateGame(c *gin.Context) {
	game, err := p.gameService.CreateGame(c.Request.Context(), gamerepo.Player{
		PlayerID: "xxx",
		Name:     "Will",
		IsReveal: false,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"game": game,
	})
}
