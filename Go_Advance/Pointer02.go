// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

package main
import (
	"fmt"
)

func NumMultiplyByTwo(slicePtr *[]int) {
	for i := range *slicePtr {
		(*slicePtr)[i] *= 2
	}
}
func main() {
	// 测试将切片中的每个元素乘以2的函数
	slice := []int{1, 2, 3, 4, 5}	
	fmt.Println("Before NumMultiplyByTwo:", slice)
	NumMultiplyByTwo(&slice)
	fmt.Println("After NumMultiplyByTwo:", slice)
}
