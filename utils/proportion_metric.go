package utils

import (
	"github.com/jaysrivatsa/sample_size_calculator/models"
	"math"
)

func SampleSizeProportion(r models.ProportionRequest) float64 {
	var b float64
	b = GetZ(r.TwoTailed, r.Alpha, r.Power)
	metricValue := r.Prop
	newMetricValue := r.Prop * (1 + r.Mde)

	a := metricValue*(1-metricValue) + newMetricValue*(1-newMetricValue)
	c := math.Pow(metricValue-newMetricValue, 2)

	sampleSize := math.Pow(b, 2) * a / c

	return math.Ceil(sampleSize)
}
