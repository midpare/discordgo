package packets

import (
	. "discordbot/src/utils"
)

type Channeltype uint

const (
	ChannelType_GuildText Channeltype = iota
	ChannelType_Dm
	ChannelType_GuildVoice
	ChannelType_GroupDm
	ChannelType_GuildCategory
	ChannelType_GuildNews
	ChannelType_GuildNewsThread Channeltype = iota + 4
	ChannelType_GuildPublicThread
	ChannelType_GuildPrivateThread
	ChannelType_GuildStageVoice
	ChannelType_GuildDirectory
	ChannelType_GuildForum
)

type Overwrite struct {
	Id    Snowflake `json:"id"`
	Type  int       `json:"type"`
	Allow string    `json:"allow"`
	Deny  string    `json:"deny"`
}

type ThreadMetadata struct {
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int    `json:"auto_archive_duration"`
	ArchiveTimestamp    string `json:"archive_timestamp"`
	Locked              bool   `json:"locked"`
	Invitable           bool   `json:"invitable"`
	CreateTimestamp     string `json:"create_timestamp"`
}

type ThreadMember struct {
	Id            Snowflake `json:"id"`
	UserId        Snowflake `json:"user_id"`
	JoinTimestamp string    `json:"join_timestamp"`
	Flags         int       `json:"flags"`
}

type Channel struct {
	Id                         Snowflake       `json:"id"`
	Type                       Channeltype     `json:"type"`
	GuildId                    Snowflake       `json:"guild_id"`
	Position                   int             `json:"position"`
	PermissionOverwrites       []*Overwrite    `json:"permission_overwrites"`
	Name                       string          `json:"name"`
	Topic                      string          `json:"topic"`
	NSFW                       bool            `json:"nsfw"`
	LastMessageId              Snowflake       `json:"last_message_id"`
	Bitrate                    int             `json:"bitrate"`
	UserLimit                  int             `json:"uesr_limit"`
	RateLimitPerUser           int             `json:"rate_limit_per_user"`
	Recipients                 []*User         `json:"recipients"`
	Icon                       string          `json:"icon"`
	OwnerId                    Snowflake       `json:"owner_id"`
	ApplicationId              Snowflake       `json:"application_id"`
	ParentId                   Snowflake       `json:"parent_id"`
	LastPinTimestamp           string          `json:"last_pin_timestamp"`
	RtcRegion                  string          `json:"rtc_region"`
	VideoQualityMode           int             `json:"video_quality_mode"`
	MessageCount               int             `json:"message_count"`
	MemberCount                int             `json:"member_count"`
	ThreadMetadata             *ThreadMetadata `json:"thread_metadata"`
	Member                     *ThreadMember   `json:"member"`
	DefaultAutoArchiveDuration int             `json:"default_auto_archive_duration"`
	Permissions                string          `json:"permissions"`
	Flags                      int             `json:"flags"`
	TotalMessageSent           int             `json:"total_message_sent"`
}
