package main

import (
	"fmt"
	"golang.org/x/image/bmp"
	"os"
	"time"
)

func main() {
	t1 := time.Now()
	f, err := os.Create("testBMP.BMP")
	if err != nil {
		fmt.Println("File creation failed")
	}
	image := traceTheRays()
	err = bmp.Encode(f, image)
	if err != nil {
		fmt.Println("err or in encoding")
	}

	delta := time.Now().Sub(t1).Seconds()
	fmt.Println(delta)

}
