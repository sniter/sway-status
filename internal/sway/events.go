package sway

type SwayLayoutChanged struct {
	Layout string `json:"xkb_active_layout_name"`
}

type SwayChangeEvent[T any] struct {
	Change string `json:"change"`
	Input  T      `json:"input"`
}

//{ "change": "xkb_layout", "input": { "identifier": "1:1:AT_Translated_Set_2_keyboard", "name": "AT Translated Set 2 keyboard",
//"type": "keyboard", "repeat_delay": 600, "repeat_rate": 25, "xkb_layout_names": [ "English (US)", "Latvian", "Russian" ], "xkb_active_layout_index": 1, "xkb_active_layout_name": "Latvian", "libinput": { "send_events": "enabled" }, "vendor": 1, "product": 1 } }
