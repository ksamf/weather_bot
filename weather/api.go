package weather

import (
	"io"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ksamf/weather_bot/models"
)

type AllowedStruct interface {
	models.Forecast | models.WeatherResponse
}

func Get(url string, msg *tgbotapi.MessageConfig) []byte {

	res, err := http.Get(url)
	if err != nil {
		msg.Text = "⚠️ ошибка при получении данных"
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		msg.Text = "❌ Город не найден"
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
