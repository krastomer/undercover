package game

import (
	"context"

	"github.com/krastomer/undercover/server/internal/service/broadcast"
)

type GameService interface {
	CreateGame()
	JoinGame()
	StartGame(ctx context.Context, gameID string) error

	VotePlayer(ctx context.Context, gameID, targetPlayerID string) error
	GuessWord(ctx context.Context, gameID string) error
}

type gameService struct {
	broadcastService broadcast.BroadcastService
}

func NewGameService(broadcastSvc broadcast.BroadcastService) GameService {
	return &gameService{
		broadcastService: broadcastSvc,
	}
}

// CreateGame implements GameService.
func (g *gameService) CreateGame() {
	panic("unimplemented")
}

// JoinGame implements GameService.
func (g *gameService) JoinGame() {
	panic("unimplemented")
}

// StartGame implements GameService.
func (g gameService) StartGame(ctx context.Context, gameID string) error {
	// g.broadcastService.

	return nil
}

// GuessWord implements GameService.
func (g *gameService) GuessWord(ctx context.Context, gameID string) error {
	panic("unimplemented")
}

// VotePlayer implements GameService.
func (g *gameService) VotePlayer(ctx context.Context, gameID string, targetPlayerID string) error {
	panic("unimplemented")
}
