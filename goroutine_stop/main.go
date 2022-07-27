package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("tick")
			case <-quit:
				fmt.Println("stop goroutine")
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(10 * time.Second)
	quit <- true
}
