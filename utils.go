package main

import "math/rand"

func randomFloat(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

//clamps value x to mi or max
func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	} else {
		return x
	}
}
