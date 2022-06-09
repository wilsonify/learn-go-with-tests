package models

type StrengthInput struct {
	Actual int32 `json:"actual,omitempty"`

	Expected int32 `json:"expected,omitempty"`
}
