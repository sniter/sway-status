package calendar

import (
	"time"
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
func (c Calendar) GetName() string     { return c.Name }
func (c Calendar) GetInstance() string { return c.Instance }
func (c Calendar) GetFullText(_ []byte) (string, error) {
	return c.render(), nil
}
