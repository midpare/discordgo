package commands

import (
	. "discordbot/src/global"
	"discordbot/src/packets"
	"fmt"
	"math/rand"
	"time"

	. "golang.org/x/text/message"
)

var Daily = &packets.ApplicationCommand{
	Name:        "출석체크",
	Category:    "도박",
	Usage:       "출석체크",
	Description: "하루에 한번 50,000 ~ 100,000원을 획득합니다.",
	Execute: func(interaction *packets.Interaction, p *Printer) (message *packets.MessagePacket, success bool) {
		message = &packets.MessagePacket{}
		guildId := interaction.GuildId
		id := interaction.User.Id

		user := Global.Database.Gambling[guildId][id]
		now := time.Now()
		format := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
		if user.DailyDate == format {
			message.Content = p.Sprintf("Already checked attendance today.")
			return message, false
		}

		random := rand.Intn(50000) + 50000

		user.Money += random
		user.DailyDate = format
		Global.Database.Gambling[guildId][id] = user

		message.Content = fmt.Sprintf("Attendance check was successful! %d has been paid.\nbalance: %d -> %d", random, user.Money, user.Money+random)
		return message, true
	},
}
