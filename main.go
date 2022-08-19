package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type IdentifyPacket struct {
	Operation uint8              `json:"op"`
	Data      IdentifyPacketData `json:"d"`
}

type IdentifyPacketData struct {
	Token   string
	Intents uint64
	Os      string
	Browser string
	Device  string
}

type BaseEvent struct {
	Operation int             `json:"op"`
	Sequence  int64           `json:"s"`
	Type      string          `json:"t"`
	Data      json.RawMessage `json:"d"`
}

type HelloEvent struct {
	HeartBeatInterval int `json:"heartbeat_interval"`
}

func main() {
	dialer := websocket.DefaultDialer
	header := http.Header{}
	header.Add("accept-encoding", "json")
	ws, _, err := dialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", header)

	if err != nil {
		log.Printf("error connecting to gateway %s", err)
	}

	d := IdentifyPacket{
		Operation: 2,
		Data: IdentifyPacketData{
			Token:   os.Getenv("DISCORD_TOKEN"),
			Intents: (1 << 22) - 1,
			Os:      "window",
			Browser: "discord",
			Device:  "discord",
		},
	}

	ws.WriteJSON(d)

	_, m, err := ws.ReadMessage()

	if err != nil {
		fmt.Println(err)
	}

	var event *BaseEvent

	decoder := json.NewDecoder(bytes.NewBuffer(m))

	if err := decoder.Decode(&event); err != nil {
		fmt.Println(err)
	}

	var interval *HelloEvent
	json.Unmarshal(event.Data, &interval)

}
