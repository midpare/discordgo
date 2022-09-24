package packets

type Embed struct {
	Title       string          `json:"tile"`
	Type        string          `json:"type"`
	Description string          `json:"description"`
	Url         string          `json:"url"`
	Timestamp   string          `json:"timestamp"`
	Color       int             `json:"color"`
	Footer      *EmbedFooter    `json:"footer"`
	Image       *EmbedImage     `json:"image"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail"`
	Video       *EmbedVideo     `json:"video"`
	Provider    *EmbedProvider  `json:"provider"`
	Author      *EmbedAuthor    `json:"author"`
	Fields      []*EmbedField   `json:"fields"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconUrl      string `json:"icon_url"`
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedImage struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedThumbnail struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedVideo struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	IconUrl      string `json:"icon_url"`
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
