package ws_auth

import (
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/memory/v2"
	"github.com/google/uuid"
	"github.com/wa-broadcast/internal/pkg"
	"go.mau.fi/whatsmeow/types"
)

type WSAuth interface {
	Auth(c *websocket.Conn)
}

type wSAuth struct {
	memStore *memory.Storage
	pkgWa    pkg.WhatsmeowClient
}

func NewWSAuth(memStore *memory.Storage, pkgWa pkg.WhatsmeowClient) WSAuth {
	return &wSAuth{memStore, pkgWa}
}

func (w *wSAuth) Auth(c *websocket.Conn) {
	container := w.pkgWa.GetContainerSql()
	newDevice := container.NewDevice()
	cl := w.pkgWa.GetClient(newDevice)

	w.pkgWa.Auth(cl, func(context string, data any) {
		if context == "success" {
			token := uuid.New().String()
			w.memStore.Set(token, []byte(data.(*types.JID).String()), 10*time.Second)
			data = token
		}
		c.WriteJSON(fiber.Map{
			"context": context,
			"data":    data,
		})
	})
}
