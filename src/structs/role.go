package structs

import . "discordbot/src/utils"

type Role struct {
	Id            Snowflake `json:"id"`
	Name          string    `json:"name"`
	Color         int       `json:"color"`
	Hoist         bool      `json:"hoist"`
	Icon          string    `json:"icon"`
	Unicode_emoji string    `json:"unicode_emoji"`
	Position      int       `json:"position"`
	Permissions   string    `json:"permissions"`
	Managed       bool      `json:"managed"`
	Mentionable   bool      `json:"mentionable"`
	Tags          *RoleTags `json:"tags"`
}

type RoleTags struct {
	BotId         Snowflake `json:"bot_id"`
	IntegrationId Snowflake `json:"integration_id"`
}
