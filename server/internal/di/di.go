package di

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/undercover/server/internal/handler"
	"github.com/krastomer/undercover/server/internal/router"
	"github.com/krastomer/undercover/server/internal/service/broadcast"
	"github.com/krastomer/undercover/server/internal/service/game"
	"github.com/krastomer/undercover/server/pkg/connection"
)

func NewContainer() (Container, error) {
	redisClient, err := connection.NewRedisClient()
	if err != nil {
		return Container{}, err
	}

	_ = redisClient

	// TODO: repository for bridge

	broadcastSvc := broadcast.NewBroadcastService()
	gameSvc := game.NewGameService(broadcastSvc)

	playerHandler := handler.NewPlayerHandler(gameSvc)
	webSocketHandler := handler.NewWebSocketHandler(broadcastSvc)

	handlers := handler.Handlers{
		PlayerHandler:    playerHandler,
		WebSocketHandler: webSocketHandler,
	}

	r := gin.Default()

	router.NewRouter(r, handlers)

	ctn := Container{
		Server: r,
	}

	return ctn, nil
}
