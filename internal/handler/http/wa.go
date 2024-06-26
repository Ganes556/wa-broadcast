package handler_http

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/wa-broadcast/internal/handler"
	"github.com/wa-broadcast/internal/pkg"
	view_broadcast "github.com/wa-broadcast/view/page/broadcast"
	view_partial_alert "github.com/wa-broadcast/view/partial/alert"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

type HandlerHttpWA interface {
	GetWABroadcast(c *fiber.Ctx) error
	PostWABroadcast(c *fiber.Ctx) error
}

type handlerHttpWA struct {
	store *session.Store
	pkgWa pkg.WhatsmeowClient
}

func NewHandlerHttpWA(store *session.Store, pkgWa pkg.WhatsmeowClient) HandlerHttpWA {
	return &handlerHttpWA{store, pkgWa}
}

type body struct {
	Numbers []string `form:"numbers"`
	Message string   `form:"message"`
}

func (h *handlerHttpWA) GetWABroadcast(c *fiber.Ctx) error {
	return handler.Render(c, view_broadcast.Broadcast())
}

func (h *handlerHttpWA) PostWABroadcast(c *fiber.Ctx) error {
	var req = new(body)
	c.BodyParser(req)
	if req.Message == "" || req.Numbers == nil {
		return handler.Render(c, view_partial_alert.Alert(400, fiber.ErrBadRequest.Message))
	}

	cl := c.Locals("client").(*whatsmeow.Client)
	if !cl.IsConnected() {
		cl.Connect()
		cl.WaitForConnection(time.Duration(5 * time.Second))
	}

	for _, v := range req.Numbers {
		v += "@s.whatsapp.net"
		jid, err := types.ParseJID(v)
		if err != nil {
			return c.SendString("err")
		}
		cl.SendMessage(context.Background(), jid, &waE2E.Message{
			Conversation: proto.String(req.Message),
		})
	}
	defer cl.Disconnect()
	return handler.Render(c, view_partial_alert.Alert(200, "success"))
}
