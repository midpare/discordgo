package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"

	. "golang.org/x/text/message"
)

var PayBack = &packets.ApplicationCommand{
	Name:        "빚갚기",
	Category:    "도박",
	Usage:       "빚갚기 <돈>",
	Description: "자신의 빚을 갚습니다.",
	Options: []*packets.ApplicationCommandOption{
		{
			Name:        "돈",
			Description: "갚을 돈을 입력합니다.",
			Type:        packets.ApplicationCommandOptionType_Integer,
			Required:    true,
			MinValue:    1,
		},
	},
	Execute: func(interaction *packets.Interaction, p *Printer) (message *packets.MessagePacket, success bool) {
		message = &packets.MessagePacket{}
		guildId := interaction.GuildId
		id := interaction.User.Id

		user := Global.Database.Gambling[guildId][id]

		money := interaction.Options["돈"].Value

		if money > user.Money {
			message.Content = p.Sprintf("Cannot enter money which is higher than your balance.\nbalance: %d", user.Money)
			return message, false
		}

		if user.Debt < money {
			message.Content = p.Sprintf("The money to pay back is more than current debt.")
			return message, false
		}

		user.Money -= money
		user.Debt -= money

		Global.Database.Gambling[guildId][id] = user

		message.Content = p.Sprintf("%d is successfully paid back!\ndebt: %d -> %d", money, user.Debt+money, user.Debt)
		return message, true
	},
}
