package weather

import (
	"fmt"
	"os"
	"time"

	"github.com/sniter/sway-status/internal/sway"
)

type ForecastProvider interface {
	Get() ([]byte, error)
}

type CachedForecast struct {
	Provider ForecastProvider
	File     string
	Ttl      time.Duration
}

func (c CachedForecast) Get() ([]byte, error) {
	info, err := os.Stat(c.File)
	if err == nil && time.Since(info.ModTime()) < c.Ttl {
		return os.ReadFile(c.File)
	}

	body, err := c.Provider.Get()
	if err != nil {
		return nil, err
	}
	_ = os.WriteFile(c.File, body, 0644)
	return body, nil
}

func Cached(f ForecastProvider, file string, ttl time.Duration) ForecastProvider {
	return CachedForecast{f, file, ttl}
}

type Weather struct {
	Provider    ForecastProvider
	Name        string
	Instance    string
	LabelFormat string
}

func (w Weather) ToBarComponent() (*sway.BarComponent, error) {
	payload, err := w.Provider.Get()
	if err != nil {
		return nil, err
	}
	component := sway.BarComponent{
		Name:     w.Name,
		Instance: w.Instance,
		FullText: fmt.Sprintf(w.LabelFormat, string(payload)),
	}
	return &component, nil
}
