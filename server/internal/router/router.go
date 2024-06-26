package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/krastomer/undercover/server/internal/handler"
)

func NewRouter(router *gin.Engine, handlers handler.Handlers) *gin.Engine {
	router.Use(cors.New(cors.Config{
		AllowWebSockets:        true,
		AllowAllOrigins:        true,
		AllowBrowserExtensions: true,
	}))

	router.GET("/ws", handlers.WebSocketHandler.Upgrade)

	// create game
	router.POST("/api/v1/game", handlers.PlayerHandler.CreateGame)
	// join game
	router.POST("/api/v1/game/:game_id", nil)
	// start game
	router.POST("/api/v1/game/:game_id/start", nil)
	// vote player
	router.POST("/api/v1/game/:game_id/vote", nil)
	// for white guy assume word
	router.POST("/api/v1/game/:game_id/assume", nil)

	return router
}
