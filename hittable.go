package main

type HitRecord struct{
	p,normal Vec3
	t float64
	frontFace bool
}

type Hittable interface {
	Hit(ray *Ray, minT float64, maxT float64, rec *HitRecord) bool

}

func (hr *HitRecord) setFaceNormal(ray *Ray, outwardNormal *Vec3){
	hr.frontFace = Dot(ray.Direction, *outwardNormal) < 0
	if hr.frontFace{
		hr.normal = *outwardNormal
	}
	hr.normal = outwardNormal.Multiply(-1.0)
}
