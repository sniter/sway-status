package battery

func PickupCapacityIcon(icons []string) func(int) string {
	return func(capacity int) string {
		if capacity <= 0 {
			return icons[0]
		}
		if capacity >= 100 {
			return icons[len(icons)-1]
		}
		capacityStep := 100 / float64(len(icons)-1)
		capacityIndex := int(float64(capacity) / capacityStep)
		return icons[capacityIndex]
	}
}

func PickupStatusIcon(icons map[BatteryStatus]string, ifNotFound string) func(BatteryStatus) string {
	return func(status BatteryStatus) string {
		value, ok := icons[status]
		if ok {
			return value
		}
		return ifNotFound
	}
}
