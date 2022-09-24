package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"
	"fmt"
	"math/rand"
)

var money = packets.ApplicationCommandOption{
	Name:        "돈",
	Description: "도박에 사용할 돈을 입력합니다.",
	Type:        packets.ApplicationCommandOptionType_Integer,
	Required:    true,
	MinValue:    1,
}

var Gambling = &packets.ApplicationCommand{
	Name:        "도박",
	Category:    "도박",
	Usage:       "도박 <돈>",
	Description: "자신의 <돈>을 걸고 도박을 진행합니다. (성공시: 2배, 실패시: 0배)",
	Options: []*packets.ApplicationCommandOption{
		&money,
	},
	Execute: func(interaction *packets.Interaction) (message *packets.MessagePacket, success bool) {
		collection := Global.Database.Gambling
		id := interaction.User.Id
		guildId := interaction.GuildId

		message = &packets.MessagePacket{}

		user := collection[guildId][id]
		money := interaction.Options["돈"].Value

		if money > user.Money {
			message.Content = fmt.Sprintf("현재 잔액보다 높은 돈은 입력할 수 없습니다.\n현재 잔액: %d원", user.Money)
			return message, false
		}

		random := rand.Intn(2)

		switch random {
		case 0:
			message.Content = fmt.Sprintf("도박에 성공했습니다! %d원이 지급되었습니다.\n현재 잔액: %d원 -> %d원", money, user.Money, user.Money+money)
			user.Money += money
			collection[guildId][id] = user
		case 1:
			message.Content = fmt.Sprintf("도박에 실패했습니다! %d원이 차감되었습니다.\n현재 잔액: %d원 -> %d원", money, user.Money, user.Money-money)
			user.Money -= money
			collection[guildId][id] = user
		}
		return message, true
	},
}
