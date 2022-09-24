package main

import (
	"discordbot/src/client"
	"discordbot/src/global"
	"log"
)

func main() {
	client, err := client.NewClient()

	if err != nil {
		log.Fatalf("error connecting to discord gateway\n%s\n", err)
	}
	global.SetGlobal()
	client.Login()
}
