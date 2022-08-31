package client

import (
	"bytes"
	"discordbot/src/events"
	. "discordbot/src/structs"
	"discordbot/src/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	sync.Mutex
	Token             string
	Websocket         *websocket.Conn
	LastSequence      int
	HeartbeatInterval int
	HeartbeatDate     int64
	ResumeDate        int64
	SessionId         string
	ResumeGatewayUrl  string
	Guilds            map[utils.Snowflake]*Guild
	Channels          map[utils.Snowflake]*Channel
}

func NewClient() (client Client, err error) {
	dialer := websocket.DefaultDialer
	header := http.Header{}
	header.Add("accept-encoding", "json")
	websocket, _, err := dialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", header)

	if err != nil {
		return Client{}, err
	}

	return Client{
		Token:         os.Getenv("DISCORD_TOKEN"),
		Websocket:     websocket,
		HeartbeatDate: time.Now().UnixMilli(),
		ResumeDate:    time.Now().UnixMilli(),
		Guilds:        make(map[utils.Snowflake]*Guild),
		Channels:      make(map[utils.Snowflake]*Channel),
	}, nil
}

func (client *Client) Login() {
	msgType, msg, err := client.Websocket.ReadMessage()

	if err != nil {
		log.Fatalf("error reading hello message \n%s", err)
	}

	client.HandleMessage(msgType, msg)
	go client.Heartbeat()
	go client.Resume()

	packet := IdentifyPacket{
		Operation: OperationCode_Identify,
		Data: IdentifyPacketData{
			Token:   client.Token,
			Intents: uint64(IntentsAll),
			Properties: IdentifyPacketDataProperties{
				Os:      "window",
				Browser: "discord",
				Device:  "discord",
			},
		},
	}

	client.Websocket.WriteJSON(packet)

	client.listen()
}

func (client *Client) Heartbeat() {
	for {
		date := client.HeartbeatDate
		now := time.Now().UnixMilli()
		if int(now)-int(date) < client.HeartbeatInterval/3 {
			continue
		} else {
			client.Lock()
			client.HeartbeatDate = time.Now().UnixMilli()
			client.Unlock()

			packet := HeartbeatPacket{
				Operation: OperationCode_Heartbeat,
				Data:      client.LastSequence,
			}

			client.Websocket.WriteJSON(packet)
		}
	}
}

func (client *Client) Resume() {
	for {
		date := client.ResumeDate
		now := time.Now().UnixMilli()
		if int(now)-int(date) < client.HeartbeatInterval+1000*3 {
			continue
		} else {
			client.Lock()
			client.ResumeDate = time.Now().UnixMilli()
			client.Unlock()

			packet := ResumePacket{
				Operation: OperationCode(OperationCode_Resume),
				Data: ResumePacketData{
					Token:     client.Token,
					SessionId: client.SessionId,
					Sequence:  client.LastSequence,
				},
			}

			client.Websocket.WriteJSON(packet)
		}
	}
}

func (client *Client) HandleMessage(msgType int, msg []byte) {
	var event *BaseEvent
	if err := json.NewDecoder(bytes.NewBuffer(msg)).Decode(&event); err != nil {
		log.Fatalf("error decoding base message \n%s", err)
	}
	client.Lock()
	defer client.Unlock()

	if event.Operation != 11 {
		client.LastSequence = event.Sequence
	}

	switch event.Operation {
	case OperationCode(OperationCode_Dispatch):
		client.HandleEvent(event)
		return
	case OperationCode(OperationCode_Hello):

		var data *HelloData

		if err := json.Unmarshal([]byte(event.Data), &data); err != nil {
			log.Fatalf("error decoding hello event data \n%s", err)
		}

		client.HeartbeatInterval = data.Heartbeat
		client.ResumeDate = time.Now().UnixMilli()
		return
	case OperationCode(OperationCode_HeartbeatACK):
		client.ResumeDate = time.Now().UnixMilli()
		return
	}
}

func (client *Client) HandleEvent(event *BaseEvent) {
	switch event.Type {
	case EventType_Ready:
		var data *Ready
		if err := json.Unmarshal(event.Data, &data); err != nil {
			log.Fatalf("error decoding ready event data \n%s", err)
		}
		client.SessionId = data.SessionId
		client.ResumeGatewayUrl = data.ResumeGatewayUrl

		events.Ready(data)
		return
	case EventType_MessageCreate:
		var msg *Message
		if err := json.Unmarshal(event.Data, &msg); err != nil {
			log.Fatalf("error decoding message create event data \n%s", err)
		}

		events.MessageCreate(msg)
		return
	case EventType_InteractionCreate:
		var interaction *Interaction
		if err := json.Unmarshal(event.Data, &interaction); err != nil {
			log.Fatalf("error decoding guild create event data")
		}

		events.InteractionCreate(interaction)
	case EventType_GuildCreate:
		var guild *Guild
		if err := json.Unmarshal(event.Data, &guild); err != nil {
			log.Fatalf("error decoding guild create event data")
		}

		client.Guilds[guild.Id] = guild
		for _, channel := range guild.Channels {
			client.Channels[channel.Id] = channel
		}
		return
	}
}

func (client *Client) listen() {
	for {
		msgType, msg, err := client.Websocket.ReadMessage()
		if err != nil {
			log.Fatalf("error listening message \n%s", err)
		}

		client.HandleMessage(msgType, msg)
	}
}
