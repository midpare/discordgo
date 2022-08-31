package events

import (
	"discordbot/src/structs"
	"fmt"
	"log"
)

func MessageCreate(msg *structs.Message) {
	if msg.Author.Bot {
		return
	}

	log.Fatalf(fmt.Sprintf("%s: %s", msg.Author.Username, msg.Content))
}
