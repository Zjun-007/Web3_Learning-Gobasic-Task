// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
package main

import (
	"fmt"
	"time"
)	
func producer(ch chan<- int) {	
	for i := 1; i <= 100; i++ {
		ch <- i // 发送整数到通道
		time.Sleep(10 * time.Millisecond) // 模拟一些工作
	}		
	close(ch) // 关闭通道，表示不再发送数据
}	
func consumer(ch <-chan int) {
	for num := range ch { // 从通道接收整数，直到通道关闭
		fmt.Println("Received:", num)
	}
}	
func main() {
	ch := make(chan int, 10) // 创建一个带缓冲区大小为10的整数类型通道	
	// 启动生产者协程
	go producer(ch)
	// 启动消费者协程
	go consumer(ch)
	// 等待一段时间，确保所有协程完成
	time.Sleep(3 * time.Second)
}