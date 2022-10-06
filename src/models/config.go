package models

import . "discordbot/src/utils"

type Config struct {
	Id       Snowflake `bson:"id"`
	GuildId  Snowflake `bson:"guildid"`
	Name     string    `bson:"name"`
	Slangs   []string  `bson:"slangs"`
	Warning  int       `bson:"warning"`
	BanTime  int       `bson:"banTime"`
	MuteTime int       `bson:"muteTime"`
}
