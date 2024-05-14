package game

import (
	"context"

	"github.com/google/uuid"
	gamerepo "github.com/krastomer/undercover/server/internal/repository/game"
	"github.com/krastomer/undercover/server/internal/service/broadcast"
)

type GameService interface {
	CreateGame(ctx context.Context, hostPlayer gamerepo.Player) (gamerepo.Game, error)
	JoinGame(ctx context.Context, gameID string, player gamerepo.Player) (gamerepo.Game, error)
	StartGame(ctx context.Context, gameID string) error

	VotePlayer(ctx context.Context, gameID, targetPlayerID string) error
	GuessWord(ctx context.Context, gameID string) error
}

type gameService struct {
	broadcastService broadcast.BroadcastService
	repository       gamerepo.Repository
	// redisClient      *redis.Client
}

func NewGameService(broadcastSvc broadcast.BroadcastService, repository gamerepo.Repository) GameService {
	return &gameService{
		broadcastService: broadcastSvc,
		repository:       repository,
	}
}

func (g *gameService) CreateGame(ctx context.Context, hostPlayer gamerepo.Player) (gamerepo.Game, error) {
	newGame := gamerepo.Game{
		GameID:  uuid.NewString(),
		Players: []gamerepo.Player{hostPlayer},
		Host:    hostPlayer,
	}

	err := g.repository.SetGame(ctx, newGame)
	if err != nil {
		return gamerepo.Game{}, err
	}

	return newGame, nil
}

func (g *gameService) JoinGame(ctx context.Context, gameID string, player gamerepo.Player) (gamerepo.Game, error) {
	game, err := g.repository.GetGame(ctx, gameID)
	if err != nil {
		return gamerepo.Game{}, err
	}

	game.Players = append(game.Players, player)
	if err := g.repository.SetGame(ctx, game); err != nil {
		return gamerepo.Game{}, err
	}

	return game, nil
}

// StartGame implements GameService.
func (g gameService) StartGame(ctx context.Context, gameID string) error {
	// large logic
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
