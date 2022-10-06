package test

import (
	"discordbot/src/client"
	"discordbot/src/packets"
	"discordbot/src/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := client.NewClient()

	emptyBoolChan := make(chan bool)
	emptyGuildMap := make(map[utils.Snowflake]*packets.Guild)
	emptyChannelMap := make(map[utils.Snowflake]*packets.Channel)

	assert.Equal(t, client.Token, os.Getenv("DISCORD_TOKEN"))
	assert.IsType(t, emptyChannelMap, client.Channels)
	assert.IsType(t, emptyGuildMap, client.Guilds)
	assert.IsType(t, emptyBoolChan, client.ReceiveHeartbeat)
}
