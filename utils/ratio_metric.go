package utils

import (
	"github.com/jaysrivatsa/sample_size_calculator/models"
	"math"
)

func SampleSizeRatio(r models.RatioRequest) float64 {
	numVar := math.Pow(r.Numstd, 2)
	denVar := math.Pow(r.Denstd, 2)

	a := math.Pow(r.Nummean, 2) / math.Pow(r.Denmean, 2)

	b := numVar/math.Pow(r.Nummean, 2) + denVar/math.Pow(r.Denmean, 2)

	c := 2 * r.Covv / (r.Nummean * r.Denmean)

	tau := a * (b - c)

	z := GetZ(r.TwoTailed, r.Alpha, r.Power)

	newMetricValue := (r.Nummean / r.Denmean) * r.Mde
	sampleSize := math.Ceil((2 * tau * math.Pow(z, 2)) / (math.Pow(newMetricValue, 2)))
	return math.Ceil(sampleSize)
}
