package main

type Settings struct {
	aspectRatio                                                            float64
	imageHeight, imageWidth, samplesPerPixel, rayBounceLimit, rowsPerChunk int
}

func (i *Settings) init() {
	i.aspectRatio = 3.0 / 2.0
	i.imageWidth = 1200
	i.imageHeight = int(float64(i.imageWidth) / i.aspectRatio)
	i.samplesPerPixel = 100
	i.rayBounceLimit = 50
	i.rowsPerChunk = 1 //i.imageHeight / runtime.NumCPU()
}
