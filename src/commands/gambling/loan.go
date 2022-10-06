package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"

	. "golang.org/x/text/message"
)

var Loan = &packets.ApplicationCommand{
	Name:        "대출",
	Category:    "도박",
	Usage:       "대출 <돈>",
	Description: "최대 100만원까지의 돈을 대출합니다.",
	Options: []*packets.ApplicationCommandOption{
		{
			Name:        "돈",
			Description: "대출할 돈을 입력합니다.",
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
		debt := interaction.Options["돈"].Value

		if user.Debt+debt > 1000000 {
			message.Content = p.Sprintf("You can't borrow more than 1 million.")
			return message, false
		}

		user.Money += debt
		user.Debt += debt
		Global.Database.Gambling[guildId][id] = user

		message.Content = p.Sprintf("%d was successfully borrowed.\ndebt: %d -> %d", debt, user.Debt-debt, user.Debt)
		return message, true
	},
}
