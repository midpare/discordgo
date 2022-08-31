package commands

import (
	"discordbot/src/structs"
)

var ping *structs.ApplicationCommand = &structs.ApplicationCommand{
	Name:        "ping",
	Description: "test description",
	Type:        structs.ApplicationCommandType_ChatInput,
	Execute: func(interaction *structs.Interaction) {
		interaction.Reply(&structs.MessagePacket{
			Content: "pong!",
		})
	},
}
