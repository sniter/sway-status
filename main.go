package main

import (
	"encoding/json"
	"fmt"
	"time"

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
		LabelFormet:  "temp %s cpu %d%% dsk %s mem %d%%/%d%% ",
		Name:         "sysmon",
		Instance:     "main",
	}.ToBarComponent()

	components := []sway.BarComponent{
		sysMonComp,
		*weatherComp,
	}
	result, err := json.Marshal(components)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s,", result)
}
