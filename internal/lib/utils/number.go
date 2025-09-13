package utils

import "math"

func RoundFloat32(val float32) float32 {
	return float32(math.Round(float64(val)*100) / 100)
}
