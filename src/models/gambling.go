package models

import (
	"context"
	. "discordbot/src/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Gambling struct {
	Id             Snowflake `bson:"id"`
	GuildId        Snowflake `bson:"guildId"`
	Money          int       `bson:"money"`
	Debt           int       `bson:"debt"`
	DailyDate      string    `bson:"dailyDate"`
	BankruptcyTime int       `bson:"bankruptcyTime"`
	BaseMoneyTime  int       `bson:"baseMoneyTime"`
	Coin           []*Coin   `bson:"coin"`
}

type Coin struct {
	Name  string  `bson:"name"`
	Count int     `bson:"count"`
	Money float64 `bson:"money"`
}

func getGamb(database *mongo.Database) map[Snowflake]map[Snowflake]*Gambling {
	gambling := database.Collection("gamblings")
	cur, _ := gambling.Find(context.TODO(), bson.M{})

	data := make(map[Snowflake]map[Snowflake]*Gambling)
	for cur.Next(context.TODO()) {
		var user *Gambling
		if err := cur.Decode(&user); err != nil {
			log.Fatalf("error decoding cursur gambling user \n%s\n", err)
		}

		_, e := data[user.GuildId]
		if !e {
			data[user.GuildId] = make(map[Snowflake]*Gambling)
		}

		data[user.GuildId][user.Id] = user
	}

	return data
}
