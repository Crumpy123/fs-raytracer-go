package main

import "math"

type Vec3 struct {
	x, y, z float64
}

func (p *Vec3) PSubtract(other Vec3) Vec3 {
	return Vec3{p.x - other.x, p.y - other.y, p.z - other.z}
}

func (p Vec3) Subtract(other Vec3) Vec3 {
	return Vec3{p.x - other.x, p.y - other.y, p.z - other.z}
}

func (p *Vec3) PAdd(other Vec3) Vec3 {
	return Vec3{p.x + other.x, p.y + other.y, p.z + other.z}
}

func (p Vec3) Add(other Vec3) Vec3 {
	return Vec3{p.x + other.x, p.y + other.y, p.z + other.z}
}

func (p *Vec3) PMultiply(x float64) Vec3 {
	return Vec3{p.x * x, p.y * x, p.z * x}
}

func (p Vec3) Multiply(x float64) Vec3 {
	return Vec3{p.x * x, p.y * x, p.z * x}
}

func (p *Vec3) PDivide(x float64) Vec3 {
	return p.PMultiply(1 / x)
}

func (p Vec3) Divide(x float64) Vec3 {
	return p.PMultiply(1 / x)
}

func (p *Vec3) PLengthSquared() float64 {
	return p.x*p.x + p.y*p.y + p.z*p.z
}

func (p Vec3) LengthSquared() float64 {
	return p.x*p.x + p.y*p.y + p.z*p.z
}

func (p *Vec3) PLength() float64 {
	return math.Sqrt(p.PLengthSquared())
}

func (p *Vec3) Length() float64 {
	return math.Sqrt(p.PLengthSquared())
}

func (p *Vec3) PUnitVector() Vec3 {
	return p.PDivide(p.PLength())
}

func (p Vec3) UnitVector() Vec3 {
	return p.PDivide(p.PLength())
}

func Dot(vec1, vec2 Vec3) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y + vec1.z*vec2.z
}
