package events

import (
	"discordbot/src/commands"
	. "discordbot/src/global"
	"discordbot/src/packets"
	"log"
)

func InteractionCreate(interaction *packets.Interaction) {
	event, e := commands.Commands[interaction.Data.Name]
	if !e {
		log.Fatalf("error get commands interaction eventName: %s\n", interaction.Data.Name)
	}

	Id := interaction.User.Id
	guildId := interaction.GuildId
	message := &packets.MessagePacket{}

	if event.Category == "도박" {
		_, e := Global.Database.Gambling[guildId][Id]
		if event.Name != "가입" && !e {
			message.Content = "가입되지 않은 유저입니다"
			return
		}
	}

	message, success := event.Execute(interaction)
	if !success {
		interaction.ErrorReply(message)
	} else {
		interaction.Reply(message)
	}
}
