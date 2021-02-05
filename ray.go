package main

type LightRay struct {
	ray         Ray
	color       Vec3
	bounceLimit int
	hitRecord   HitRecord
	attenuation Vec3
}

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r *Ray) At(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (lightRay *LightRay) resetLightRay(setting *Settings) {
	lightRay.bounceLimit = setting.rayBounceLimit
	lightRay.color = Vec3{1, 1, 1}
}
