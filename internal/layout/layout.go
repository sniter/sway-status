package layout

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/sniter/sway-status/internal/common/cache"
	"github.com/sniter/sway-status/internal/common/source"
	"github.com/sniter/sway-status/internal/sway"
)

type Layout struct {
	InitialValue source.Source
	Cache        cache.Cache[string, []byte]
	Renderer     func(string) string
	Name         string
	Instance     string
	LabelFormat  string
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

func (l Layout) readKeyboardEvent(input []byte) (string, bool) {
	var event sway.SwayChangeEvent[sway.SwayLayoutChanged]
	err := json.Unmarshal(input, &event)
	if err != nil {
		return "", false
	}
	if event.Change != "xkb_layout" || event.Input.Layout == "" {
		return "", false
	}
	return fmt.Sprintf(l.LabelFormat, l.Renderer(event.Input.Layout)), true
}

func (l Layout) readInitialEvent(input []byte) (string, bool) {
	var event sway.TickEvent
	if err := json.Unmarshal(input, &event); err != nil || !event.First {
		// log.Printf("Failed decode initial event: %s\n%s", err, string(input))
		return "", false
	}

	if value, err := l.InitialValue.ReadString(); err == nil {
		return fmt.Sprintf(l.LabelFormat, l.Renderer(value)), true
	}
	// log.Printf("failed read init value: %s", err)
	return "", false
}

func (l Layout) GetName() string     { return l.Name }
func (l Layout) GetInstance() string { return l.Instance }
func (l Layout) GetFullText(input []byte) (string, error) {
	if len(input) == 0 {
		return l.fullTextFromCache()
	}

	if locale, ok := l.readInitialEvent(input); ok {
		l.Cache.Put(cacheKey, []byte(locale))
		return locale, nil
	}

	if locale, ok := l.readKeyboardEvent(input); ok {
		l.Cache.Put(cacheKey, []byte(locale))
		return locale, nil
	}

	return l.fullTextFromCache()
}
