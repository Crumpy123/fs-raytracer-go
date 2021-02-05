package main

type Material interface {
	Reflect(lightRay *LightRay, rng RNG)
}

type Lambertian struct {
	albedo Vec3
}

type Metal struct {
	albedo Vec3
}

func (l Lambertian) Reflect(lightRay *LightRay, rng RNG) {
	scatterDirection := lightRay.hitRecord.normal.Add(randUnitVector(rng))

	//catch degenerate scatter direction
	if scatterDirection.nearZero() {
		scatterDirection = lightRay.hitRecord.normal
	}

	lightRay.ray = Ray{lightRay.hitRecord.point, scatterDirection}
	lightRay.color = lightRay.color.MulWithVec3(l.albedo)
}

func (m Metal) Reflect(lightRay *LightRay, rng RNG) {
	reflected := reflectVec(lightRay.ray.Direction.UnitVector(), lightRay.hitRecord.normal)
	lightRay.ray = Ray{lightRay.hitRecord.point, reflected}
	lightRay.color = lightRay.color.MulWithVec3(m.albedo)
}
