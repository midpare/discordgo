package events

import (
	"discordbot/src/packets"
	"fmt"
	"log"
)

func MessageCreate(message *packets.Message) {
	if message.Author.Bot {
		return
	}

	log.Printf(fmt.Sprintf("%s: %s", message.Author.Username, message.Content))
}
