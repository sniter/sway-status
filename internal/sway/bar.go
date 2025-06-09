package sway

import (
	"encoding/json"
	"fmt"
)

type BarComponent struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
	FullText string `json:"full_text"`
}

type ComponentBuilder interface {
	GetName() string
	GetInstance() string
	GetFullText(input []byte) (string, error)
}

func BuildComponent(c ComponentBuilder, input []byte) BarComponent {
	fullText, err := c.GetFullText(input)
	if err != nil {
		return notAvailableComponent(c, err)
	}
	return BarComponent{
		Name:     c.GetName(),
		Instance: c.GetInstance(),
		FullText: fullText,
	}
}

func notAvailableComponent(c ComponentBuilder, err error) BarComponent {
	return BarComponent{
		Name:     c.GetName(),
		Instance: c.GetInstance(),
		FullText: err.Error(),
	}
}

func errorComponent(err error) BarComponent {
	return BarComponent{
		Name:     "error",
		Instance: "sway-status",
		FullText: err.Error(),
	}
}

type SwayDelegate interface {
	Handle(event []byte) string
	OnError(error) string
}
type SimpleSwayDelegate struct {
	Components []ComponentBuilder
}

func (m SimpleSwayDelegate) Handle(event []byte) string {
	statusBar := make([]BarComponent, len(m.Components))
	for idx, component := range m.Components {
		statusBar[idx] = BuildComponent(component, event)
	}
	jsonBar, err := json.Marshal(statusBar)
	if err != nil {
		m.OnError(err)
	}
	return fmt.Sprintf("%s,", string(jsonBar))
}

func (m SimpleSwayDelegate) OnError(err error) string {
	statusBar := []BarComponent{errorComponent(err)}
	jsonBar, err := json.Marshal(statusBar)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%s,", string(jsonBar))
}
