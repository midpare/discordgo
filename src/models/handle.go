package models

import (
	"context"
	. "discordbot/src/utils"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Gambling map[Snowflake]map[Snowflake]*Gambling
}

func Handle() *Database {
	dbOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB_URI"))
	dbClient, err := mongo.Connect(context.TODO(), dbOptions)
	mongodb := dbClient.Database("discordbot")

	if err != nil {
		log.Fatalf("error connecting to mongodb\n%s\n", err)
	}

	return &Database{
		Gambling: getGamb(mongodb),
	}
}
