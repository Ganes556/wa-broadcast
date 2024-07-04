package handler_http

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"

	"github.com/wa-broadcast/internal/handler"
	"github.com/wa-broadcast/internal/pkg"
	view_broadcast "github.com/wa-broadcast/view/page/broadcast"
	view_partial_alert "github.com/wa-broadcast/view/partial/alert"
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

	// check number
	jids := make([]types.JID, len(req.Numbers))
	for i, v := range req.Numbers {
		parsedJid, err := types.ParseJID(v + "@s.whatsapp.net")
		if err != nil {
			log.Println(err)
			return c.SendString("err")
		}
		jids[i] = parsedJid
	}

	// sending process
	wg := new(sync.WaitGroup)
	guard := make(chan struct{}, 10)
	wg.Add(len(req.Numbers))
	for _, jid := range jids {
		guard <- struct{}{}
		go func(v types.JID) {
			defer func() {
				wg.Done()
				<-guard
			}()
			cl.SendMessage(context.Background(), jid, &waE2E.Message{
				Conversation: proto.String(req.Message),
			})
		}(jid)
	}
	wg.Wait()
	defer cl.Disconnect()
	return handler.Render(c, view_partial_alert.Alert(200, "success"))
}
