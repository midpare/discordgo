package packets

import (
	. "discordbot/src/utils"
)

type GuildScheduledEventPrivacyLevel uint
type GuildScheduledEventStatus uint
type GuildScheduledEventEntityTypes uint

const (
	GuildScheduledEvent_GuildOnly GuildScheduledEventPrivacyLevel = iota + 2
)

const (
	GuildScheduledEventStatus_Scheduled GuildScheduledEventStatus = iota + 1
	GuildScheduledEventStatus_Active
	GuildScheduledEventStatus_Completed
	GuildScheduledEventStatus_Canceled
)

const (
	GuildScheduledEventEntityTypes_StageInstance GuildScheduledEventEntityTypes = iota + 1
	GuildScheduledEventEntityTypes_Voice
	GuildScheduledEventEntityTypes_External
)

type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location"`
}

type GuildScheduledEvent struct {
	Id                 Snowflake
	GuildId            Snowflake
	ChannelId          Snowflake
	CreatorId          Snowflake
	Name               string
	Description        string
	ScheduledStartTime string
	ScheduledEndTime   string
	PrivacyLevel       GuildScheduledEventPrivacyLevel
	Status             GuildScheduledEventStatus
	EntityType         GuildScheduledEventEntityTypes
	EntityId           Snowflake
	EntityMetadata     *GuildScheduledEventEntityMetadata
	Creator            *User
	UserCount          int
	image              string
}
