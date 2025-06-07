package battery

import (
	"errors"
	"fmt"
)

type SimpleBatteryRenderer struct {
	CapacityIcons []string
	StatusIcons   map[BatteryStatus]string
	Format        string
}

func validateCapacity(capacity int) (int, error) {
	if capacity < 0 {
		return 0, errors.New("battery capacity < 0%")
	} else if capacity > 100 {
		return 100, errors.New("battery capacity > 100%")
	} else {
		return capacity, nil
	}
}

func (s SimpleBatteryRenderer) detectCapacityIdx(capacity int) int {
	if capacity == 0 {
		return 0
	}
	capacityStep := 100 / float64(len(s.CapacityIcons)-1)
	return int(float64(capacity) / capacityStep)
}

func (s SimpleBatteryRenderer) Render(capacity int, status BatteryStatus) (string, error) {
	newCapacity, err := validateCapacity(capacity)
	capacityIdx := s.detectCapacityIdx(newCapacity)

	return fmt.Sprintf(s.Format, s.StatusIcons[status], s.CapacityIcons[capacityIdx], newCapacity), err
}
