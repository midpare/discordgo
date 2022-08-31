package structs

import (
	. "discordbot/src/utils"
)

type ClientStatus struct {
	Desktop string `json:"desktop"`
	Mobile  string `json:"mobile"`
	Web     string `json:"web"`
}

type PresencesUpdate struct {
	User         *User         `json:"user"`
	GuildId      Snowflake     `json:"guild_id"`
	Status       string        `json:"status"`
	Activities   []*Activity   `json:"activities"`
	ClientStatus *ClientStatus `json:"client_status"`
}
