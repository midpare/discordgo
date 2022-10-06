package main

import (
	"discordbot/src/client"
	"discordbot/src/language"
)

func main() {
	client := client.NewClient()

	language.SetUp()
	client.Login()
}
