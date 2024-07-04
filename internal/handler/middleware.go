package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mau.fi/whatsmeow/types"

	"github.com/wa-broadcast/internal/pkg"
)

type HandlerMiddleware interface {
	WS(c *fiber.Ctx) error
	Authorized(c *fiber.Ctx) error
}

type handlerMiddleware struct {
	store *session.Store
	pkgWa pkg.WhatsmeowClient
}

func NewHandlerMiddleware(store *session.Store, pkgWa pkg.WhatsmeowClient) HandlerMiddleware {
	return &handlerMiddleware{store, pkgWa}
}

func (m *handlerMiddleware) WS(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (m *handlerMiddleware) Authorized(c *fiber.Ctx) error {
	sess, err := m.store.Get(c)
	if err != nil {
		panic(err)
	}

	jid, ok := sess.Get("jid").(string)

	if !ok || jid == "" {
		if c.Path() != "/auth" {
			return c.Status(301).Redirect("/auth")
		}
		return c.Next()
	}

	jidParsed, _ := types.ParseJID(jid)

	container := m.pkgWa.GetContainerSql()
	device, err := container.GetDevice(jidParsed)

	if device == nil || err != nil {
		if c.Path() != "/auth" {
			return c.Status(301).Redirect("/auth")
		}
		return c.Next()
	}

	cl := m.pkgWa.GetClient(device)
	c.Locals("client", cl)
	c.Locals("myjid", jid)

	if c.Path() == "/auth" {
		return c.Status(301).Redirect("/wa/broadcast")
	}

	return c.Next()
}
