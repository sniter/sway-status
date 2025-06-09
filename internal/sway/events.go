package sway

type SwayLayoutChanged struct {
	Layout string `json:"xkb_active_layout_name"`
}

type SwayChangeEvent[T any] struct {
	Change string `json:"change"`
	Input  T      `json:"input"`
}
