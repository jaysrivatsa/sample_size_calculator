package utils

import "github.com/montanaflynn/stats"

func GetZ(twoTailed bool, alpha float64, power float64) float64 {
	var b float64
	if twoTailed {
		b = stats.NormPpf(1-alpha/2, 0, 1) + stats.NormPpf(power, 0, 1)
	} else {
		b = stats.NormPpf(1-alpha, 0, 1) + stats.NormPpf(power, 0, 1)
	}
	return b
}
