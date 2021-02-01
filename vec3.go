package main

import "math"

type Vec3 struct {
	x, y, z float64
}

func (p Vec3) Sub(other Vec3) Vec3 {
	return Vec3{p.x - other.x, p.y - other.y, p.z - other.z}
}

func (p *Vec3) SubInPlace(other Vec3) {
	*p = p.Sub(other)
}

func (p Vec3) Add(other Vec3) Vec3 {
	return Vec3{p.x + other.x, p.y + other.y, p.z + other.z}
}

func (p *Vec3) AddInPlace(other Vec3) {
	*p = p.Add(other)
}

func (p Vec3) Mul(x float64) Vec3 {
	return Vec3{p.x * x, p.y * x, p.z * x}
}

func (p *Vec3) MulInPlace(x float64) {
	*p = p.Mul(x)
}

func (p Vec3) Divide(x float64) Vec3 {
	return p.Mul(1 / x)
}

func (p *Vec3) DivInPlace(x float64) {
	*p = p.Divide(x)
}

func (p Vec3) LenSquared() float64 {
	return p.x*p.x + p.y*p.y + p.z*p.z
}

func (p Vec3) Len() float64 {
	return math.Sqrt(p.LenSquared())
}

func (p Vec3) UnitVector() Vec3 {
	return p.Divide(p.Len())
}

func Dot(vec1, vec2 Vec3) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y + vec1.z*vec2.z
}
