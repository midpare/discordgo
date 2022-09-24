package packets

import . "discordbot/src/utils"

type Emoji struct {
	Id            Snowflake `json:"id"`
	Name          string    `json:"name"`
	Roles         []*Role   `json:"roles"`
	User          *User     `json:"user"`
	RequireColons bool      `json:"require_colons"`
	Managed       bool      `json:"managed"`
	Animated      bool      `json:"animated"`
	Available     bool      `json:"available"`
}
