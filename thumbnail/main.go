package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//data, err := makeThumbnail([]string{
	//	"lala.png",
	//	"gg.png",
	//	"kreakrea.png",
	//})
	//
	//fmt.Println(data)
	//fmt.Println(err)

	value := makeThumbnailWithWaitGroup([]string{
		"lala.png",
		"gg.png",
		"kreakrea.png",
	})

	fmt.Println(value)
}

func ImageFile(infile string) (string, error) {
	seconds := rand.Intn(5) + 1

	time.Sleep(time.Duration(seconds) * time.Second)

	fmt.Println("finish" + strconv.Itoa(seconds))
	return "foo" + strconv.Itoa(seconds) + ".thump.jpg", nil
}

type item struct {
	thumbnail string
	err       error
}

func makeThumbnail(filenames []string) ([]string, error) {
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbnail, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	var thumbfiles []string
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}

		thumbfiles = append(thumbfiles, it.thumbnail)
	}

	return thumbfiles, nil
}

//func makeThumbnailWithWaitGroup(filenames <-chan string) int64 {
func makeThumbnailWithWaitGroup(filenames []string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()

			_, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}

			//info, _ := os.Stat(thump)
			sizes <- 5
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
