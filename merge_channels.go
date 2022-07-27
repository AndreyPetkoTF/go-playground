package main

//
//func main() {
//	for num := range joinChannels(a, b, c) {
//		fmt.Println(num)
//	}
//}
//
//func joinChannels(chs ...<-chan int) <-chan int {
//	mergedChan := make(chan int)
//
//	go func() {
//		wg := &sync.WaitGroup{}
//
//		wg.Add(len(chs))
//
//		for _, ch := range chs {
//
//		}
//	}
//}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	for _, c := range cs {
		go func() {
			for v := range c {
				out <- v
			}
		}()
	}

	return out
}
