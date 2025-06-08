package sway

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

func notAvailableComponent(c ComponentBuilder, _ error) BarComponent {
	return BarComponent{
		Name:     c.GetName(),
		Instance: c.GetInstance(),
		FullText: "N/A",
	}
}
