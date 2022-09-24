package commands

import (
	packet "discordbot/src/packets"
)

var Ping = &packet.ApplicationCommand{
	Name:        "ping",
	Category:    "기본",
	Usage:       "ping",
	Description: "봇의 현재상태를 확인합니다",
	Execute: func(interaction *packet.Interaction) (message *packet.MessagePacket, success bool) {
		return &packet.MessagePacket{
			Content: "pong!",
		}, true
	},
}
