package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	x, y, z float64
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{v.x - other.x, v.y - other.y, v.z - other.z}
}

func (v *Vec3) SubInPlace(other Vec3) {
	*v = v.Sub(other)
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.x + other.x, v.y + other.y, v.z + other.z}
}

func (v *Vec3) AddInPlace(other Vec3) {
	*v = v.Add(other)
}

func (v Vec3) Mul(x float64) Vec3 {
	return Vec3{v.x * x, v.y * x, v.z * x}
}

func (v Vec3) MulWithVec3(other Vec3) Vec3 {
	return Vec3{v.x * other.x, v.y * other.y, v.z * other.z}
}

func (v *Vec3) MulInPlace(x float64) {
	*v = v.Mul(x)
}

func (v Vec3) Divide(x float64) Vec3 {
	return v.Mul(1 / x)
}

func (v *Vec3) DivInPlace(x float64) {
	*v = v.Divide(x)
}

func (v Vec3) LenSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

func (v Vec3) UnitVector() Vec3 {
	return v.Divide(v.Len())
}

func (v *Vec3) UnitVectorInPlace() {
	*v = v.UnitVector()
}

func Dot(vec1, vec2 Vec3) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y + vec1.z*vec2.z
}

func randVec3() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func randVec3WBound(min, max float64, rng RNG) Vec3 {
	return Vec3{randFloatWBound(min, max, rng), randFloatWBound(min, max, rng), randFloatWBound(min, max, rng)}
}

func randInUnitSphere(rng RNG) Vec3 {
	for {
		p := randVec3WBound(-1, 1, rng)
		if p.LenSquared() < 1 {
			return p
		}
	}
}

func randUnitVector(rng RNG) Vec3 {
	return randInUnitSphere(rng).UnitVector()
}

//return true if vector is near zero
func (v Vec3) nearZero() bool {
	s := 1e-8
	return (math.Abs(v.x) < s) && (math.Abs(v.y) < s) && (math.Abs(v.z) < s)
}

func reflectVec(a Vec3, b Vec3) Vec3 {
	return a.Sub(b.Mul(Dot(a, b) * 2))
}

func refractVec(uv Vec3, n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(Dot(uv.Mul(-1), n), 1.0)
	//perpendicular
	rOutPerp := uv.Add(n.Mul(cosTheta)).Mul(etaiOverEtat)
	rOutParallel := n.Mul(-1 * math.Sqrt(math.Abs(1.0-rOutPerp.LenSquared())))
	return rOutPerp.Add(rOutParallel)

}
