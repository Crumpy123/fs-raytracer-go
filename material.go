package main

type Material interface {
	Scatter(ray *Ray, hRecord *HitRecord, attenuation *Vec3, scattered *Ray) bool
}

type Lambertian struct {
	albedo Vec3
}

type Metal struct {
	albedo Vec3
}

func (m Lambertian) Scatter(ray *Ray, hRecord *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	scatterDirection := hRecord.normal.Add(randUnitVector())

	//catch degenerate scatter direction
	if scatterDirection.nearZero() {
		scatterDirection = hRecord.normal
	}

	*scattered = Ray{hRecord.p, scatterDirection}
	*attenuation = m.albedo
	return true
}

func (m Metal) Scatter(ray *Ray, hRecord *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := reflectVec(ray.Direction.UnitVector(), hRecord.normal)
	*scattered = Ray{hRecord.p, reflected}
	*attenuation = m.albedo
	return Dot(scattered.Direction, hRecord.normal) > 0
}
