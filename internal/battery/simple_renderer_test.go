package battery

import "testing"

type TestCase struct {
	Capacity int
	Status   BatteryStatus
	Expected string
}

func TestSimpleBatteryRenderer(t *testing.T) {
	renderer := SimpleBatteryRenderer{
		CapacityIcons: []string{"0", "25", "50", "75", "100"},
		StatusIcons: map[BatteryStatus]string{
			BatteryCharging:      "C",
			BatteryFull:          "F",
			BatteryNotCharging:   "N",
			BatteryDischarging:   "D",
			BatteryUnknownStatus: "?",
		},
		Format: "%s%s %d%%",
	}

	if result, err := renderer.Render(-1, BatteryCharging); err == nil {
		t.Errorf("when capacity < 0, then error expected, but got: %s", result)
	}

	if result, err := renderer.Render(101, BatteryCharging); err == nil {
		t.Errorf("when capacity > 100, then error expected, but got: %s", result)
	}

	testData := []TestCase{
		{0, BatteryCharging, "C0 0%"},
		{1, BatteryFull, "F0 1%"},
		{24, BatteryNotCharging, "N0 24%"},
		{25, BatteryDischarging, "D25 25%"},
		{99, BatteryUnknownStatus, "?75 99%"},
		{100, BatteryCharging, "C100 100%"},
	}

	for i, testCase := range testData {
		result, err := renderer.Render(testCase.Capacity, testCase.Status)
		if err != nil {
			t.Errorf("row #%d, expected result, got error: %s", i, err)
		}
		if result != testCase.Expected {
			t.Errorf("row #%d, expected: %s, got: %s", i, testCase.Expected, result)
		}
	}
}
