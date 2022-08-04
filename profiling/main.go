package main

import (
	"bufio"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

/**
Make sure that you have: `brew install graphviz`

then run for profiling:
./profiling -cpuprofile cpu.prof -memprofile mem.prof ./testdata/*
to generate profiling files
then run `go tool pprof cpu.prof` to open specific file which you want to debug and run 'png' there

Or just:
go tool pprof -png cpu.prof

Also, you can just add import _ "net/http/pprof" in import block and run it if you have web server

top10 - fast check biggest one

*/

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file `file`")
	memprofile = flag.String("memprofile", "", "write mem profile to `file`")
	maxWorkers = flag.Int("w", 4, "max workers")
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

	version4(*maxWorkers)
	//version1()
}

func version1() {
	if len(flag.Args()) == 0 {
		log.Error("No files process")
		return
	}

	result := make(map[string]int)

	start := time.Now()
	for _, fn := range flag.Args() {
		processFile(result, fn)
	}

	defer fmt.Printf("Processing took: %v\n\n", time.Since(start))
	//printResult(result)
}

func processFile(result map[string]int, fn string) {
	var w string
	r, err := os.Open(fn)
	if err != nil {
		log.Warn(err)
		return
	}
	defer r.Close()

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		w = strings.ToLower(sc.Text())
		result[w] = result[w] + 1
	}
}

func printResult(result map[string]int) {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "---")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
