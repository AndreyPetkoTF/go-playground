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

func version4(maxWorkers int) {
	if len(flag.Args()) == 0 {
		log.Error("No files process")
		return
	}

	workersWG := &sync.WaitGroup{}
	partialResults := make(chan map[string]int, maxWorkers)
	workQueue := make(chan string, maxWorkers)
	reducerWG := &sync.WaitGroup{}
	finalResult := make(map[string]int)

	start := time.Now()

	reducer(reducerWG, finalResult, partialResults)
	for i := 0; i < maxWorkers; i++ {
		processFile4(workersWG, partialResults, workQueue)
	}

	for _, fn := range flag.Args() {
		workQueue <- fn
	}

	close(workQueue)
	workersWG.Wait()
	close(partialResults)
	reducerWG.Wait()
	defer fmt.Printf("Processing took: %v\n", time.Since(start))
	//printResult(finalResult)

}

func processFile4(wg *sync.WaitGroup, results chan<- map[string]int, workQueue <-chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		var w string
		for fn := range workQueue {
			res := make(map[string]int)
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
				res[w] += 1
			}

			results <- res
		}
	}()
}

func reducer(wg *sync.WaitGroup, finalResult map[string]int, in <-chan map[string]int) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for res := range in {
			for k, v := range res {
				finalResult[k] += v
			}
		}
	}()
}
