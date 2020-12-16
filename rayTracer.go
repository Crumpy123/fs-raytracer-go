package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

func printProgress(height, h int) {
	if math.Mod(((float64(height)-float64(h))/float64(height))*100, 1) == 0 {
		fmt.Println(int(((float64(height)-float64(h))/float64(height))*100), "%")
	}
}

func hitSphere(center Vec3, radius float64, ray Ray) float64 {
	oc := ray.Origin.PSubtract(center)
	a := ray.Direction.LengthSquared()//Dot(ray.Direction, ray.Direction)
	halfB := Dot(oc, ray.Direction)
	//b := Dot(oc, ray.Direction) * 2.0
	c := oc.LengthSquared() - radius*radius
	discriminant := halfB * halfB -a*c
	if discriminant < 0{
		return -1.0
	}else{
		return (-halfB-math.Sqrt(discriminant))/a//(-b - math.Sqrt(discriminant)) / (2.0*a)
	}
}

func rayColor(ray *Ray, world *HittableList) Vec3 {
	var rec HitRecord
	if world.HitSomething(ray,0,math.MaxFloat64, &rec){
		return rec.normal.Add(Vec3{1,1,1}).Multiply(0.5)
	}

	unitDirection := ray.Direction.UnitVector()
	t := 0.5 * (unitDirection.y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.Multiply(1.0-t).Add(Vec3{0.5, 0.7, 1.0}.Multiply(t))
}

func traceTheRays() image.Image {
	//Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	//World
	var world HittableList
	world.Add(Sphere{Vec3{0,0,-1}, 0.5})
	world.Add(Sphere{Vec3{0,-100.5, -1},100})

	//Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := Vec3{0, 0, 0}
	horizontal := Vec3{viewportWidth, 0, 0}
	vertical := Vec3{0, viewportHeight, 0}

	lowerLeftCorner := origin.Subtract(horizontal.Divide(2.0)).Subtract(vertical.Divide(2.)).Subtract(Vec3{0,0,focalLength})

	renderImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for y := imageHeight; y >= 0; y-- {
		printProgress(imageHeight, y)
		for x := 0; x < imageWidth; x++ {
			u := float64(x) / (float64(imageWidth) -1)
			v := float64(y) / (float64(imageHeight) -1)
			//	b := 0.25 * 255

			ray := Ray{origin, lowerLeftCorner.Add(horizontal.Multiply(u)).Add(vertical.Multiply(v)).Subtract(origin)}
			rtColor := rayColor(&ray, &world)

			ir := uint8(math.Round(rtColor.x * 255))
			ig := uint8(math.Round(rtColor.y * 255))
			ib := uint8(math.Round(rtColor.z * 255))
			renderImage.SetRGBA(x, imageHeight-y, color.RGBA{R: ir, G: ig, B: ib, A: 255})
		}
	}
	return renderImage
}