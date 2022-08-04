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

func version2() {
	if len(flag.Args()) == 0 {
		log.Error("No files process")
		return
	}

	result := make(map[string]int)
	resLock := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	start := time.Now()
	for _, fn := range flag.Args() {
		processFile2(wg, fn, result, resLock)
	}

	wg.Wait()

	defer fmt.Printf("Processing took: %v\n", time.Since(start))
	//printResult(result)
}

func processFile2(wg *sync.WaitGroup, fn string, result map[string]int, resLock *sync.Mutex) {
	wg.Add(1)
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

		for sc.Scan() {
			w = strings.ToLower(sc.Text())
			resLock.Lock()
			result[w] = result[w] + 1
			resLock.Unlock()
		}
	}()
}
