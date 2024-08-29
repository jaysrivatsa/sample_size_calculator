package models

type AdditiveRequest struct {
	Avg       float64 `json:"avg" required:"true"`
	Std       float64 `json:"std" required:"true"`
	Mde       float64 `json:"mde" required:"true"`
	Alpha     float64 `json:"alpha" required:"true"`
	Power     float64 `json:"power" required:"true"`
	TwoTailed bool    `json:"two_tailed" required:"true"`
}
