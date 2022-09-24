package commands

import (
	. "discordbot/src/global"
	"discordbot/src/models"
	"discordbot/src/packets"
	. "discordbot/src/utils"
)

var Join = &packets.ApplicationCommand{
	Name:        "가입",
	Category:    "도박",
	Usage:       "가입",
	Description: "도박 관련 명령어를 사용할수있게 가입을 합니다.",
	Execute: func(interaction *packets.Interaction) (message *packets.MessagePacket, success bool) {
		collection := Global.Database.Gambling

		id := interaction.User.Id
		guildId := interaction.GuildId
		message = &packets.MessagePacket{}

		_, e := collection[guildId][id]
		if e {
			message.Content = "이미 가입된 유저입니다."
			return message, false
		}

		_, e = collection[guildId]
		if !e {
			collection[guildId] = make(map[Snowflake]models.Gambling)
		}

		collection[guildId][id] = models.Gambling{
			Id:      id,
			GuildId: guildId,
		}

		message.Content = "성공적으로 가입되었습니다!"
		return message, true
	},
}
