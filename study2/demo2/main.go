package main

import (
	"fmt"
	"time"
)

// 限制 channel 只能寫入資料
func Write(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 100) // 等待一下以便观察
	}
	close(ch) // 关闭通道，告知读取端没有更多的数据可以读取
}

// 限制 channel 只能讀出資料
func Read(ch <-chan int) {
	// 這行會噴錯, 原因是因為這裡只能讀取
	// ch <- 1
	for value := range ch { // 使用 range 迭代通道，直到通道被关闭
		fmt.Println("读取:", value)
	}
}

func main() {
	ch := make(chan int)

	go Write(ch) // 启动写入协程
	go Read(ch)  // 启动读取协程

	time.Sleep(time.Second) // 等待一秒，以确保协程完成
}
