package service

type FilterData struct {
	After  *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type MinMaxFilterInput struct {
	Min float64  `json:"min"`
	Max *float64 `json:"max,omitempty"`
}
