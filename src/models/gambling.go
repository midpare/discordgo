package models

import . "discordbot/src/utils"

type Gambling struct {
	Id             Snowflake `bson:"id"`
	GuildId        Snowflake `bson:"guildId"`
	Money          int       `bson:"money"`
	Debt           int       `bson:"debt"`
	DailyDate      int       `bson:"dailyDate"`
	BankruptcyTime int       `bson:"bankruptcyTime"`
	BaseMoneyTime  int       `bson:"baseMoneyTime"`
	Coin           []*Coin   `bson:"coin"`
}

type Coin struct {
	Name  string  `bson:"name"`
	Count int     `bson:"count"`
	Money float64 `bson:"money"`
}
