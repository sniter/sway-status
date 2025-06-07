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

func (p PowerSupplyProvider) GetCapacity() (int, error) {
	capacityPath := fmt.Sprintf("/sys/class/power_supply/%s/capacity", p.Device)
	capacityData, err := os.ReadFile(capacityPath)
	if err != nil {
		return -1, err
	}
	capacityStr := strings.TrimSpace(string(capacityData))
	capacityVal, err := strconv.Atoi(capacityStr)
	if err != nil {
		return -2, err
	}
	return capacityVal, nil
}

func (p PowerSupplyProvider) GetStatus() (BatteryStatus, error) {
	statusPath := fmt.Sprintf("/sys/class/power_supply/%s/status", p.Device)

	statusData, err := os.ReadFile(statusPath)
	if err != nil {
		return BatteryUnknownStatus, err
	}
	statusStr := strings.TrimSpace(string(statusData))
	status, err := ParseBatteryStatus(statusStr)
	if err != nil {
		return BatteryUnknownStatus, err
	}
	return status, nil
}
