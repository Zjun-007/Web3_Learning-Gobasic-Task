// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，
// 同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

package main
import (		
	"fmt"	
	"time"
	"sync"
)	
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()		
	start := time.Now()	
	// 模拟任务执行时间
	time.Sleep(time.Duration(100+id*50) * time.Millisecond) 
	duration := time.Since(start)	
	fmt.Printf("Task %d completed in %v\n", id, duration)
}
func main() {
	var wg sync.WaitGroup		
	tasks := 10
	// 启动多个任务协程			
	wg.Add(tasks)
	for i := 1; i <= tasks; i++ {
		go task(i, &wg)
	}
	wg.Wait() // 等待所有任务完成
}

