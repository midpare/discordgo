package commands

import "discordbot/src/packets"

var Normal_commands = make(map[string]*packets.ApplicationCommand)

func init() {
	Normal_commands[Ping.Name] = Ping
}
