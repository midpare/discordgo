package structs

type SelectMenu struct {
	BaseComponent
	Options     []SelectMenuOptions `json:"options"`
	Placeholder string              `json:"placeholder"`
	MinValues   int                 `json:"min_values"`
	MaxValues   int                 `json:"max_values"`
	Disabled    bool                `json:"disabled"`
}

type SelectMenuOptions struct {
	Label       string
	Value       string
	Description string
	Emoji       Emoji
	Default     bool
}
