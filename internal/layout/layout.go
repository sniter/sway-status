package layout

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/sniter/sway-status/internal/common/cache"
	"github.com/sniter/sway-status/internal/sway"
)

type Layout struct {
	Cache       cache.Cache[string, []byte]
	Renderer    func(string) string
	Name        string
	Instance    string
	LabelFormat string
}

func BasicRenderer(layout string) string {
	switch {
	case strings.HasPrefix(layout, "Russian"):
		return "RU"
	case strings.HasPrefix(layout, "Latvian"):
		return "LV"
	case strings.HasPrefix(layout, "English"):
		return "EN"
	default:
		if len(layout) >= 2 {
			return strings.ToUpper(layout[:2])
		}
		return layout
	}
}

const cacheKey string = "layout"

func (l Layout) fullTextFromCache() (string, error) {
	value, ok := l.Cache.Get(cacheKey)
	if ok {
		return string(value), nil
	}
	return "", errors.New("cache miss")
}

func (l Layout) GetName() string     { return l.Name }
func (l Layout) GetInstance() string { return l.Instance }
func (l Layout) GetFullText(input []byte) (string, error) {
	if len(input) == 0 {
		return l.fullTextFromCache()
	}

	var event sway.SwayChangeEvent[sway.SwayLayoutChanged]
	err := json.Unmarshal(input, &event)
	if err != nil {
		return "", err
	}

	// Extra validation
	if event.Change != "xkb_layout" || event.Input.Layout == "" {
		return l.fullTextFromCache()
	}

	locale := fmt.Sprintf(l.LabelFormat, l.Renderer(event.Input.Layout))
	l.Cache.Put(cacheKey, []byte(locale))
	return locale, nil
}
