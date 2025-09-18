package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/allegro/bigcache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ksamf/weather_bot/database"
	"github.com/ksamf/weather_bot/models"
	"github.com/ksamf/weather_bot/weather"
)

func WeatherCommand(chatId int64, city string, msg *tgbotapi.MessageConfig, cache *bigcache.BigCache) {
	if city == "" {
		msg.Text = "⚠️ Укажи город. Пример: /weather Москва"
	}
	val, _ := cache.Get(city)

	var data models.WeatherResponse
	if len(val) != 0 {
		if err := json.Unmarshal(val, &data); err != nil {
			log.Printf("Error unmarhaling cache data: %v", err)
		}
	} else {
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru", city, os.Getenv("WEATHER_API_KEY"))

		body := weather.Get(url, msg)
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal(err)
		}
		cache.Set(city, body)
	}
	msg.Text = fmt.Sprintf("Погода в %s:\nТемпература: %.0f°C\nОщущается как: %.0f°C\nОписание: %s", data.Name, data.Main.Temp, data.Main.FeelsLike, data.Weather[0].Description)
	go database.InsertHistory(chatId, data.Name, data.Main.Temp, data.Weather[0].Description)
}
func ForecastCommand(city string, msg *tgbotapi.MessageConfig, cache *bigcache.BigCache) {
	if city == "" {
		msg.Text = "⚠️ Укажи город. Пример: /weather Москва"
	}
	val, _ := cache.Get(city)
	var data models.Forecast

	if len(val) != 0 {
		if err := json.Unmarshal(val, &data); err != nil {
			log.Printf("Error unmarhaling: %v", err)
		}
	} else {
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric&lang=ru", city, os.Getenv("WEATHER_API_KEY"))
		body := weather.Get(url, msg)
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal(err)
		}
		cache.Set(city, body)
	}
	str := fmt.Sprintf("Погода в %s:\n", data.City.Name)

	for _, d := range data.List {
		str = str + fmt.Sprintf("Дата: %s\nТемпература: %.0f°C\nОщущается как: %.0f°C\nОписание: %s\n\n", d.Dt, d.Temperature.Temp, d.Temperature.FeelsLike, d.Weather[0].Description)

	}
	msg.Text = str
}
func HistoryCommand(chatId int64, l string, msg *tgbotapi.MessageConfig) {
	limit, err := strconv.Atoi(l)
	if err != nil {
		limit = 5
	}
	history := database.GetHistory(chatId, limit)
	str := "История запросов:\n"
	for i, h := range history {
		str = str + fmt.Sprintf("%d. %s - %.0f°C (%s) [%s]\n", i+1, h.City, h.Temperature, h.Description, h.CreatedAt.Format(time.DateTime))
	}

	msg.Text = str
}
