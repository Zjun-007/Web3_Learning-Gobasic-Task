package main	
import "fmt"

func func1(a int, b int) (int, int) {
 	 return a, b
	}

func main() {
    s1 := []int{5, 4, 3, 2, 1}	
    // 下标访问切片
	fmt.Println(s1)
    // _ := s1[0]
    // _ := s1[1]		
	// _ := s1[2]
	//fmt.Println("e1 = ", e1, "; e2 = ", e2, "; e3 = ", e3)

	// 使用range遍历切片
	for i, v := range s1 {
		fmt.Println("index = ", i, "value = ", v)
	}



	fmt.Println(func1(3,6))
}