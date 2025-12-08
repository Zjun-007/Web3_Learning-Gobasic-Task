// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

package main

import (
	"fmt"
	"sync"
)		
func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup	
	// 启动10个协程，每个协程对计数器进行1000次递增操作		
	numGoroutines := 10
	incrementsPerGoroutine := 1000	
	wg.Add(numGoroutines)		
	for i := 0; i < numGoroutines; i++ {	
		go func() {
			defer wg.Done()		
			mu.Lock()
			for j := 0; j < incrementsPerGoroutine; j++ {					
				counter++ // 递增计数器	
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}	