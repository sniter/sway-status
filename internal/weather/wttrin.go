package weather

import (
	"encoding/json"
	"fmt"

	"github.com/sniter/sway-status/internal/common"
)

type WttrIn struct {
	Fetch        common.Fetch
	Location     string
	WindDirIcons func(string) string
	WttrFormat   string
	Format       string
}

type WttrWeatherCode string

type WttrWeatherForecast struct {
	Conditions []WttrCurrentCondition `json:"current_condition"`
}

type WttrWeatherDescription struct {
	Value string `json:"value"`
}

type WttrCurrentCondition struct {
	Temperature   string                   `json:"temp_C"`
	WeatherCode   string                   `json:"weatherCode"`
	WindDirection string                   `json:"winddir16Point"`
	WindSpeed     string                   `json:"windspeedKmph"`
	Description   []WttrWeatherDescription `json:"weatherDesc"`
}

func WindDirIcon(dir string) string {
	switch dir {
	case "N":
		return "󰁝"
	case "NNE":
		return "󰁝"
	case "NE":
		return "󰁜"
	case "ENE":
		return "󰁔"
	case "E":
		return "󰁔"
	case "ESE":
		return "󰁔"
	case "SE":
		return "󰁃"
	case "SSE":
		return "󰁅"
	case "S":
		return "󰁅"
	case "SSW":
		return "󰁅"
	case "SW":
		return "󰁂"
	case "WSW":
		return "󰁍"
	case "W":
		return "󰁍"
	case "WNW":
		return "󰁍"
	case "NW":
		return "󰁛"
	case "NNW":
		return "󰁝"
	default:
		return "?"
	}
}

func (w WttrIn) decode(payload []byte) (string, error) {
	var p WttrWeatherForecast
	err := json.Unmarshal(payload, &p)
	if err != nil {
		return "N/A", err
	}
	condition := p.Conditions[0]
	description := condition.Description[0].Value
	temp := fmt.Sprintf("%sC", condition.Temperature)
	windDir := w.WindDirIcons(condition.WindDirection)
	windSpeed := condition.WindSpeed
	wind := fmt.Sprintf("%s%skm/h", windDir, windSpeed)

	return fmt.Sprintf(w.Format, description, temp, wind), err
}

func (w WttrIn) Get() ([]byte, error) {
	resp, err := w.Fetch(fmt.Sprintf("https://wttr.in/%s?format=%s", w.Location, w.WttrFormat))
	if err != nil {
		return nil, err
	}

	label, err := w.decode(resp)
	return []byte(label), err
}
