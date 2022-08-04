package main

import (
	"bufio"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
	"time"
)

func version3() {
	if len(flag.Args()) == 0 {
		log.Error("No files process")
		return
	}

	results := make(chan map[string]int)
	wg := &sync.WaitGroup{}

	start := time.Now()
	for _, fn := range flag.Args() {
		wg.Add(1)
		processFile3(fn, results, wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	result := make(map[string]int)
	for item := range results {
		for i, v := range item {
			result[i] += v
		}
	}

	defer fmt.Printf("Processing took: %v\n", time.Since(start))
	//printResult(result)
}

func processFile3(fn string, results chan<- map[string]int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		var w string

		r, err := os.Open(fn)
		if err != nil {
			log.Warn(err)
			return
		}
		defer r.Close()

		sc := bufio.NewScanner(r)
		sc.Split(bufio.ScanWords)

		fileResult := make(map[string]int)
		for sc.Scan() {
			w = strings.ToLower(sc.Text())
			fileResult[w] = fileResult[w] + 1
		}

		results <- fileResult
	}()
}
