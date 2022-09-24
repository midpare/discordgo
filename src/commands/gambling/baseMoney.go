package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"

	. "golang.org/x/text/message"
)

var BaseMoney = &packets.ApplicationCommand{
	Name:        "기초자금",
	Category:    "도박",
	Usage:       "기초자금",
	Description: "기초자금 25,000원을 획득합니다. 돈이 0원일때만 명령어 사용이 가능합니다.",
	Execute: func(interaction *packets.Interaction, p *Printer) (message *packets.MessagePacket, success bool) {
		message = &packets.MessagePacket{}
		guildId := interaction.GuildId
		id := interaction.User.Id

		user := Global.Database.Gambling[guildId][id]

		if user.Money > 0 {
			message.Content = p.Sprintf("Cannot pay baseMoney because you have money or coin.")
			return message, false
		}

		user.Money = 25000
		Global.Database.Gambling[guildId][id] = user
		message.Content = p.Sprintf("%d was successfully paid!", 25000)
		return message, true
	},
}
