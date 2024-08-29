package routes

import (
	"github.com/kelbwah/GoChatter/handlers"
	"github.com/kelbwah/GoChatter/server"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func Init(app *echo.Echo) {
	server := server.NewServer()
	app.Static("/public", "ui/public")

	app.GET("/", handlers.HandleShowHome)
	app.GET("/ws", func(c echo.Context) error {
		websocket.Handler(server.HandleWS).ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
