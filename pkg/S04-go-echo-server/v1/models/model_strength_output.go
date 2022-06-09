package models

type StrengthOutput struct {

	Actual int32 `json:"actual,omitempty"`

	Expected int32 `json:"expected,omitempty"`

	Strength float32 `json:"strength,omitempty"`
}
