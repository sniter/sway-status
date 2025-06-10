package battery

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PowerSupplyProvider struct {
	Device string
}

func (p PowerSupplyProvider) Report() BatteryReport {
	return BatteryReport{
		Capacity: p.Capacity(),
		Status:   p.Status(),
	}
}

func (p PowerSupplyProvider) Capacity() int {
	capacityPath := fmt.Sprintf("/sys/class/power_supply/%s/capacity", p.Device)
	capacityData, err := os.ReadFile(capacityPath)
	if err != nil {
		return -1
	}
	capacityStr := strings.TrimSpace(string(capacityData))
	capacityVal, err := strconv.Atoi(capacityStr)
	if err != nil {
		return -2
	}
	return capacityVal
}

func (p PowerSupplyProvider) Status() BatteryStatus {
	statusPath := fmt.Sprintf("/sys/class/power_supply/%s/status", p.Device)

	statusData, err := os.ReadFile(statusPath)
	if err != nil {
		return BatteryUnknownStatus
	}
	statusStr := strings.TrimSpace(string(statusData))
	status, err := ParseBatteryStatus(statusStr)
	if err != nil {
		return BatteryUnknownStatus
	}
	return status
}
