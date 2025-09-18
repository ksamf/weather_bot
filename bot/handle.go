package bot

import (
	"github.com/allegro/bigcache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, cacheWeather, cacheForecast *bigcache.BigCache) {
	if update.Message == nil || !update.Message.IsCommand() {
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "start":
		msg.Text = "Привет! Я бот прогноза погоды 🌤\nИспользуй команду /w <город>, чтобы узнать погоду.\nИсользуй /h <n> что бы посмотреть последние n запросы\nИспользуй команду /f <город>, чтобы узнать прогноз погоды."
	case "w":
		city := update.Message.CommandArguments()
		WeatherCommand(update.Message.Chat.ID, city, &msg, cacheWeather)

	case "h":
		limit := update.Message.CommandArguments()
		HistoryCommand(update.Message.Chat.ID, limit, &msg)

	case "f":
		city := update.Message.CommandArguments()
		ForecastCommand(city, &msg, cacheForecast)
	}
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

}
