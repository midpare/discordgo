package structs

import . "discordbot/src/utils"

type GuildMember struct {
	User                       *User       `json:"user"`
	Nick                       string      `json:"nick"`
	Avatar                     string      `json:"avatar"`
	Roles                      []Snowflake `json:"roles"`
	Joined_at                  string      `json:"joined_at"`
	PremiumSince               string      `json:"premium_since"`
	Deaf                       bool        `json:"deaf"`
	Mute                       bool        `json:"mute"`
	Pending                    bool        `json:"pending"`
	Permissions                string      `json:"permissions"`
	CommunicationDisabledUntil string      `json:"communication_disabled_until"`
}
