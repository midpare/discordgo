package packets

import (
	. "discordbot/src/utils"
)

type VoiceState struct {
	GuildId                 Snowflake    `json:"guild_id"`
	ChannelId               Snowflake    `json:"channel_id"`
	UserId                  Snowflake    `json:"user_id"`
	Member                  *GuildMember `json:"member"`
	SessionId               string       `json:"session_id"`
	Deaf                    bool         `json:"deaf"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              bool         `json:"self_stream"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	RequestToSpeakTimestamp string       `json:"request_to_speak_timestamp"`
}
