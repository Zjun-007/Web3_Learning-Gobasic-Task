// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。


package main

import (
	"fmt"
)
// 增加10的函数，接收一个整数指针
func addTen(numPtr *int) {
	*numPtr += 10
}				
func main() {
	// 测试增加10的函数
	num := 5
	fmt.Println("Before addTen:", num)
	numPtr := &num
	 addTen(numPtr) //numPtr := &num  //(*int)   *numPtr = num + 10
	fmt.Println("After addTen:", num) // num = 15	
}	