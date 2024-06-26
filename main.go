package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory/v2"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/storage/sqlite3"
	_ "github.com/joho/godotenv/autoload"
	"github.com/wa-broadcast/internal/handler"
	handler_http "github.com/wa-broadcast/internal/handler/http"
	ws_auth "github.com/wa-broadcast/internal/handler/ws"
	"github.com/wa-broadcast/internal/pkg"
)

func main() {
	app := fiber.New()
	app.Static("/", "./static")

	var storage fiber.Storage
	if os.Getenv("DATABASE") == "posgres" {
		storage = postgres.New(postgres.Config{
			ConnectionURI: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("HOST"),
				os.Getenv("PORT"),
				os.Getenv("USERNAME"),
				os.Getenv("PASSWORD"),
				os.Getenv("DATABASE_NAME")),
		})
	} else {
		if err := os.MkdirAll("database", 0775); err != nil {
			panic(err)
		}
		storage = sqlite3.New(sqlite3.Config{
			Database:        "./database/wa-app.db?_foreign_keys=on",
			Table:           "fiber_storage",
			Reset:           false,
			GCInterval:      10 * time.Second,
			MaxOpenConns:    100,
			MaxIdleConns:    100,
			ConnMaxLifetime: 1 * time.Second,
		})
	}

	store := session.New(session.Config{
		Storage:           storage,
		CookieSessionOnly: true,
		CookieHTTPOnly:    true,
	})

	memStore := memory.New()

	pkgWa := pkg.NewWhatsappmeow(os.Getenv("DATABASE"))

	h_middleware := handler.NewHandlerMiddleware(store, pkgWa)

	app.Use("/ws", h_middleware.WS)

	ws_auth := ws_auth.NewWSAuth(memStore, pkgWa)
	{
		app.Get("/ws/auth", websocket.New(ws_auth.Auth))
	}

	app.Use(h_middleware.Authorized)

	h_auth := handler_http.NewHandlerHttpAuth(store, memStore, pkgWa)
	{
		app.Post("/auth", h_auth.PostAuth)
		app.Get("/auth", h_auth.GetAuth)
		app.Get("/logout", h_auth.Logout)
	}

	h_wa := handler_http.NewHandlerHttpWA(store, pkgWa)
	{
		app.Post("/wa/broadcast", h_wa.PostWABroadcast)
		app.Get("/wa/broadcast", h_wa.GetWABroadcast)
	}

	app.Listen(":8000")
}
