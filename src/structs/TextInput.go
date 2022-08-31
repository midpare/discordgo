package structs

type TextInputStyle uint

const (
	Short TextInputStyle = iota + 1
	Paragraph
)

type TextInput struct {
	BaseComponent
	Style       TextInputStyle `json:"style"`
	Label       string         `json:"label"`
	MinLength   int            `json:"min_length"`
	MaxLength   int            `json:"max_length"`
	Required    bool           `json:"required"`
	Value       string         `json:"value"`
	Placeholder string         `json:"placeholder"`
}
