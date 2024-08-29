package utils

import (
	"github.com/jaysrivatsa/sample_size_calculator/models"
	"github.com/montanaflynn/stats"
	"math"
)

func SampleSizeAdditive(r models.AdditiveRequest) float64 {
	var b float64
	if r.TwoTailed {
		b = stats.NormPpf(1-r.Alpha/2, 0, 1) + stats.NormPpf(r.Power, 0, 1)
	} else {
		b = stats.NormPpf(1-r.Alpha, 0, 1) + stats.NormPpf(r.Power, 0, 1)
	}
	c := r.Avg*(1+r.Mde) - r.Avg
	sampleSize := 2 * math.Pow(r.Std*b/c, 2)
	return math.Ceil(sampleSize)
}
