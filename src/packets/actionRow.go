package packets

type ComponentType uint

const (
	ComponentType_ActionRow ComponentType = iota + 1
	ComponentType_Button
	ComponentType_SelectMenu
	ComponentType_TextInput
)

type MessageComponent interface {
	setCustomId(id string)
}

type ActionRow struct {
	Type       ComponentType       `json:"type"`
	Components [5]MessageComponent `json:"components"`
}
