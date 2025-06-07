package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sniter/sway-status/internal/battery"
	"github.com/sniter/sway-status/internal/calendar"
	"github.com/sniter/sway-status/internal/layout"
	"github.com/sniter/sway-status/internal/sway"
	"github.com/sniter/sway-status/internal/sysmon"
	"github.com/sniter/sway-status/internal/weather"
)

func main() {
	weatherComp, err := weather.Weather{
		Provider:    weather.Cached(weather.WttrIn{Url: "https://wttr.in/Jurmala?format=%C+%t+%w"}, "/tmp/wttr_cache", 20*time.Minute),
		LabelFormat: " %s ",
		Name:        "weather",
		Instance:    "main",
	}.ToBarComponent()
	if err != nil {
		panic(err)
	}
	sysMonComp := sysmon.SysMon{
		DiskName:     "/nvme0",
		CpuStatsFile: "/tmp/cpu_stat",
		LabelFormat:  "temp: %s cpu: %d%% dsk: %s mem: %d%%/%d%% ",
		Name:         "sysmon",
		Instance:     "main",
	}.ToBarComponent()

	battery, _ := battery.Battery{
		Provider: battery.PowerSupplyProvider{Device: "BAT0"},
		Renderer: battery.SimpleBatteryRenderer{
			CapacityIcons: []string{"󰁺", "󰁻", "󰁼", "󰁽", "󰁾", "󰁿", "󰂀", "󰂁", "󰂂", "󰁹"},
			StatusIcons: map[battery.BatteryStatus]string{
				battery.BatteryCharging:      "",
				battery.BatteryFull:          "",
				battery.BatteryNotCharging:   "",
				battery.BatteryDischarging:   "",
				battery.BatteryUnknownStatus: "",
			},
			Format: "%s%s %d%%",
		},
		LabelFormat: " %s ",
		Name:        "battery",
		Instance:    "bat0",
	}.ToBarComponent()

	layout := layout.Layout{
		Renderer:    layout.BasicRenderer,
		Name:        "layout",
		LabelFormat: " %s ",
	}.ToBarComponent("English (US)")

	calendar := calendar.Calendar{
		Name:     "calendar",
		Instance: "Riga",
		Format:   "Mon Jan 2 15:04",
	}.ToBarComponent()

	components := []sway.BarComponent{
		sysMonComp,
		*weatherComp,
		layout,
		battery,
		calendar,
	}
	result, err := json.Marshal(components)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s,", result)
}
