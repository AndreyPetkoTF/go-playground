package main

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//runSignal()
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()

		//ctx, stop := context.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		//defer stop()

		//fmt.Println("run cancel goroutine")
		//time.Sleep(2 * time.Second)
		//fmt.Println("run cancel")
		//cancel()
	}()

	//run(ctx)
	runWaitGroup(ctx)

}

func run(ctx context.Context) {
	wait := make(chan struct{}, 1)

	go func() {
		defer func() {
			wait <- struct{}{}
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	go func() {
		defer func() {
			wait <- struct{}{}
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	<-wait
	<-wait

	fmt.Println("main done")
}

func runWaitGroup(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello from loop 1")
			}
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello from loop 2")
			}
		}
	}()

	wg.Wait()

	fmt.Println("Main done")
}

func runSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			fmt.Println("Break the loop")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Hello in a loop")
		}
	}
}
