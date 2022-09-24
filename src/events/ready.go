package events

import (
	"discordbot/src/commands"
	"discordbot/src/packets"
	"discordbot/src/rest"
	"fmt"
	"log"
)

func Ready(data *packets.Ready) {

	log.Println("success to login!")
	arr := commands.Handle()
	rest.Put(fmt.Sprintf("/applications/%s/commands", data.User.Id), arr)
}
