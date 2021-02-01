package main

type HittableList struct {
	objects []Hittable
}

func (h *HittableList) Add(newHittable Hittable) {
	h.objects = append(h.objects, newHittable)
}

func (h *HittableList) PClear() {
	h.objects = []Hittable{}
}

func (h *HittableList) HitSomething(ray *Ray, minT float64, maxT float64, rec *HitRecord) bool {
	tempRec := rec
	hitAnything := false
	closestSoFar := maxT

	for _, object := range h.objects {
		if object.Hit(ray, minT, closestSoFar, tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			rec = tempRec
		}
	}

	return hitAnything
}
