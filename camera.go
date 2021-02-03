package main

type Camera struct {
	aspectRation, viewportHeight, viewportWidth, focalLength float64

	origin, horizontal, vertical, lowerLeftCorner Vec3
}

func (c *Camera) init() {
	c.aspectRation = 16.0 / 9.0
	c.viewportHeight = 2.0
	c.viewportWidth = c.aspectRation * c.viewportHeight
	c.focalLength = 1
	c.origin = Vec3{0, 0, 0}
	c.horizontal = Vec3{c.viewportWidth, 0, 0}
	c.vertical = Vec3{0, c.viewportHeight, 0}
	c.lowerLeftCorner = c.origin.Sub(c.horizontal.Divide(2.0)).Sub(c.vertical.Divide(2.0)).Sub(Vec3{0, 0, c.focalLength})
}

func (c *Camera) getRay(u, v float64) Ray {
	return Ray{c.origin, c.lowerLeftCorner.Add(c.horizontal.Mul(u)).Add(c.vertical.Mul(v)).Sub(c.origin)}
}
