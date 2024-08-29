package models

type ProportionRequest struct {
	Prop      float64 `json:"prop" required:"true"`
	Mde       float64 `json:"mde" required:"true"`
	Alpha     float64 `json:"alpha" required:"true"`
	Power     float64 `json:"power" required:"true"`
	TwoTailed bool    `json:"two_tailed" required:"true"`
}
