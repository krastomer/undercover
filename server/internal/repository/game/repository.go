package game

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	prefixGame     = "game:%s"
	prefixGameVote = "game:%s:vote"
)

type Repository interface {
	GetGame(ctx context.Context, gameID string) (Game, error)
	SetGame(ctx context.Context, game Game) error
}

type repository struct {
	redisClient *redis.Client
}

func NewRepository(redisClient *redis.Client) Repository {
	return &repository{redisClient: redisClient}
}

func (r repository) GetGame(ctx context.Context, gameID string) (Game, error) {
	result, err := r.redisClient.Get(ctx, prefixGame).Result()
	if err != nil {
		return Game{}, err
	}

	var game Game
	if err := json.Unmarshal([]byte(result), &game); err != nil {
		return Game{}, err
	}

	return game, nil
}

func (r repository) SetGame(ctx context.Context, game Game) error {
	if err := r.redisClient.Set(ctx, fmt.Sprintf("%s%s", prefixGame, game.GameID), game, 0).Err(); err != nil {
		return err
	}

	return nil
}
