package main

import (
	"github.com/ksamf/weather_bot/bot"
	"github.com/ksamf/weather_bot/database"
)

func main() {

	database.InitDb()
	bot.Start()
	defer database.DB.Close()

}
