package handler

import "github.com/krastomer/undercover/server/internal/service/game"

type PlayerHandler interface {
}

type playerHandler struct {
	gameService game.GameService
}

func NewPlayerHandler(gameService game.GameService) PlayerHandler {
	return playerHandler{gameService: gameService}
}
