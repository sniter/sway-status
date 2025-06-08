package sway

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalSwayChangeEvent(t *testing.T) {
	input := `{ "change": "xkb_layout", "input": { "identifier": "1:1:AT_Translated_Set_2_keyboard", "name": "AT Translated Set 2 keyboard", "type": "keyboard", "repeat_delay": 600, "repeat_rate": 25, "xkb_layout_names": [ "English (US)", "Latvian", "Russian" ], "xkb_active_layout_index": 0, "xkb_active_layout_name": "English (US)", "libinput": { "send_events": "enabled" }, "vendor": 1, "product": 1 } }`
	var actual SwayChangeEvent[SwayLayoutChanged]
	err := json.Unmarshal([]byte(input), &actual)
	if err != nil {
		t.Errorf("Decoding failed with error: %s", err)
	}
	expected := SwayChangeEvent[SwayLayoutChanged]{
		Change: "xkb_layout",
		Input: SwayLayoutChanged{
			Layout: "English (US)",
		},
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Mismatch (-expected +actual):\n%s", diff)
	}
}
