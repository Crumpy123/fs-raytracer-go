package main

import "math"

type Camera struct {
	aspectRation, viewportHeight, viewportWidth, focalLength, vFov, theta, h, lensRadius float64

	origin, horizontal, vertical, lowerLeftCorner, lookFrom, lookAt, vUp, u, v, w Vec3
}

func (c *Camera) init(lookFrom Vec3, lookAt Vec3, vUP Vec3, vFov, aspectRatio, aperture, focusDist float64) {
	c.lookFrom = lookFrom
	c.lookAt = lookAt
	c.vUp = vUP
	c.vFov = vFov
	c.aspectRation = aspectRatio

	c.theta = degreeSToRadians(vFov)
	c.h = math.Tan(c.theta / 2)
	c.viewportHeight = 2.0 * c.h
	c.viewportWidth = c.aspectRation * c.viewportHeight

	c.w = c.lookFrom.SubVec3(c.lookAt).UnitVector()
	c.u = Cross(c.vUp, c.w).UnitVector()
	c.v = Cross(c.w, c.u)

	//c.focalLength = 1
	c.origin = c.lookFrom
	c.horizontal = c.u.Mul(c.viewportWidth).Mul(focusDist)
	c.vertical = c.v.Mul(c.viewportHeight).Mul(focusDist)
	c.lowerLeftCorner = c.origin.Sub(c.horizontal.Divide(2.0)).Sub(c.vertical.Divide(2.0)).Sub(c.w.Mul(focusDist))

	c.lensRadius = aperture / 2.0
}

func (c *Camera) getRay(s, t float64, rng RNG) Ray {
	rd := randomInUnitDisc(rng).Mul(c.lensRadius)
	offset := c.u.Mul(rd.x).Add(c.v.Mul(rd.y))

	return Ray{c.origin.Add(offset), c.lowerLeftCorner.Add(c.horizontal.Mul(s)).Add(c.vertical.Mul(t)).Sub(c.origin).Sub(offset)}
}
