package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Token        string
	Websocket    *websocket.Conn
	LastSequence int
	Heartbeat    int
}

func (client *Client) heartbeatTimer() {
	date := time.Now().UnixMilli()
	for {
		now := time.Now().UnixMilli()
		if int(now)-int(date) < client.Heartbeat {
			continue
		} else {
			client.sendHeartbeat()
			return
		}
	}
}

func (client *Client) sendHeartbeat() {
	packet := HeartbeatPacket{
		Operation: OperationCode(Heartbeat),
		Sequence:  client.LastSequence,
	}

	client.Websocket.WriteJSON(packet)
}

func (client *Client) handleEvent() {
	var baseEvent *BaseEvent

	if err := json.NewDecoder(bytes.NewBuffer(m)).Decode(&baseEvent); err != nil {
		log.Printf("error decoding base event \n%s", err)
	}

	// fmt.Printf("%+v", baseEvent)

	switch baseEvent.Operation {
	case OperationCode(Hello):
		var event *BasicEvent
		if err := json.NewDecoder(bytes.NewBuffer(m)).Decode(&event); err != nil {
			log.Printf("error decoding basic event \n%s", err)
		}

		var data *HelloEvent
		if err := json.Unmarshal([]byte(event.Data), &data); err != nil {
			log.Printf("error Unmarshaling data \n%s", err)
		}
		client.Heartbeat = 5000
		client.heartbeatTimer()
	case OperationCode(HeartbeatACK):
		fmt.Println(baseEvent)
	}
}

func (client *Client) listen() {
	for {
		if client.Websocket == nil {
			log.Printf("error handling Event\nreason: Ws is nil")
		}

		msgType, msg, err := client.Websocket.ReadMessage()

		if err != nil {
			log.Printf("error reading websocket message \n%s", err)
		}

		
	}
}

func newClient() Client {
	return Client{
		Token: os.Getenv("DISCORD_TOKEN"),
	}
}
