package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	c.Logger().Info("hello")
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Read
			var msg Msg
			err := websocket.JSON.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			fmt.Printf("%s\n", msg)

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()
	e.Logger.Fatal(e.Start(":1323"))
}
