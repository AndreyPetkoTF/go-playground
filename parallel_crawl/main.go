package main

import (
	"flag"
	"fmt"
	"go-playground/parallel_crawl/links"
	"log"
	"os"
)

var depth *int

func init() {
	depth = flag.Int("depth", 3000, "depth")
}

type Links struct {
	Links []string
	Depth int
}

func main() {
	flag.Parse()

	workList := make(chan Links)
	var n int // количество ожидающий отправки в список

	n++
	go func() { workList <- Links{os.Args[2:3], 1} }()

	seen := make(map[string]bool)
	for list := range workList {
		//for ; n > 0; n-- {
		//	list := <-workList
		for _, link := range list.Links {
			if !seen[link] {
				fmt.Println(link)
				seen[link] = true
				n++

				if list.Depth < *depth {
					go func(link string) {
						workList <- crawlDepth(link, list.Depth)
					}(link)
				}
			}
		}
	}
}

var limitRunCrawlersAtSameTime = 20
var tokens = make(chan struct{}, limitRunCrawlersAtSameTime)

func crawlDepth(url string, depth int) Links {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}

	return Links{list, depth + 1}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

/**
until on of goroutines will finish and run <-tokens - another, which more then limit, will be blocked on send to tokens channel `tokens <- struct{}{}`
*/
func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}

	return list
}

func mainCrawl3() {
	worklist := make(chan []string)
	unseenList := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenList {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenList <- link
			}
		}
	}
}
