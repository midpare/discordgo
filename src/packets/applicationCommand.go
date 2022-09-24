package packets

import (
	. "discordbot/src/utils"
)

type Interaction_Execute func(*Interaction) (*MessagePacket, bool)

type ApplicationCommandType uint
type ApplicationCommandOptionType uint

const (
	ApplicationCommandType_ChatInput ApplicationCommandType = iota + 1
	ApplicationCommandType_User
	ApplicationCommandType_Message
)

const (
	ApplicationCommandOptionType_SubCommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionType_SubCommandGroup
	ApplicationCommandOptionType_String
	ApplicationCommandOptionType_Integer
	ApplicationCommandOptionType_Boolean
	ApplicationCommandOptionType_User
	ApplicationCommandOptionType_Channel
	ApplicationCommandOptionType_Role
	ApplicationCommandOptionType_Mentionable
	ApplicationCommandOptionType_Number
	ApplicationCommandOptionType_Attachment
)

type ApplicationCommandOptionChoice struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ApplicationCommandOption struct {
	Type         ApplicationCommandOptionType      `json:"type"`
	Name         string                            `json:"name"`
	Description  string                            `json:"description"`
	Required     bool                              `json:"required"`
	Choices      []*ApplicationCommandOptionChoice `json:"choices"`
	Options      []*ApplicationCommandOption       `json:"options"`
	ChannelTypes []Channeltype                     `json:"channel_types"`
	MinValue     int                               `json:"min_value,omitempty"`
	MaxValue     int                               `json:"max_value,omitempty"`
	Minlength    int                               `json:"min_length,omitempty"`
	MaxLength    int                               `json:"max_length,omitempty"`
	Autocomplete bool                              `json:"autocomplete,omitempty"`
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    ApplicationCommandOptionType               `json:"type"`
	Value   int                                        `json:"value"`
	Options []*ApplicationCommandInteractionDataOption `json:"options"`
	Focused bool                                       `json:"focuses"`
}

type ApplicationCommand struct {
	Id                       Snowflake                   `json:"id,omitempty"`
	Type                     ApplicationCommandType      `json:"type,omitempty"`
	ApplicationId            Snowflake                   `json:"application_id,omitempty"`
	GuildId                  Snowflake                   `json:"guild_id,omitempty"`
	Name                     string                      `json:"name"`
	Description              string                      `json:"description"`
	Options                  []*ApplicationCommandOption `json:"options"`
	DefaultMemberPermissions string                      `json:"default_member_permission,omitempty"`
	DmPermission             bool                        `json:"dm_permission,omitempty"`
	DefaultPermission        bool                        `json:"default_permission,omitempty"`
	Version                  Snowflake                   `json:"version,omitempty"`
	Category                 string                      `json:"-"`
	Usage                    string                      `json:"-"`
	Execute                  Interaction_Execute         `json:"-"`
}
