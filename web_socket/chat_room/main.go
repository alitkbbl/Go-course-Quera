package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	rooms   = make(map[string]map[*websocket.Conn]bool)
	roomsMu sync.Mutex
)

func wsChatRoom(c echo.Context) error {
	roomId := c.Param("roomId")
	username := c.Param("username")

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	roomsMu.Lock()
	if rooms[roomId] == nil {
		rooms[roomId] = make(map[*websocket.Conn]bool)
	}
	rooms[roomId][conn] = true
	roomsMu.Unlock()

	defer func() {
		roomsMu.Lock()
		delete(rooms[roomId], conn)
		roomsMu.Unlock()
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		broadcastMessage := username + ": " + string(msg)

		roomsMu.Lock()
		for c2 := range rooms[roomId] {
			if c2 == conn {
				continue
			}
			err := c2.WriteMessage(websocket.TextMessage, []byte(broadcastMessage))
			if err != nil {
				c2.Close()
				delete(rooms[roomId], c2)
			}
		}
		roomsMu.Unlock()
	}

	return nil
}

func main() {
	e := echo.New()
	e.GET("/ws/chat/:roomId/user/:username", wsChatRoom)
	e.Logger.Fatal(e.Start(":8080"))
}
