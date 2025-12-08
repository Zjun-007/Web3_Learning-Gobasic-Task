// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，
// 另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。


package main

import (		
	"fmt"	
	"time"
	"sync"
)	
// 打印奇数的协程函数
func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()		
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd:", i)
		time.Sleep(100 * time.Millisecond) // 模拟一些工作
	}	
}
// 打印偶数的协程函数
func printEven(wg *sync.WaitGroup) {	
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("Even:", i)
		time.Sleep(150 * time.Millisecond) // 模拟一些工作
	}
}

func main() {
	var wg sync.WaitGroup		
	// 启动打印奇数和偶数的协程			
	wg.Add(2)	
	go printOdd(&wg)
	go printEven(&wg)		
	wg.Wait() // 等待奇数和偶数打印完成			
}				

				