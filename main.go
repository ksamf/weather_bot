package main

import (
	"github.com/ksamf/weather_bot/bot"
	"github.com/ksamf/weather_bot/database"
)

func main() {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %s", err)
	// }
	database.InitDb()
	bot.Start()
	defer database.DB.Close()

}
