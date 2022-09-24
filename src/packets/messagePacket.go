package packets

import (
	. "discordbot/src/utils"
)

type AllowedMentionType string

const (
	AllowedMentionType_Role     = "roles"
	AllowedMentionType_User     = "users"
	AllowedMentionType_Everyone = "everyone"
)

type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`
	Roles       []Snowflake          `json:"roles"`
	Users       []Snowflake          `json:"users"`
	RepliedUser bool                 `json:"replied_user"`
}

type MessagePacket struct {
	Content          string             `json:"content"`
	Tts              bool               `json:"tts"`
	Embeds           []*Embed           `json:"embeds"`
	AllowedMentions  *AllowedMentions   `json:"allowed_mentions"`
	MessageReference *MessageReference  `json:"message_reference"`
	Components       []MessageComponent `json:"components"`
	StickerIds       []Snowflake        `json:"sticker_ids"`
	PayloadJson      string             `json:"payload_json"`
	Attachments      []*Attachment      `json:"attachments"`
	Flags            *MessageFlags      `json:"flags"`
}
