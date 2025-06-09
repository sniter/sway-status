package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"time"

	"github.com/sniter/sway-status/internal/battery"
	"github.com/sniter/sway-status/internal/calendar"
	"github.com/sniter/sway-status/internal/common"
	"github.com/sniter/sway-status/internal/common/cache"
	"github.com/sniter/sway-status/internal/layout"
	"github.com/sniter/sway-status/internal/sway"
	"github.com/sniter/sway-status/internal/sysmon"
	"github.com/sniter/sway-status/internal/weather"
)

func makeHandler() sway.SimpleSwayDelegate {
	return sway.SimpleSwayDelegate{
		Components: []sway.ComponentBuilder{
			sysmon.SysMon{
				DiskName:     "/nvme0",
				CpuStatsFile: "/tmp/cpu_stat",
				LabelFormat:  "temp: %s cpu: %d%% dsk: %s mem: %d%%/%d%% ",
				Name:         "sysmon",
				Instance:     "main",
			},

			weather.Weather{
				Provider: weather.WttrIn{
					Fetch:        common.Fetch(common.FetchFrom).ReadThrough(fnv.New64(), 20*time.Minute),
					Location:     "Riga",
					WttrFormat:   "j1",
					WindDirIcons: weather.WindDirIcon,
					Format:       "%s %s %s",
				},
				LabelFormat: " %s ",
				Name:        "weather",
				Instance:    "main",
			},

			battery.Battery{
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
			},

			layout.Layout{
				Cache:       cache.MakeTempFileCache("layout_"),
				Renderer:    layout.BasicRenderer,
				Name:        "layout",
				LabelFormat: " %s ",
			},

			calendar.Calendar{
				Name:     "calendar",
				Instance: "Local",
				Format:   "Mon Jan 2 15:04",
			},
		},
	}
}

func main() {
	handler := makeHandler()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`{"version": 1, "click_events": false}`)
	fmt.Println("[")
	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Println(handler.Handle(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(handler.OnError(err))
	}
}
