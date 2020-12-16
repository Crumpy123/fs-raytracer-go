package main

import (
	"fmt"
	"golang.org/x/image/bmp"
	"os"
)

func main() {
	f, err := os.Create("testBMP.BMP")
	if err != nil {
		fmt.Println("File creation failed")
	}
	image := traceTheRays()
	err = bmp.Encode(f, image)
	if err != nil {
		fmt.Println("err or in encoding")
	}

}
