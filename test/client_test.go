package client_test

import (
	"discordbot/src/client"
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client, err := client.NewClient()

	if err != nil {
		t.Errorf("error testing new Client \n%s\n", err)
	}

	time := time.Now().UnixMilli()
	if time-client.ResumeDate > 3000 {
		t.Errorf("expect client.ResumeDate = %d, but client.ResumeDate = %d\n", time, client.ResumeDate)
	}

	if client.Token != os.Getenv("DISCORD_TOKEN") {
		t.Errorf("Expect client.Token = %s, but client.Token = %s\n", os.Getenv("DISCORD_TOKEN"), client.Token)
	}

	if client.Websocket == nil {
		t.Error("Expect client.Websocket exist but it is nil\n")
	}

	client.Websocket.Close()
}
