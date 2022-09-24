package commands

import (
	gambling "discordbot/src/commands/gambling"
	normal "discordbot/src/commands/normal"

	"discordbot/src/packets"
)

var Commands = make(map[string]*packets.ApplicationCommand)

func Handle() []*packets.ApplicationCommand {
	for k, v := range gambling.Gambling_commands {
		Commands[k] = v
	}

	for k, v := range normal.Normal_commands {
		Commands[k] = v
	}

	arr := []*packets.ApplicationCommand{}

	for _, val := range Commands {
		arr = append(arr, val)
	}

	return arr
}
