package weather

import (
	"testing"
)

func TestWttrInDecode(t *testing.T) {
	w := WttrIn{
		Location:     "Foo",
		WindDirIcons: WindDirIcon,
		WttrFormat:   "j1",
		Format:       "%s %s %s",
	}
	payload := `{"current_condition":[{"FeelsLikeC":"17","FeelsLikeF":"63","cloudcover":"25","humidity":"88","localObsDateTime":"2025-06-07 09:22 PM","observation_time":"06:22 PM","precipInches":"0.0","precipMM":"0.0","pressure":"1007","pressureInches":"30","temp_C":"17","temp_F":"63","uvIndex":"0","visibility":"10","visibilityMiles":"6","weatherCode":"113","weatherDesc":[{"value":"Sunny"}],"weatherIconUrl":[{"value":""}],"winddir16Point":"SSW","winddirDegree":"206","windspeedKmph":"6","windspeedMiles":"4"}]}`
	actual, err := w.decode([]byte(payload))
	if err != nil {
		t.Errorf("Failed decode JSON: %s", err)
	}
	expected := "Sunny 17C 󰁅6km/h"

	if actual != expected {
		t.Errorf(`"%s" != "%s"`, actual, expected)
	}
}
