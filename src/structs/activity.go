package structs

import (
	. "discordbot/src/utils"
)

type ActivityType uint
type ActivityFlags uint

const (
	Activity_Game ActivityType = iota
	Activity_Streaming
	Activity_Listening
	Activity_Watching
	Activity_Custom
	Activity_Competing
)

const (
	ActivityFlags_Instance ActivityFlags = 1 << iota
	ActivityFlags_Join
	ActivityFlags_Spectate
	ActivityFlags_JoinRequest
	ActivityFlags_Sync
	ActivityFlags_Play
	ActivityFlags_PartyPrivacyFriends
	ActivityFlags_PartyPrivacyVoiceChannel
	ActivityFlags_Embedded
)

type ActivityTimestamps struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type ActivityParty struct {
	Id   string `json:"id"`
	Size [2]int `json:"size"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
}

type ActivitySecrets struct {
	Join     string `json:"join"`
	Spectate string `json:"spectate"`
	Match    string `json:"match"`
}

type ActivityButton struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type Activity struct {
	Name          string              `json:"name"`
	Type          ActivityType        `json:"type"`
	Url           string              `json:"url"`
	CreatedAt     int                 `json:"created_at"`
	Timestamps    *ActivityTimestamps `json:"timestamps"`
	ApplicationId Snowflake           `json:"application_id"`
	Details       string              `json:"details"`
	State         string              `json:"state"`
	Emoji         *Emoji              `json:"emoji"`
	Party         *ActivityParty      `json:"party"`
	Assets        *ActivityAssets     `json:"assets"`
	Secrets       *ActivitySecrets    `json:"secrets"`
	Instance      bool                `json:"instance"`
	Flags         ActivityFlags       `json:"flags"`
	Buttons       []*ActivityButton   `json:"buttons"`
}
