package main

import "math"

type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) Hit(ray *Ray, minT float64, maxT float64, rec *HitRecord) bool {
	oc := ray.Origin.Sub(s.center)
	a := ray.Direction.LenSquared()
	halfB := Dot(oc, ray.Direction)
	c := oc.LenSquared() - (s.radius * s.radius)

	discriminant := (halfB * halfB) - (a * c)
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	root := (-halfB - sqrtd) / a
	if root < minT || maxT < root {
		root = (-halfB + sqrtd) / a
		if root < minT || maxT < root {
			return false
		}
	}

	rec.t = root
	rec.p = ray.At(rec.t)
	outwardNormal := (rec.p.Sub(s.center)).Divide(s.radius)
	rec.setFaceNormal(ray, &outwardNormal)
	//rec.normal = (rec.p.Sub(s.center)).Divide(s.radius)
	return true
}