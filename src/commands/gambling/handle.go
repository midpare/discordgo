package commands

import "discordbot/src/packets"

var Gambling_commands = make(map[string]*packets.ApplicationCommand)

func init() {
	Gambling_commands[AllIn.Name] = AllIn
	Gambling_commands[Balance.Name] = Balance
	Gambling_commands[BaseMoney.Name] = BaseMoney
	Gambling_commands[Daily.Name] = Daily
	Gambling_commands[Gambling.Name] = Gambling
	Gambling_commands[Half.Name] = Half
	Gambling_commands[Join.Name] = Join
}