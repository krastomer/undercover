package broadcast

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

type BroadcastService interface {
	NewClient(ctx context.Context, playerID string, w http.ResponseWriter, r *http.Request) error
	SendMessage(ctx context.Context, playerID string, msg string) error
	SendMessageAll(ctx context.Context, msg string) error
}

type broadcastService struct {
	client map[string]*websocket.Conn
}

func NewBroadcastService() BroadcastService {
	return &broadcastService{
		client: map[string]*websocket.Conn{},
	}
}

func (b *broadcastService) NewClient(ctx context.Context, playerID string, w http.ResponseWriter, r *http.Request) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	b.client[playerID] = conn
	return nil
}

func (b *broadcastService) SendMessage(ctx context.Context, playerID string, msg string) error {
	client, ok := b.client[playerID]
	if !ok {
		return errors.New("player lost connection")
	}

	return client.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (b *broadcastService) SendMessageAll(ctx context.Context, msg string) error {
	for playerID := range b.client {
		b.SendMessage(ctx, playerID, msg)
	}

	return nil
}
