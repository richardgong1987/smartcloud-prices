package models

type Price struct {
	Kind   string  `json:"kind"`
	Amount float64 `json:"amount"`
}

type SmartcloudResponse struct {
	Prices []Price `json:"prices"`
}
