package packets

import . "discordbot/src/utils"

type Attachment struct {
	Id          Snowflake `json:"id"`
	Filename    string    `json:"filename"`
	Description string    `json:"description"`
	ContentType string    `json:"content_type"`
	Size        int       `json:"size"`
	Url         string    `json:"url"`
	ProxyUrl    string    `json:"proxy_url"`
	Height      int       `json:"height"`
	Width       int       `json:"width"`
	Ephemeral   bool      `json:"ephemeral"`
}
