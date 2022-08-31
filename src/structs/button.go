package structs

type ButtonStyle uint

const (
	ButtonStyle_Primary ButtonStyle = iota + 1
	ButtonStyle_Secondary
	ButtonStyle_Success
	ButtonStyle_Danger
	ButtonStyle_Link
)

type Button struct {
	BaseComponent
	Style    ButtonStyle `json:"style"`
	Label    string      `json:"label"`
	Emoji    Emoji       `json:"emoji"`
	Url      string      `json:"url"`
	Disabled bool        `json:"disabled"`
}
