package calendar

import (
	"time"

	"github.com/sniter/sway-status/internal/sway"
)

type Calendar struct {
	Format      string
	LabelFormat string
	Name        string
	Instance    string
}

func (c Calendar) render() string {
	return time.Now().Format(c.Format)
}

func (c Calendar) ToBarComponent() sway.BarComponent {
	return sway.BarComponent{
		Name:     c.Name,
		Instance: c.Instance,
		FullText: c.render(),
	}
}
