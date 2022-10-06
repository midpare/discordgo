package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"

. "golang.org/x/text/message"
)

var Debt = &packets.ApplicationCommand{
	Name: "빚",
	Category: "도박",
	Usage: "빚",
	Description: "현재 자신의 빚을 확인합니다.",
	Execute: func(interaction *packets.Interaction, p *Printer) (message *packets.MessagePacket, success bool) {
		message = &packets.MessagePacket{}
		guildId := interaction.GuildId
		id := interaction.User.Id

		user := Global.Database.Gambling[guildId][id]

		message.Content = p.Sprintf("%s's debt is %d.", interaction.Member.DisplayName, user.Debt)
		return message, true
	},
}