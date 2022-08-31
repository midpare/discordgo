package structs

import . "discordbot/src/utils"

type channelMentions struct {
	Id      Snowflake `json:"id"`
	GuildId Snowflake `json:"guild_id"`
	Type    int       `json:"type"`
	Name    string    `json:"name"`
}
