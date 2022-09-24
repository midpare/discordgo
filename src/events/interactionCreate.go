package events

import (
	"discordbot/src/commands"
	. "discordbot/src/global"
	"discordbot/src/packets"
	"log"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func InteractionCreate(interaction *packets.Interaction) {
	p := message.NewPrinter(language.Korean)
	message := &packets.MessagePacket{}

	if interaction.GuildId == "" {
		message.Content = "명령어는 서버에서만 사용할 수 있습니다."
		interaction.Reply(message)
		return
	}

	event, e := commands.Commands[interaction.Data.Name]
	if !e {
		log.Fatalf("error get commands interaction eventName: %s\n", interaction.Data.Name)
	}

	Id := interaction.User.Id
	guildId := interaction.GuildId

	if event.Category == "도박" {
		_, e := Global.Database.Gambling[guildId][Id]
		if event.Name != "가입" && !e {
			message.Content = p.Sprintf("You are not join in gambling")
			interaction.ErrorReply(message)
			return
		}
	}

	message, success := event.Execute(interaction, p)
	if !success {
		interaction.ErrorReply(message)
	} else {
		interaction.Reply(message)
	}
}
