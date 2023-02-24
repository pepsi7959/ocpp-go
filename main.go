package main

import (
	"github.com/lorenzodonini/ocpp-go/ws"
	"log"
)

func main() {
	server := ws.NewServer()

	server.SetNewClientHandler(func(ws ws.Channel) {
		log.Println("new client connected !!!")
	})

	server.SetDisconnectedClientHandler(func(ws ws.Channel) {
		log.Println("client disconnected !!!")
	})

	server.SetMessageHandler(func(ws ws.Channel, data []byte) error {
		log.Printf("recev data: %s\n", string(data))
		return nil
	})

	server.Start(8080, "ws/:id")
}
