package structs

import (
	. "discordbot/src/utils"
)

type StageInstancePrivacyLevel uint

const (
	Public StageInstancePrivacyLevel = iota + 1
	GuildOnly
)

type StageInstance struct {
	Id                    Snowflake    `json:"id"`
	GuildId               Snowflake    `json:"guild_id"`
	ChannelId             Snowflake    `json:"channel_id"`
	Topic                 string       `json:"topic"`
	PrivacyLevel          StageInstancePrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool         `json:"discoverable_disabled"`
	GuildScheduledEventId Snowflake    `json:"guild_scheduled_event_id"`
}
