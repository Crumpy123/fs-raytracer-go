package main

import (
	"fmt"
	"golang.org/x/image/bmp"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	pprofFile, err := os.Create("pprof.txt")
	if err != nil {
		fmt.Println("pprof file creation failed")
	}

	err = pprof.StartCPUProfile(pprofFile)
	if err != nil {
		fmt.Println("Pprof starting failed")
	}
	defer pprof.StopCPUProfile()

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
