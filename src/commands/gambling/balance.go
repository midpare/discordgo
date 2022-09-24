package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"

	. "golang.org/x/text/message"
)

var Balance = &packets.ApplicationCommand{
	Name:        "잔액",
	Category:    "도박",
	Usage:       "잔액",
	Description: "현재 잔액을 확인합니다.",
	Execute: func(interaction *packets.Interaction, p *Printer) (message *packets.MessagePacket, success bool) {
		message = &packets.MessagePacket{}
		guildId := interaction.GuildId
		id := interaction.User.Id

		user := Global.Database.Gambling[guildId][id]

		message.Content = p.Sprintf("%s's balance is %d.", interaction.Member.DisplayName, user.Money)
		return message, true
	},
}
