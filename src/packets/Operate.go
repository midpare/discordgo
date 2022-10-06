package packets

import (
	"encoding/json"
	"time"
)

type OperationCode uint
type Intent uint
type EventType string

const (
	OperationCode_Dispatch OperationCode = iota
	OperationCode_Heartbeat
	OperationCode_Identify
	OperationCode_PresenceUpdate
	OperationCode_VoiceStateUpdate
	OperationCode_Resume OperationCode = iota + 1
	OperationCode_Reconnect
	OperationCode_RequestGuildMembers
	OperationCode_InvalidSession
	OperationCode_Hello
	OperationCode_HeartbeatACK
)

const (
	IntentsNone                       Intent = iota
	IntentGuilds                      Intent = 1<<iota - 1
	IntentGuildMembers                Intent = 1 << iota
	IntentGuildBans                   Intent = 1 << iota
	IntentGuildEmojis                 Intent = 1 << iota
	IntentGuildIntegrations           Intent = 1 << iota
	IntentGuildWebhooks               Intent = 1 << iota
	IntentGuildInvites                Intent = 1 << iota
	IntentGuildVoiceStates            Intent = 1 << iota
	IntentGuildPresences              Intent = 1 << iota
	IntentGuildMessages               Intent = 1 << iota
	IntentGuildMessageReactions       Intent = 1 << iota
	IntentGuildMessageTyping          Intent = 1 << iota
	IntentDirectMessages              Intent = 1 << iota
	IntentDirectMessageReactions      Intent = 1 << iota
	IntentDirectMessageTyping         Intent = 1 << iota
	IntentMessageContent              Intent = 1 << iota
	IntentGuildScheduledEvents        Intent = 1 << iota
	IntentAutoModerationConfiguration Intent = 1<<iota + 3
	IntentAutoModerationExecution     Intent = 1 << iota

	IntentsAll = IntentGuilds |
		IntentGuildBans |
		IntentGuildEmojis |
		IntentGuildIntegrations |
		IntentGuildWebhooks |
		IntentGuildInvites |
		IntentGuildVoiceStates |
		IntentGuildMessages |
		IntentGuildMessageReactions |
		IntentGuildMessageTyping |
		IntentDirectMessages |
		IntentDirectMessageReactions |
		IntentDirectMessageTyping |
		IntentGuildScheduledEvents |
		IntentAutoModerationConfiguration |
		IntentAutoModerationExecution |
		IntentGuildMembers |
		IntentGuildPresences |
		IntentMessageContent
)

const (
	EventType_Ready             = EventType("READY")
	EventType_MessageCreate     = EventType("MESSAGE_CREATE")
	EventType_GuildCreate       = EventType("GUILD_CREATE")
	EventType_InteractionCreate = EventType("INTERACTION_CREATE")
)

type BaseEvent struct {
	Operation OperationCode   `json:"op"`
	Sequence  int             `json:"s"`
	Type      EventType       `json:"t"`
	Data      json.RawMessage `json:"d"`
}

type Identify struct {
	Operation OperationCode `json:"op"`
	Data      IdentifyData  `json:"d"`
}

type IdentifyData struct {
	Token      string                 `json:"token"`
	Intents    Intent                 `json:"intents"`
	Properties IdentifyDataProperties `json:"properties"`
}

type IdentifyDataProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type Resume struct {
	Operation OperationCode `json:"op"`
	Data      ResumeData    `json:"d"`
}

type ResumeData struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Sequence  int    `json:"seq"`
}

type Heartbeat struct {
	Operation OperationCode `json:"op"`
	Data      int           `json:"d"`
}

type HelloData struct {
	Heartbeat time.Duration `json:"heartbeat_interval"`
}
