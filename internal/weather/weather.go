package weather

import (
	"fmt"
)

type ForecastProvider interface {
	Get() ([]byte, error)
}

type Weather struct {
	Provider    ForecastProvider
	Name        string
	Instance    string
	LabelFormat string
}

func (w Weather) GetName() string     { return w.Name }
func (w Weather) GetInstance() string { return w.Instance }
func (w Weather) GetFullText(_ []byte) (string, error) {
	payload, err := w.Provider.Get()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(w.LabelFormat, payload), nil
}
