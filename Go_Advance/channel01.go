// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
package main

import (
	"fmt"
	"time"
)	

func generateNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i // 发送整数到通道
		time.Sleep(100 * time.Millisecond) // 模拟一些工作
	}
	close(ch) // 关闭通道，表示不再发送数据
}
func printNumbers(ch <-chan int) {
	for num := range ch { // 从通道接收整数，直到通道关闭
		fmt.Println("Received:", num)
	}
}
func main() {
	ch := make(chan int) // 创建一个整数类型的通道	
	// 启动生成数字的协程
	go generateNumbers(ch)
	// 启动打印数字的协程
	go printNumbers(ch)
	// 等待一段时间，确保所有协程完成
	time.Sleep(2 * time.Second)
}	