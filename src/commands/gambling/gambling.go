package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"
	"math/rand"

	. "golang.org/x/text/message"
)

var Gambling = &packets.ApplicationCommand{
	Name:        "도박",
	Category:    "도박",
	Usage:       "도박 <돈>",
	Description: "자신의 <돈>을 걸고 도박을 진행합니다. (성공시: 2배, 실패시: 0배)",
	Options: []*packets.ApplicationCommandOption{
		{
			Name:        "돈",
			Description: "도박에 사용할 돈을 입력합니다.",
			Type:        packets.ApplicationCommandOptionType_Integer,
			Required:    true,
			MinValue:    1,
		},
	},
	Execute: func(interaction *packets.Interaction, p *Printer) (message *packets.MessagePacket, success bool) {
		id := interaction.User.Id
		guildId := interaction.GuildId

		message = &packets.MessagePacket{}

		user := Global.Database.Gambling[guildId][id]
		money := interaction.Options["돈"].Value

		if money > user.Money {
			message.Content = p.Sprintf("Cannot enter money which is higher than your balance.\nbalance: %d", user.Money)
			return message, false
		}

		random := rand.Intn(2)

		switch random {
		case 0:
			message.Content = p.Sprintf("Gambling succeeded! %d has been paid.\nbalance: %d -> %d", money, user.Money, user.Money+money)
			user.Money += money
		case 1:
			message.Content = p.Sprintf("Gambling failed! %d has been deducated.\nbalance: %d -> %d", money, user.Money, user.Money-money)
			user.Money -= money
		}
		Global.Database.Gambling[guildId][id] = user
		return message, true
	},
}
