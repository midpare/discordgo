package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	client := newClient()

	dialer := websocket.DefaultDialer
	header := http.Header{}
	header.Add("accept-encoding", "json")
	websocket, _, err := dialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", header)

	if err != nil {
		log.Printf("error connecting to gateway \n%s", err)
	}

	client.Websocket = websocket
	packet := IdentifyPacket{
		Operation: OperationCode(Identify),
		Data: IdentifyPacketData{
			Token:   client.Token,
			Intents: (1 << 22) - 1,
			Os:      "window",
			Browser: "discord",
			Device:  "discord",
		},
	}
	client.listen()
	client.handleEvent()
	client.Websocket.WriteJSON(packet)
}
