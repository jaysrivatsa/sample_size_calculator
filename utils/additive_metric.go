package utils

import (
	"github.com/jaysrivatsa/sample_size_calculator/models"
	"math"
)

func SampleSizeAdditive(r models.AdditiveRequest) float64 {
	var b float64
	b = GetZ(r.TwoTailed, r.Alpha, r.Power)
	c := r.Avg*(1+r.Mde) - r.Avg
	sampleSize := 2 * math.Pow(r.Std*b/c, 2)
	return math.Ceil(sampleSize)
}
