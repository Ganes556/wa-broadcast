package handler_http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory/v2"
	"github.com/wa-broadcast/internal/handler"
	"github.com/wa-broadcast/internal/pkg"
	view_auth "github.com/wa-broadcast/view/page/auth"
	view_partial_alert "github.com/wa-broadcast/view/partial/alert"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

type HandlerHttpAuth interface {
	GetAuth(c *fiber.Ctx) error
	PostAuth(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type handlerHttpAuth struct {
	store    *session.Store
	memStore *memory.Storage
	pkgWa    pkg.WhatsmeowClient
}

func NewHandlerHttpAuth(store *session.Store, memStore *memory.Storage, pkgWa pkg.WhatsmeowClient) HandlerHttpAuth {
	return &handlerHttpAuth{store, memStore, pkgWa}
}

type auth struct {
	Token string `json:"token"`
}

func (h *handlerHttpAuth) GetAuth(c *fiber.Ctx) error {
	return handler.Render(c, view_auth.Login())
}

func (h *handlerHttpAuth) Logout(c *fiber.Ctx) error {
	sess, _ := h.store.Get(c)
	jid, ok := sess.Get("jid").(string)
	if !ok {
		c.Set("HX-Redirect", "/auth")
		return c.SendStatus(fiber.StatusSeeOther)
	}

	cl := c.Locals("client").(*whatsmeow.Client)
	if !cl.IsConnected() {
		cl.Connect()
		cl.WaitForConnection(time.Duration(5 * time.Second))
	}
	cl.Logout()
	if cl.IsConnected() {
		cl.Disconnect()
	}
	sess.Delete(jid)
	c.Set("HX-Redirect", "/auth")
	return c.SendStatus(fiber.StatusSeeOther)
}

func (h *handlerHttpAuth) PostAuth(c *fiber.Ctx) error {
	req := new(auth)
	c.BodyParser(req)
	if req.Token == "" {
		return handler.Render(c.Status(400), view_partial_alert.Alert(400, fiber.ErrBadRequest.Message))
	}

	jidStr, err := h.memStore.Get(req.Token)
	if err != nil {
		return handler.Render(c.Status(500), view_partial_alert.Alert(500, fiber.ErrInternalServerError.Message))
	}

	if err := h.memStore.Delete(req.Token); err != nil {
		return handler.Render(c.Status(500), view_partial_alert.Alert(500, fiber.ErrInternalServerError.Message))
	}

	jid, err := types.ParseJID(string(jidStr))

	if err != nil {
		return handler.Render(c.Status(500), view_partial_alert.Alert(500, fiber.ErrInternalServerError.Message))
	}

	container := h.pkgWa.GetContainerSql()
	device, err := container.GetDevice(jid)
	if err != nil {
		panic(err)
	}

	if device == nil || device.ID == nil {
		return handler.Render(c, view_partial_alert.Alert(401, fiber.ErrUnauthorized.Message))
	}

	sess, _ := h.store.Get(c)
	sess.Set("jid", string(jidStr))
	sess.SetExpiry(24 * time.Hour * 14)
	sess.Save()
	c.Set("Location", "/wa/broadcast")
	return c.SendStatus(200)

}
