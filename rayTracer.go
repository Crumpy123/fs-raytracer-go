package main

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
)

type Job struct {
	startRow    int
	set         *Settings
	cam         *Camera
	world       *HittableList
	renderImage *image.RGBA
}

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

func rayColor(ray *Ray, world *HittableList, depth int, rng RNG) Vec3 {
	var rec HitRecord

	if depth <= 0 {
		return Vec3{0, 0, 0}
	}

	//hit object
	if world.HitSomething(ray, 0.001, math.MaxFloat64, &rec) {
		var scattered Ray
		var attenuation Vec3
		if rec.material.Scatter(ray, &rec, &attenuation, &scattered, rng) {
			return attenuation.MulWithVec3(rayColor(&scattered, world, depth-1, rng))
		} else {
			return Vec3{0, 0, 0}
		}
	}

	// hit nothing
	unitDirection := ray.Direction.UnitVector()
	t := 0.5 * (unitDirection.y + 1.0)

	return Vec3{1.0, 1.0, 1.0}.Mul(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.Mul(t))
}

func traceWorker(workCh chan *Job, wg *sync.WaitGroup) {
	rng := NewRNG()

	for work := range workCh {
		endRow := work.startRow + work.set.rowsPerChunk
		for ; work.startRow < endRow && work.startRow < work.set.imageHeight; work.startRow++ {
			//fmt.Println(startRow)
			for x := 0; x < work.set.imageWidth; x++ {
				pixelColor := Vec3{0, 0, 0}
				for s := 0; s < work.set.samplesPerPixel; s++ {
					u := (float64(x) + rng.Float64()) / (float64(work.set.imageWidth) - 1)
					v := (float64(work.startRow) + rng.Float64()) / (float64(work.set.imageHeight) - 1)

					ray := work.cam.getRay(u, v)
					pixelColor.AddInPlace(rayColor(&ray, work.world, work.set.maxDepth, rng))
				}
				writeColor(work.renderImage, &pixelColor, work.set, x, work.startRow)
			}
		}
	}
	wg.Done()
}

func traceTheRays() image.Image {

	//Image
	var set Settings
	set.init()

	//World
	var world HittableList

	mGround := Lambertian{Vec3{0.8, 0.8, 0.0}}
	mCenter := Lambertian{Vec3{0.7, 0.3, 0.3}}
	mLeft := Metal{Vec3{0.8, 0.8, 0.8}}
	mRight := Metal{Vec3{0.8, 0.6, 0.2}}

	world.Add(Sphere{Vec3{0, -100.5, -1}, 100, mGround})
	world.Add(Sphere{Vec3{0, 0, -1}, 0.5, mCenter})
	world.Add(Sphere{Vec3{-1, 0, -1}, 0.5, mLeft})
	world.Add(Sphere{Vec3{1, 0, -1}, 0.5, mRight})

	//Camera
	var cam Camera
	cam.init()

	//rendering
	renderImage := image.NewRGBA(image.Rect(0, 0, set.imageWidth, set.imageHeight))

	workCh := make(chan *Job)
	finishing := &sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		finishing.Add(1)
		go traceWorker(workCh, finishing)
	}

	for rowCount := 0; rowCount < set.imageHeight; rowCount += set.rowsPerChunk {
		workCh <- &Job{
			rowCount,
			&set,
			&cam,
			&world,
			renderImage,
		}
	}

	close(workCh)

	finishing.Wait()

	return renderImage
}
