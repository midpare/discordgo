package structs

import . "discordbot/src/utils"

type StickerFormat uint
type StickerType uint

const (
	StickerFormat_PNG StickerFormat = iota + 1
	StickerFormat_APNG
	StickerFormat_LOTTIE
)

const (
	StickerType_Standard StickerType = iota + 1
	StickerType_Guild
)

type Sticker struct {
	Id          Snowflake     `json:"id"`
	PackId      Snowflake     `json:"pack_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Tags        string        `json:"tags"`
	Asset       string        `json:"asset"`
	Type        StickerType   `json:"type"`
	FormatType  StickerFormat `json:"format_type"`
	Available   bool          `json:"available"`
	GuildId     Snowflake     `json:"guild_id"`
	User        *User         `json:"user"`
	SortValue   int           `json:"sort_value"`
}

type StickerItem struct {
	Id         Snowflake     `json:"id"`
	Name       string        `json:"name"`
	FormatType StickerFormat `json:"format_type"`
}
