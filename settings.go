package main

type Settings struct {
	aspectRatio                                                      float64
	imageHeight, imageWidth, samplesPerPixel, maxDepth, rowsPerChunk int
}

func (i *Settings) init() {
	i.aspectRatio = 16.0 / 9.0
	i.imageWidth = 400
	i.imageHeight = int(float64(i.imageWidth) / i.aspectRatio)
	i.samplesPerPixel = 100
	i.maxDepth = 50
	i.rowsPerChunk = i.imageHeight
}
