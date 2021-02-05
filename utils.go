package main

import (
	"math"
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

func randFloatWBoundExpensive(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat(rng RNG) float64 {
	return rng.Float64()
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

func degreeSToRadians(degreeS float64) float64 {
	return degreeS * (math.Pi / 180)
}

func radiansToDegree(radians float64) float64 {
	return radians * (180 / math.Pi)
}
