package battery

import (
	"fmt"

	"github.com/sniter/sway-status/internal/sway"
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

func (b Battery) failedComponent() sway.BarComponent {
	return sway.BarComponent{
		Name:     b.Name,
		Instance: b.Instance,
		FullText: "N/A",
	}
}

func (b Battery) ToBarComponent() (sway.BarComponent, error) {
	capacity, err := b.Provider.GetCapacity()
	if err != nil {
		return b.failedComponent(), err
	}
	status, err := b.Provider.GetStatus()
	if err != nil {
		return b.failedComponent(), err
	}
	value, err := b.Renderer.Render(capacity, status)
	if err != nil {
		return b.failedComponent(), err
	}
	component := sway.BarComponent{
		Name:     b.Name,
		Instance: b.Instance,
		FullText: fmt.Sprintf(b.LabelFormat, value),
	}
	return component, nil
}

// func getBattery() string {
// 	statusPath := "/sys/class/power_supply/BAT0/status"
//
// 	statusData, err := os.ReadFile(statusPath)
// 	if err != nil {
// 		return "N/A"
// 	}
//
// 	icons := []string{"󰁺", "󰁻", "󰁼", "󰁽", "󰁾", "󰁿", "󰂀", "󰂁", "󰂂", "󰁹"}
// 	iconIdx := capacityVal / 10
// 	if iconIdx >= len(icons) {
// 		iconIdx = len(icons) - 1
// 	}
//
// 	statusIcon := ""
// 	switch statusStr {
// 	case "Charging":
// 		statusIcon = ""
// 	case "Full":
// 		statusIcon = ""

// 	case "Not charging":
// 		statusIcon = ""
// 	case "Discharging":
// 		statusIcon = ""
// 	default:
// 		statusIcon = statusStr
// 	}
//
// 	return fmt.Sprintf("%s%s %d%%", statusIcon, icons[iconIdx], capacityVal)
// }
