package structs

import (
	"discordbot/src/rest"
	. "discordbot/src/utils"
	"encoding/json"
	"fmt"
)

type MessageType uint
type MessageActivityType uint
type MessageFlags uint

const (
	MessageType_Default MessageType = iota
	MessageType_RecipientAdd
	MessageType_RecipientRemove
	MessageType_Call
	MessageType_ChannelNameChange
	MessageType_ChannelIconChange
	MessageType_ChannelPinnedMessage
	MessageType_UserJoin
	MessageType_GuildBoost
	MessageType_GuildBoostTier1
	MessageType_GuildBoostTier2
	MessageType_GuildBoostTier3
	MessageType_ChannelFollowAdd
	MessageType_GuildlDiscoveryDisqualified MessageType = iota + 1
	MessageType_GuildDiscoveryRequalified
	MessageType_GuildDiscoveryGracePeriodInitialWarning
	MessageType_GuildDiscoveryGracePeriodFinalWarning
	MessageType_ThreadCreated
	MessageType_Reply
	MessageType_ChatInputCommand
	MessageType_ThreadStarterMessage
	MessageType_GuildInviteReminder
	MessageType_ContextMenuCommand
	MessageType_AutoModerationAction
)

const (
	MessageActivityType_Join MessageActivityType = iota + 1
	MessageActivityType_Spectate
	MessageActivityType_Listem
	MessageActivityType_JoinRequest
)

const (
	MessageFlags_CrossPosted MessageFlags = 1 << iota
	MessageFlags_IsCrosspost
	MessageFlags_SuppressEmbeds
	MessageFlags_SourceMessageDeleted
	MessageFlags_Urgent
	MessageFlags_HasThread
	MessageFlags_Ephemeral
	MessageFlags_Loading
	MessageFlags_FailedToMentionSomerolesInThread
)

type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyId string              `json:"party_id"`
}

type MessageReference struct {
	MessageId       Snowflake `json:"message_id"`
	ChannelId       Snowflake `json:"channel_id"`
	GuildId         Snowflake `json:"guild_id"`
	FailIfNotExists bool      `json:"fail_if_not_exists"`
}

type MessageInteraction struct {
	Id     Snowflake       `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   User            `json:"user"`
	Member GuildMember     `json:"member"`
}

type Message struct {
	Id                Snowflake           `json:"id"`
	ChannelId         Snowflake           `json:"channel_id"`
	GuildId           Snowflake           `json:"guild_id"`
	Author            User                `json:"author"`
	Content           string              `json:"content"`
	Timestamp         string              `json:"timestamp"`
	EditedTimestamp   string              `json:"edited_timestamp"`
	Tts               bool                `json:"tts"`
	MentionEveryone   bool                `json:"mention_everyone"`
	Mentions          []User              `json:"mentions"`
	MentionRoles      []Role              `json:"mention_roles"`
	MentionChannel    []channelMentions   `json:"mention_channels"`
	Member            *GuildMember        `json:"member"`
	Attachments       []Attachment        `json:"attachments"`
	Embeds            []Embed             `json:"embeds"`
	Reactions         []Reaction          `json:"reactions"`
	Nonce             string              `json:"nonce"`
	Pinned            bool                `json:"pinned"`
	WebhookId         Snowflake           `json:"webhook_id"`
	Type              *MessageType        `json:"type"`
	Activity          *MessageActivity    `json:"activity"`
	Application       *Application        `json:"application"`
	ApplicationId     Snowflake           `json:"application_id"`
	MessageReference  *MessageReference   `json:"message_reference"`
	Flags             MessageFlags        `json:"flags"`
	ReferencedMessage *Message            `json:"referenced_message"`
	Interaction       *MessageInteraction `json:"interaction"`
	Thread            *Channel            `json:"thread"`
	StickerItems      []*StickerItem      `json:"sticker_items"`
	Stickers          []*Sticker          `json:"stickers"`
	Position          int                 `json:"position"`
}

func (msg *Message) Send(packet *MessagePacket) (*Message, error) {
	bytes := rest.Post(fmt.Sprintf("/channels/%s/messages", msg.ChannelId), packet)

	var data *Message
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}
