package main

type Settings struct {
	aspectRatio                                                            float64
	imageHeight, imageWidth, samplesPerPixel, rayBounceLimit, rowsPerChunk int
}

func (i *Settings) init() {
	i.aspectRatio = 16.0 / 9.0
	i.imageWidth = 400
	i.imageHeight = int(float64(i.imageWidth) / i.aspectRatio)
	i.samplesPerPixel = 100
	i.rayBounceLimit = 50
	i.rowsPerChunk = 1 //i.imageHeight / runtime.NumCPU()
}
