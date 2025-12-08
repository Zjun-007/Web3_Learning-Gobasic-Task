package main	
import "fmt"

// func func1(a int, b int) (int, int) {
//  	 return a, b
// 	}

// func main() {
//     s1 := []int{5, 4, 3, 2, 1}	
//     // 下标访问切片
// 	fmt.Println(s1)
//     // _ := s1[0]
//     // _ := s1[1]		
// 	// _ := s1[2]
// 	//fmt.Println("e1 = ", e1, "; e2 = ", e2, "; e3 = ", e3)

// 	// 使用range遍历切片
// 	for i, v := range s1 {
// 		fmt.Println("index = ", i, "value = ", v)
// 	}



// 	fmt.Println(func1(3,6))
// }

func main() {
// 声明一个空切片	
s3 := []int{}
fmt.Println("s3 = ", s3)

// append函数追加元素
s3 = append(s3)
s3 = append(s3, 1)
s3 = append(s3, 2, 3)
fmt.Println("s3 = ", s3)

//
s4 := []int{1, 2, 4, 5}
s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
fmt.Println("s4 = ", s4)

// 删除切片元素
s5 := []int{1,2,3,5,4}
s5 = append(s5[:3], s5[4:]...)
fmt.Println("s5 = ", s5)

// copy函数复制切片
src1 := []int{1, 2, 3}
dst1 := make([]int, 4, 5)

src2 := []int{1, 2, 3, 4, 5}
dst2 := make([]int, 3, 3)

fmt.Println("before copy, src1 = ", src1)
fmt.Println("before copy, dst1 = ", dst1)

fmt.Println("before copy, src2 = ", src2)
fmt.Println("before copy, dst2 = ", dst2)

// copy(dst1, src1)
copy(src1, dst1)
copy(dst2, src2)

fmt.Println("after copy, src1 = ", src1)
fmt.Println("after copy, dst1 = ", dst1)

fmt.Println("after copy, src2 = ", src2)
fmt.Println("after copy, dst2 = ", dst2)

}