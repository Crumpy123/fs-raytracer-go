package main

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
)

type RenderJob struct {
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

func setLightRayColor(lightRay *LightRay, world *HittableList, rng RNG) {
	for ; lightRay.bounceLimit > 0; lightRay.bounceLimit-- {
		//check for a hit
		if world.HitSomething(&lightRay.ray, 0.001, math.MaxFloat64, &lightRay.hitRecord) {
			//calculate scattering
			lightRay.hitRecord.material.Reflect(lightRay, rng)
			if Dot(lightRay.ray.Direction, lightRay.hitRecord.normal) <= 0 {
				return
			}
			lightRay.bounceLimit -= 1
		} else {
			// hit nothing
			t := 0.5 * (lightRay.ray.Direction.UnitVector().y + 1.0)

			lightRay.color = lightRay.color.MulWithVec3(Vec3{1.0, 1.0, 1.0}.Mul(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.Mul(t)))
			return
		}
	}

	if lightRay.bounceLimit <= 0 {
		lightRay.color = Vec3{0, 0, 0}
		return
	}

}

func renderPixelChunk(job *RenderJob) {
	rng := NewRNG()
	lightRay := LightRay{}

	endRow := job.startRow + job.set.rowsPerChunk
	for ; job.startRow < endRow && job.startRow < job.set.imageHeight; job.startRow++ {
		//fmt.Println(startRow)
		for x := 0; x < job.set.imageWidth; x++ {
			pixelColor := Vec3{0, 0, 0}
			for s := 0; s < job.set.samplesPerPixel; s++ {
				u := (float64(x) + rng.Float64()) / (float64(job.set.imageWidth) - 1)
				v := (float64(job.startRow) + rng.Float64()) / (float64(job.set.imageHeight) - 1)
				lightRay.resetLightRay(job.set)
				lightRay.ray = job.cam.getRay(u, v)
				setLightRayColor(&lightRay, job.world, rng)
				pixelColor.AddInPlace(lightRay.color)
			}
			writeColor(job.renderImage, &pixelColor, job.set, x, job.startRow)
		}
	}
}

func startWorker(workCh chan *RenderJob, wg *sync.WaitGroup) {
	for job := range workCh {
		renderPixelChunk(job)
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

	jobChannel := make(chan *RenderJob)
	workGroup := &sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		workGroup.Add(1)
		go startWorker(jobChannel, workGroup)
	}

	for rowCount := 0; rowCount < set.imageHeight; rowCount += set.rowsPerChunk {
		jobChannel <- &RenderJob{
			rowCount,
			&set,
			&cam,
			&world,
			renderImage,
		}
	}

	close(jobChannel)

	workGroup.Wait()

	return renderImage
}
