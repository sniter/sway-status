package layout

import (
	"fmt"
	"strings"

	"github.com/sniter/sway-status/internal/sway"
)

type Layout struct {
	Renderer    func(string) string
	Name        string
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

func (l Layout) ToBarComponent(input string) sway.BarComponent {
	// NOTE: decode event
	return sway.BarComponent{
		Name:     l.Name,
		Instance: "main",
		FullText: fmt.Sprintf(l.LabelFormat, l.Renderer(input)),
	}
}
