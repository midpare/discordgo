package global

import (
	"discordbot/src/models"
)

type globalStruct struct {
	Database *models.Database
}

var Global globalStruct

func init() {
	database := models.Handle()
	Global = globalStruct{
		Database: database,
	}
}
