package packets

import (
	"bytes"
	"discordbot/src/rest"
	. "discordbot/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"
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
	Id             Snowflake                                           `json:"id"`
	ApplicationId  Snowflake                                           `json:"application_id"`
	Type           InteractionType                                     `json:"type"`
	Data           InteractionData                                     `json:"data"`
	GuildId        Snowflake                                           `json:"guild_id"`
	ChannelId      Snowflake                                           `json:"channel_id"`
	Member         *GuildMember                                        `json:"member"`
	User           *User                                               `json:"user"`
	Token          string                                              `json:"token"`
	Version        int                                                 `json:"version"`
	Message        *Message                                            `json:"message"`
	AppPermissions string                                              `json:"app_permissions"`
	Locale         string                                              `json:"locale"`
	GuildLocale    string                                              `json:"guild_locale"`
	Options        map[string]*ApplicationCommandInteractionDataOption `json:"-"`
}

func (interaction *Interaction) Reply(data *MessagePacket) []byte {
	response := InteractionResponse{
		Type: InteractionCallbackType_ChannelMessageWithSource,
		Data: *data,
	}
	bytes := rest.Post(fmt.Sprintf("/interactions/%s/%s/callback", interaction.Id, interaction.Token), response)
	return bytes
}

func (interaction *Interaction) FetchReply() *Message {
	var message *Message
	data := rest.Get(fmt.Sprintf("/webhooks/%s/%s/messages/@original", interaction.ApplicationId, interaction.Token))

	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&message); err != nil {
		log.Fatalf("error fetching interactin reply decode message \n%s\n", err)
	}

	return message
}

func (interaction *Interaction) ErrorReply(data *MessagePacket) {
	interaction.Reply(data)
	message := interaction.FetchReply()

	time.Sleep(time.Second * 3)

	message.Delete()
}

func (interaction *Interaction) Parse() {
	interaction.Options = make(map[string]*ApplicationCommandInteractionDataOption)
	for _, val := range interaction.Data.Options {
		interaction.Options[val.Name] = val
	}
	interaction.User = interaction.Member.User
	interaction.Data.Options = nil
	interaction.Member.User = nil
}
