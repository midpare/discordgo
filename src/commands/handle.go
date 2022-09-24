package commands

import (
	. "discordbot/src/commands/gambling"
	. "discordbot/src/commands/normal"
	"discordbot/src/packets"
)

var Commands = make(map[string]*packets.ApplicationCommand)

func Handle() []*packets.ApplicationCommand {
	Commands[Ping.Name] = Ping
	Commands[Gambling.Name] = Gambling
	Commands[Join.Name] = Join
	arr := []*packets.ApplicationCommand{}

	for _, val := range Commands {
		arr = append(arr, val)
	}

	return arr
}
