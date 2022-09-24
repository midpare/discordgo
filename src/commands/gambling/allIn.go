package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"
	"math/rand"

	. "golang.org/x/text/message"
)

var AllIn = &packets.ApplicationCommand{
	Name:        "올인",
	Category:    "도박",
	Usage:       "올인",
	Description: "자신의 모든 돈을 걸고 도박을 합니다.",
	Execute: func(interaction *packets.Interaction, p *Printer) (message *packets.MessagePacket, success bool) {
		message = &packets.MessagePacket{}
		guildId := interaction.GuildId
		id := interaction.User.Id

		user := Global.Database.Gambling[guildId][id]

		random := rand.Intn(2)

		switch random {
		case 0:
			message.Content = p.Sprintf("Gambling succeeded! %d has been paid.\nbalance: %d -> %d", user.Money, user.Money, user.Money*2)
			user.Money *= 2
		case 1:
			message.Content = p.Sprintf("Gambling failed! %d has been deducated.\nbalance: %d -> %d", user.Money, user.Money, 0)
			user.Money = 0
		}
		Global.Database.Gambling[guildId][id] = user
		return message, true
	},
}
