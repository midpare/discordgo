package structs

import "encoding/json"

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

	IntentsAllWithoutPrivileged = IntentGuilds |
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
		IntentAutoModerationExecution

	IntentsAll = IntentsAllWithoutPrivileged |
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

type IdentifyPacket struct {
	Operation OperationCode      `json:"op"`
	Data      IdentifyPacketData `json:"d"`
}

type IdentifyPacketData struct {
	Token      string                       `json:"token"`
	Intents    uint64                       `json:"intents"`
	Properties IdentifyPacketDataProperties `json:"properties"`
}

type IdentifyPacketDataProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type ResumePacket struct {
	Operation OperationCode    `json:"op"`
	Data      ResumePacketData `json:"d"`
}

type ResumePacketData struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Sequence  int    `json:"seq"`
}

type HeartbeatPacket struct {
	Operation OperationCode `json:"op"`
	Data      int           `json:"d"`
}

type HelloData struct {
	Heartbeat int `json:"heartbeat_interval"`
}
