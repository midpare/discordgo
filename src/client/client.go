package client

import (
	"bytes"
	"discordbot/src/events"
	. "discordbot/src/global"
	"discordbot/src/models"
	. "discordbot/src/packets"
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
	HeartbeatInterval time.Duration
	ReceiveHeartbeat  chan bool
	IsConnect         chan bool
	SessionId         string
	ResumeGatewayUrl  string
	Guilds            map[utils.Snowflake]*Guild
	Channels          map[utils.Snowflake]*Channel
}

func NewClient() Client {
	return Client{
		Token:            os.Getenv("DISCORD_TOKEN2"),
		ReceiveHeartbeat: make(chan bool),
		Guilds:           make(map[utils.Snowflake]*Guild),
		Channels:         make(map[utils.Snowflake]*Channel),
	}
}

func (client *Client) Login() {
	dialer := websocket.DefaultDialer
	header := http.Header{}
	header.Add("accept-encoding", "json")
	websocket, _, err := dialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", header)

	if err != nil {
		log.Fatalf("error connecting to discord gateway\n%s\n", err)
	}

	defer websocket.Close()

	client.Lock()
	client.Websocket = websocket
	client.Unlock()

	messageType, message, err := client.Websocket.ReadMessage()

	if err != nil {
		log.Fatalf("error reading hello message \n%s", err)
	}

	client.handleMessage(messageType, message)

	go client.heartbeat()
	go client.resume()

	packet := Identify{
		Operation: OperationCode_Identify,
		Data: IdentifyData{
			Token:   client.Token,
			Intents: IntentsAll,
			Properties: IdentifyDataProperties{
				Os:      "window",
				Browser: "discord",
				Device:  "discord",
			},
		},
	}

	client.Websocket.WriteJSON(packet)

	client.listen()
}

func (client *Client) heartbeat() {
	ticker := time.NewTicker(client.HeartbeatInterval / 3)

	for {
		select {
		case <-ticker.C:
			packet := Heartbeat{
				Operation: OperationCode_Heartbeat,
				Data:      client.LastSequence,
			}

			client.Websocket.WriteJSON(packet)
		}
	}
}

func (client *Client) resume() {
	timer := time.NewTimer(client.HeartbeatInterval + (time.Second * 3))

	for {
		select {
		case <-timer.C:
			log.Println("connect failed!")
		case <-client.ReceiveHeartbeat:
			timer.Stop()
			timer = time.NewTimer(client.HeartbeatInterval + (time.Second * 3))
		}
	}
}

func (client *Client) handleMessage(messageType int, message []byte) {
	var event *BaseEvent
	if err := json.NewDecoder(bytes.NewBuffer(message)).Decode(&event); err != nil {
		log.Fatalf("error decoding base message \n%s", err)
	}
	client.Lock()
	defer client.Unlock()

	if event.Operation != 11 {
		client.LastSequence = event.Sequence
	}

	switch event.Operation {
	case OperationCode(OperationCode_Dispatch):
		client.handleEvent(event)
	case OperationCode(OperationCode_Hello):

		var data *HelloData

		if err := json.Unmarshal([]byte(event.Data), &data); err != nil {
			log.Fatalf("error decoding hello event data \n%s", err)
		}

		client.HeartbeatInterval = data.Heartbeat * time.Millisecond
	case OperationCode(OperationCode_HeartbeatACK):
		client.ReceiveHeartbeat <- true
	}
}

func (client *Client) handleEvent(event *BaseEvent) {
	switch event.Type {
	case EventType_Ready:
		var data *Ready
		if err := json.Unmarshal(event.Data, &data); err != nil {
			log.Fatalf("error decoding ready event data \n%s", err)
		}
		client.SessionId = data.SessionId
		client.ResumeGatewayUrl = data.ResumeGatewayUrl

		events.Ready(data)
	case EventType_MessageCreate:
		var message *Message
		if err := json.Unmarshal(event.Data, &message); err != nil {
			log.Fatalf("error decoding message create event data \n%s", err)
		}

		events.MessageCreate(message)
	case EventType_InteractionCreate:
		var interaction *Interaction
		if err := json.Unmarshal(event.Data, &interaction); err != nil {
			log.Fatalf("error decoding guild create event data \n%s", err)
		}
		interaction.Parse()
		events.InteractionCreate(interaction)
	case EventType_GuildCreate:
		var guild *Guild
		if err := json.Unmarshal(event.Data, &guild); err != nil {
			log.Fatalf("error decoding guild create event data \n%s", err)
		}

		client.Guilds[guild.Id] = guild
		Global.Database.Gambling[guild.Id] = make(map[utils.Snowflake]*models.Gambling)
		for _, channel := range guild.Channels {
			client.Channels[channel.Id] = channel
		}
	}
}

func (client *Client) listen() {
	for {
		messageType, message, err := client.Websocket.ReadMessage()
		if err != nil {
			log.Fatalf("error listening message \n%s", err)
		}

		client.handleMessage(messageType, message)
	}
}
