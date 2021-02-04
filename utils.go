package main

import (
	"math/rand"
	"time"
)

type RNG interface {
	Float64() float64
}

func NewRNG() RNG {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randFloatWBound(min, max float64, rng RNG) float64 {
	return min + rng.Float64()*(max-min)
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
