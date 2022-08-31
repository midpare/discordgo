package events

import (
	"discordbot/src/commands"
	"discordbot/src/rest"
	"discordbot/src/structs"
	"fmt"
	"log"
)

func Ready(data *structs.Ready) {

	log.Println("success to login!")
	rest.Put(fmt.Sprintf("/applications/%s/commands", data.User.Id), commands.Handle())
}
