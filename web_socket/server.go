package main

import (
	_ "fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = "localhost:8080"

var upgrader = websocket.Upgrader{}

// TODO: add a websocket handler by `upgrader`

func main() {
	// TODO: Add handler to receive and send messages

	log.Fatal(http.ListenAndServe(addr, nil))
}
