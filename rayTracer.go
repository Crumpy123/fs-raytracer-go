package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func writeColor(renderImage *image.RGBA, pixelColor *Vec3, set *Settings, x, y int) {
	scale := 1.0 / float64(set.samplesPerPixel)
	r := math.Sqrt(pixelColor.x * scale)
	g := math.Sqrt(pixelColor.y * scale)
	b := math.Sqrt(pixelColor.z * scale)
	//math.Round(r * 255)
	ir := uint8(255 * clamp(r, 0.0, 0.999))
	ig := uint8(255 * clamp(g, 0.0, 0.999))
	ib := uint8(255 * clamp(b, 0.0, 0.999))
	renderImage.SetRGBA(x, set.imageHeight-y-1, color.RGBA{R: ir, G: ig, B: ib, A: 255})
}

func rayColor(ray *Ray, world *HittableList, depth int) Vec3 {
	var rec HitRecord

	if depth <= 0 {
		return Vec3{0, 0, 0}
	}

	if world.HitSomething(ray, 0.001, math.MaxFloat64, &rec) {
		target := rec.p.Add(randInHemisphere(rec.normal))
		r := Ray{rec.p, target.Sub(rec.p)}
		return rayColor(&r, world, depth-1).Mul(.5)
	}

	unitDirection := ray.Direction.UnitVector()
	t := 0.5 * (unitDirection.y + 1.0)

	return Vec3{1.0, 1.0, 1.0}.Mul(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.Mul(t))
}

func processChunkOfImage(startRow int, set Settings, cam Camera, world *HittableList, renderImage *image.RGBA, ch chan int) {

	endRow := startRow + set.rowsPerChunk

	for ; startRow < endRow && startRow < set.imageHeight; startRow++ {
		//fmt.Println(startRow)
		for x := 0; x < set.imageWidth; x++ {
			pixelColor := Vec3{0, 0, 0}
			for s := 0; s < set.samplesPerPixel; s++ {
				u := (float64(x) + rand.Float64()) / (float64(set.imageWidth) - 1)
				v := (float64(startRow) + rand.Float64()) / (float64(set.imageHeight) - 1)

				ray := cam.getRay(u, v)
				pixelColor.AddInPlace(rayColor(&ray, world, set.maxDepth))
			}
			writeColor(renderImage, &pixelColor, &set, x, startRow)
		}
	}
	<-ch
}

func traceTheRays() image.Image {

	//Image
	var set Settings
	set.init()

	//World
	var world HittableList
	world.Add(Sphere{Vec3{0, 0, -1}, 0.5})
	world.Add(Sphere{Vec3{0, -100.5, -1}, 100})

	//Camera
	var cam Camera
	cam.init()

	//rendering
	renderImage := image.NewRGBA(image.Rect(0, 0, set.imageWidth, set.imageHeight))

	ch := make(chan int, runtime.NumCPU())

	for rowCount := 0; rowCount < set.imageHeight; rowCount += set.rowsPerChunk {
		//printProgress(set.imageHeight, rowCount)

		ch <- 0
		go processChunkOfImage(rowCount, set, cam, &world, renderImage, ch)

	}

	close(ch)

	for len(ch) > 0 {
		time.Sleep(time.Millisecond * 50)
	}

	return renderImage
}
