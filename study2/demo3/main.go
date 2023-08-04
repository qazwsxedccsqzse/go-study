package main

import (
	"fmt"
	"sync"
)

func main() {
	str := []byte("foobar")
	ch := make(chan byte, len(str))
	next := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(str); i++ {
		ch <- str[i]
	}

	next <- 1

	close(ch)

	go func() {
		defer wg.Done()
		for {
			nextSignal, ok := <-next
			if !ok {
				return
			}
			if v, ok := <-ch; ok {
				fmt.Println("goroutine01", string(v))
			} else {
				// 會造成兩邊都關 channel
				close(ch)
				return
			}
			next <- nextSignal
		}
	}()

	go func() {
		defer wg.Done()
		for {
			nextSignal, ok := <-next
			if !ok {
				return
			}
			if v, ok := <-ch; ok {
				fmt.Println("goroutine02", string(v))
			} else {
				// 會造成兩邊都關 channel
				close(ch)
				return
			}
			next <- nextSignal
		}
	}()

	wg.Wait()
}
