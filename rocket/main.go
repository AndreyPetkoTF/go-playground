package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	countdown5()
}

func countdown() {
	fmt.Println("start countdown")

	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

	fmt.Println("hello")
}

func countdown2() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("start countdown")

	select {
	case <-time.After(5 * time.Second):
	case <-abort:
		fmt.Println("Cancel")
		return
	}

	fmt.Println("Hello")
}

func countdown3() {
	ch := make(chan int, 1) // empty

	// при i == 0 записали 0
	// при i == 1 прочитали и вывели 0
	// при i == 2 записали 2
	// при i == 3 прочитали и вывели 2
	// .....

	// работает только для буфера 1
	// так как если он не полный и не пустой то case будет выбран случайным образом

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // вывели 0
			fmt.Println(x)
		case ch <- i: // п записано 0
		}
	}
}

func countdown4() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("start countdown")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("aborted")
			return

		}
	}

	// В данном случае происходит утечка goroutine так как после цикла никто не "слушает" ticker,
	//но он все еще продолжает работать и отправлять тики. Пример с исправленной этой проблемой в countdown5

	fmt.Println("launch")
}

func countdown5() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	ticker := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-ticker.C:
		case <-abort:
			fmt.Println("Canceled")
			return
		}
	}

	ticker.Stop()
	fmt.Println("launch")
}
