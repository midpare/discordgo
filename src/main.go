package main

import (
	"discordbot/src/client"
	"log"
)

func main() {
	client, err := client.NewClient()

	if err != nil {
		log.Fatalf("error connecting to discord gateway\n%s\n", err)
	}

	client.Login()
}
