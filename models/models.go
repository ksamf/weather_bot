package models

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
	} `json:"main"`
	Name string `json:"name"`
}

type Forecast struct {
	List []struct {
		Temperature struct {
			Temp      float32 `json:"temp"`
			FeelsLike float32 `json:"feels_like"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Dt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Name string `json:"name"`
	} `json:"city"`
}
