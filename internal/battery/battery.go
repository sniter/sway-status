package battery

import (
	"fmt"
)

type BatteryStatus int

const (
	BatteryCharging BatteryStatus = iota
	BatteryFull
	BatteryNotCharging
	BatteryDischarging
	BatteryUnknownStatus
)

func ParseBatteryStatus(statusStr string) (BatteryStatus, error) {
	switch statusStr {
	case "Charging":
		return BatteryCharging, nil
	case "Full":
		return BatteryFull, nil
	case "Not charging":
		return BatteryNotCharging, nil
	case "Discharging":
		return BatteryDischarging, nil
	default:
		return BatteryUnknownStatus, fmt.Errorf("unknown battery status: %s", statusStr)
	}
}

type BatteryInfo interface {
	GetCapacity() (int, error)
	GetStatus() (BatteryStatus, error)
}

type BatteryRenderer interface {
	Render(capacity int, status BatteryStatus) (string, error)
}

type Battery struct {
	Provider    BatteryInfo
	Renderer    BatteryRenderer
	LabelFormat string
	Name        string
	Instance    string
}

func (b Battery) GetName() string     { return b.Name }
func (b Battery) GetInstance() string { return b.Instance }
func (b Battery) GetFullText(_ []byte) (string, error) {
	capacity, err := b.Provider.GetCapacity()
	if err != nil {
		return "", err
	}
	status, err := b.Provider.GetStatus()
	if err != nil {
		return "", err
	}
	value, err := b.Renderer.Render(capacity, status)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(b.LabelFormat, value), nil
}
