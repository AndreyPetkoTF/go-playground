package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpu", "", "write cpu profile to file `file`")
	memprofile = flag.String("mem", "", "write mem profile to `file`")
)

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()

		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	//fmt.Println(myPow(2.00000, -1))

	for i := 0; i < 100000000; i++ {
		fastPowMain(2, 4)
	}

	fmt.Println("finish")
}

func myPow(x float64, n int) float64 {
	var res float64 = 1

	for i := 0; i < abs(n); i++ {
		res *= x
	}

	if n >= 0 {
		return res
	} else {
		return float64(1) / res
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func fastPowMain(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}

	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
