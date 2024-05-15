package broadcast

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

type BroadcastService interface {
	NewClient(ctx context.Context, playerID string, w http.ResponseWriter, r *http.Request) error
	SendMessage(ctx context.Context, msg string, playerIDs []string) error
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
		CheckOrigin: func(r *http.Request) bool {
			// TODO: check secure
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	b.client[playerID] = conn
	return nil
}

func (b *broadcastService) SendMessage(ctx context.Context, msg string, playerIDs []string) error {
	for _, playerID := range playerIDs {
		client, ok := b.client[playerID]
		if !ok {
			return errors.New("player lost connection")
		}

		client.WriteMessage(websocket.TextMessage, []byte(msg))
	}

	return nil
}
