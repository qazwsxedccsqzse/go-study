package main

import (
	"fmt"
	"sync"
)

func addByShareMemoryFail(n int) []int {
	var ints []int
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			ints = append(ints, i)
		}(i)
	}

	wg.Wait()
	return ints
}

func addByShareMemory(n int) []int {
	var ints []int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()
		}(i)
	}

	wg.Wait()

	return ints
}

func addByShareCommunicate(n int) []int {
	var ints []int
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(channel chan<- int, order int) {
			channel <- order
		}(channel, i)
	}

	for i := range channel {
		ints = append(ints, i)

		if len(ints) == n {
			break
		}
	}

	close(channel)

	return ints
}

func main() {

	// 這裡會因為 goroutine 針對同個記憶體位置進行讀寫
	// 會因為時間差造成覆寫問題
	foo := addByShareMemoryFail(10)
	fmt.Println(len(foo))
	fmt.Println(foo)

	// 1. 透過上鎖解決
	// foo := addByShareMemory(10)
	// fmt.Println(len(foo))
	// fmt.Println(foo)

	// 2. 透過 share memory by communicating
	// Golang 名言
	// Do not communicate by sharing memory; instead, share memory by communicating.
	// foo = addByShareCommunicate(10)
	// fmt.Println(len(foo))
	// fmt.Println(foo)
}
