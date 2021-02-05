package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
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

func traceTheRays() image.Image {

	//Image
	var set Settings
	set.init()

	//World

	/*
		mGround := Lambertian{Vec3{0.8, 0.8, 0.0}}
		mCenter := Lambertian{Vec3{0.1, 0.2, 0.5}}
		mLeft := Dielectric{1.5}
		mRight := Metal{Vec3{0.8, 0.6, 0.2}, 0.1}

		world.Add(Sphere{Vec3{0, -100.5, -1}, 100, mGround})
		world.Add(Sphere{Vec3{0, 0, -1}, 0.5, mCenter})
		world.Add(Sphere{Vec3{-1, 0, -1}, -0.45, mLeft})
		world.Add(Sphere{Vec3{1, 0, -1}, 0.5, mRight})

	*/

	var world HittableList
	world = getRandomScene()

	//Camera
	/*
		var cam Camera
		lookFrom := Vec3{3, 3, 2}
		lookAt := Vec3{0, 0, -1}
		vUp := Vec3{0, 1, 0}
		distToFocus := lookFrom.SubVec3(lookAt).Len()
		aperture := 2.0
	*/
	var cam Camera
	lookFrom := Vec3{13, 2, 3}
	lookAt := Vec3{0, 0, 0}
	vUp := Vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.1
	cam.init(lookFrom, lookAt, vUp, 20, set.aspectRatio, aperture, distToFocus)

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

func startWorker(workCh chan *RenderJob, wg *sync.WaitGroup) {
	for job := range workCh {
		renderPixelChunk(job)
	}
	wg.Done()
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
				lightRay.ray = job.cam.getRay(u, v, rng)
				setLightRayColor(&lightRay, job.world, rng)
				pixelColor.AddInPlace(lightRay.color)
			}
			writeColor(job.renderImage, &pixelColor, job.set, x, job.startRow)
		}
	}
}

func setLightRayColor(lightRay *LightRay, world *HittableList, rng RNG) {
	for ; lightRay.bounceLimit > 0; lightRay.bounceLimit-- {
		//check for a hit
		if world.HitSomething(&lightRay.ray, 0.001, math.MaxFloat64, &lightRay.hitRecord) {
			//calculate scattering
			lightRay.hitRecord.material.Reflect(lightRay, rng)
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
func getRandomScene() (world HittableList) {
	groundMaterial := Lambertian{Vec3{0.5, 0.5, 0.5}}
	world.Add(Sphere{Vec3{0, -1000, 0}, 1000, groundMaterial})
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := rand.Float64()
			center := Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if center.SubVec3(Vec3{4.0, 0.2, 0.0}).Len() > 0.9 {
				var material Material
				if chooseMaterial < 0.8 {
					//diffuse
					color := randVec3().MulWithVec3(randVec3())
					material = Lambertian{color}
					world.Add(Sphere{center, 0.2, material})
				} else if chooseMaterial < 0.95 {
					//metal
					color := randVec3().MulWithVec3(randVec3())
					fuzz := randFloatWBoundExpensive(0, 0.5)
					material = Metal{color, fuzz}
					world.Add(Sphere{center, 0.2, material})
				} else {
					//glass
					material = Dielectric{1.5}
					world.Add(Sphere{center, 0.2, material})
				}
			}
		}
	}
	material1 := Dielectric{1.5}
	world.Add(Sphere{Vec3{0, 1, 0}, 1.0, material1})

	material2 := Lambertian{Vec3{0.4, 0.2, 0.1}}
	world.Add(Sphere{Vec3{-4, 1, 0}, 1.0, material2})

	material3 := Metal{Vec3{0.7, 0.6, 0.5}, 0.0}
	world.Add(Sphere{Vec3{4, 1, 0}, 1.0, material3})
	return
}
