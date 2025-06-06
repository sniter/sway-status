package weather

import (
	"io"
	"net/http"
)

type WttrIn struct {
	Url string
}

func (w WttrIn) Get() ([]byte, error) {
	resp, err := http.Get(w.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
