package main

type Settings struct {
	aspectRatio                                        float64
	imageHeight, imageWidth, samplesPerPixel, maxDepth int
}

func (i *Settings) declare() {
	i.aspectRatio = 16.0 / 9.0
	i.imageWidth = 400
	i.imageHeight = int(float64(i.imageWidth) / i.aspectRatio)
	i.samplesPerPixel = 10
	i.maxDepth = 50
}
