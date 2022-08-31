package events

import (
	"discordbot/src/commands"
	"discordbot/src/structs"
	"log"
)

func InteractionCreate(interaction *structs.Interaction) {
	event, e := commands.Commands[interaction.Data.Name]

	if !e {
		log.Fatalf("error get commands interaction eventName: %s\n", interaction.Data.Name)
	}

	event.Execute(interaction)
}
