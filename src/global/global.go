package global

import (
	"context"
	"discordbot/src/models"
	. "discordbot/src/utils"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Gambling map[Snowflake]map[Snowflake]models.Gambling
}

type globalStruct struct {
	DbClient *mongo.Client
	Mongodb  *mongo.Database
	Database *Database
}

var Global globalStruct

func init() {
	dbOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB_URI"))
	dbClient, err := mongo.Connect(context.TODO(), dbOptions)

	if err != nil {
		log.Fatalf("error connecting to mongodb\n%s\n", err)
	}

	mongodb := dbClient.Database("discordbot")
	gambling := mongodb.Collection("gamblings")

	database := Database{
		Gambling: make(map[Snowflake]map[Snowflake]models.Gambling),
	}

	cur, _ := gambling.Find(context.TODO(), bson.M{})

	for cur.Next(context.TODO()) {
		var user *models.Gambling
		if err := cur.Decode(&user); err != nil {
			log.Fatalf("error decoding cursur gambling user \n%s\n", err)
		}

		_, e := database.Gambling[user.GuildId]
		if !e {
			database.Gambling[user.GuildId] = make(map[Snowflake]models.Gambling)
		}

		database.Gambling[user.GuildId][user.Id] = *user
	}

	Global = globalStruct{
		DbClient: dbClient,
		Mongodb:  mongodb,
		Database: &database,
	}
}
