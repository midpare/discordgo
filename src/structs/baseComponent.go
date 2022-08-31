package structs

type BaseComponent struct {
	Type     ComponentType `json:"type"`
	CustomId string        `json:"custom_id"`
}

func (component *BaseComponent) setCustomId(id string) {
	component.CustomId = id
}
