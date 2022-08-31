package structs

import (
	"discordbot/src/rest"
	. "discordbot/src/utils"
	"fmt"
)

type InteractionType uint
type InteractionCallbackType uint

const (
	InteractionType_Ping InteractionType = iota + 1
	InteractionType_ApplicationCommand
	InteractionType_MessageComponent
	InteractionType_ApplicationCommandAutocompleted
	InteractionType_ModalSubmit
)

const (
	InteractionCallbackType_Pong                     InteractionCallbackType = iota + 1
	InteractionCallbackType_ChannelMessageWithSource InteractionCallbackType = iota + 3
	InteractionCallbackType_DeferredChannelMessageWithSource
	InteractionCallbackType_DeferredUpdateMessage
	InteractionCallbackType_UpdateMessage
	InteractionCallbackType_ApplicationCommandAutocompleteResult
	InteractionCallbackType_Modal
)

type ResolvedData struct {
	Users      map[Snowflake]*User        `json:"users"`
	Members    map[Snowflake]*GuildMember `json:"members"`
	Roles      map[Snowflake]*Role        `json:"roles"`
	Channels   map[Snowflake]*Channel     `json:"channels"`
	Messages   map[Snowflake]*Message     `json:"messages"`
	Attachment map[Snowflake]*Attachment  `json:"attachment"`
}

type InteractionData struct {
	Id       Snowflake                                  `json:"id"`
	Name     string                                     `json:"name"`
	Type     int                                        `json:"type"`
	Resolved *ResolvedData                              `json:"resolved"`
	Options  []*ApplicationCommandInteractionDataOption `json:"options"`
	GuildId  Snowflake                                  `json:"guild_id"`
	TargetId Snowflake                                  `json:"target_id"`
}

type InteractionResponse struct {
	Type InteractionCallbackType `json:"type"`
	Data MessagePacket           `json:"data"`
}

type Interaction struct {
	Id             Snowflake       `json:"id"`
	ApplicationId  Snowflake       `json:"application_id"`
	Type           InteractionType `json:"type"`
	Data           InteractionData `json:"data"`
	GuildId        Snowflake       `json:"guild_id"`
	ChannelId      Snowflake       `json:"channel_id"`
	Member         *GuildMember    `json:"member"`
	User           *User           `json:"user"`
	Token          string          `json:"token"`
	Version        int             `json:"version"`
	Message        *Message        `json:"message"`
	AppPermissions string          `json:"app_permissions"`
	Locale         string          `json:"locale"`
	GuildLocale    string          `json:"guild_locale"`
}

func (interaction *Interaction) Reply(message *MessagePacket) []byte {
	response := InteractionResponse{
		Type: InteractionCallbackType_ChannelMessageWithSource,
		Data: *message,
	}
	bytes := rest.Post(fmt.Sprintf("/interactions/%s/%s/callback", interaction.Id, interaction.Token), response)
	return bytes
}
