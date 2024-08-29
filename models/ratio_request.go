package models

type RatioRequest struct {
	Nummean   float64 `json:"nummean" required:"true"`
	Denmean   float64 `json:"denmean" required:"true"`
	Numstd    float64 `json:"numstd" required:"true"`
	Denstd    float64 `json:"denstd" required:"true"`
	Covv      float64 `json:"covv" required:"true"`
	Mde       float64 `json:"mde" required:"true"`
	Alpha     float64 `json:"alpha" required:"true"`
	Power     float64 `json:"power" required:"true"`
	TwoTailed bool    `json:"two_tailed" required:"true"`
}
