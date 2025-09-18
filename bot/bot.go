package bot

import (
	"log"
	"os"
	"time"

	"github.com/allegro/bigcache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	cacheWeatherConfig := bigcache.DefaultConfig(60 * time.Minute)
	cacheWeather, err := bigcache.NewBigCache(cacheWeatherConfig)
	if err != nil {
		log.Fatal(err)
	}
	cacheForecastConfig := bigcache.DefaultConfig(300 * time.Minute)
	cacheForecast, err := bigcache.NewBigCache(cacheForecastConfig)
	if err != nil {
		log.Fatal(err)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		go handleUpdate(update, bot, cacheWeather, cacheForecast)

	}
}
