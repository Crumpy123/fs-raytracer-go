package main

import "math"

type Material interface {
	Reflect(lightRay *LightRay, rng RNG)
}

type Lambertian struct {
	color Vec3
}

func (l Lambertian) Reflect(lightRay *LightRay, rng RNG) {
	scatterDirection := lightRay.hitRecord.normal.Add(randUnitVector(rng))

	//catch degenerate scatter direction
	if scatterDirection.nearZero() {
		scatterDirection = lightRay.hitRecord.normal
	}

	lightRay.ray = Ray{lightRay.hitRecord.point, scatterDirection}
	lightRay.color = lightRay.color.MulWithVec3(l.color)
}

type Metal struct {
	color Vec3
	fuzz  float64
}

func (m Metal) Reflect(lightRay *LightRay, rng RNG) {
	reflected := reflectVec(lightRay.ray.Direction.UnitVector(), lightRay.hitRecord.normal)
	lightRay.ray = Ray{lightRay.hitRecord.point, reflected.Add(randInUnitSphere(rng).Mul(m.fuzz))}
	lightRay.color = lightRay.color.MulWithVec3(m.color)
}

type Dielectric struct {
	indexOfRefraction float64
}

func (d Dielectric) Reflect(lightRay *LightRay, rng RNG) {
	lightRay.color = lightRay.color.MulWithVec3(Vec3{1, 1, 1})
	var refractionRatio float64
	if lightRay.hitRecord.frontFace {
		refractionRatio = 1.0 / d.indexOfRefraction
	} else {
		refractionRatio = d.indexOfRefraction
	}
	unitDirection := lightRay.ray.Direction.UnitVector()
	cosTheta := math.Min(Dot(unitDirection.Mul(-1), lightRay.hitRecord.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0
	var direction Vec3
	if cannotRefract {
		direction = reflectVec(unitDirection, lightRay.hitRecord.normal)
	} else {
		direction = refractVec(unitDirection, lightRay.hitRecord.normal, refractionRatio)
	}
	//refracted := refractVec(lightRay.ray.Direction.UnitVector(), lightRay.hitRecord.normal, refractionRatio)

	lightRay.ray = Ray{lightRay.hitRecord.point, direction}
}
