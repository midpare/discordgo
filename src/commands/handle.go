package commands

import "discordbot/src/structs"

var Commands = make(map[string]*structs.ApplicationCommand)

func Handle() []*structs.ApplicationCommand {
	Commands[ping.Name] = ping

	arr := []*structs.ApplicationCommand{}

	for _, val := range Commands {
		arr = append(arr, val)
	}

	return arr
}
