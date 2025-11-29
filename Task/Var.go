package main		


import "fmt"	

// func main() {
//     // 方式1，类型推导，用得最多
//     a := 1
//     // 方式2，完整的变量声明写法
//     var b int = 2
//     // 方式3，仅声明变量，但是不赋值，
//     var c int
//     fmt.Println(a, b, c)
// 	fmt.Println(method2(a, b))
// }

// // 方式4，直接在返回值中声明
// func method2(a,b int) (sum int, diff int) {
// 	sum = a + b
// 	diff = a - b
//     // 这种方式必须声明return关键字
//     // 并且同样不需要使用，并且也不用必须给这种变量赋值
//     return 
// }

// func method3() (a int, b string) {
//     a = 1
//     b = "test"
//     return
// }

// func method4() (a int, b string) {
//     return
// }


/*
var a, b, c int = 1, 2, 3

var e, f, g int

var h, i, j = 1, 2, "test"

func main() {
    var k, l, m int = 1, 2, 3
    var n, o, p int
    q, r, s := 1, 2, "test"
    fmt.Println(k, l, m, n, o, p, q, r, s)
}
	*/

var s1 complex64 = 1 + 2i
var s string = ("hello world!")
var bytes []byte = []byte(s)	

var s2 string = "abc，你好，世界！"
var runes []rune = []rune(s2)	

var s3 string = "Hello\nworld!\n"
var s4 string = `Hello
world!
`

var s5 string = "Go语言"
var bytes2 []byte = []byte(s5)
var runes2 []rune = []rune(s5)


func main() {
	fmt.Println(s1)		
	x := real(s1)
	y := imag(s1)
	fmt.Println(x, y)
	fmt.Println(bytes,len(s))
	fmt.Println(runes, len(s2))
	fmt.Println(s3)
	fmt.Println(s4)	
	fmt.Println(s3==s4)
	fmt.Println(bytes2,len(s5))
	fmt.Println(runes2, len(s5))
	fmt.Println("string sub: ", s5[0:8])
    fmt.Println("bytes sub: ", string(bytes2[0:8]))
    fmt.Println("runes sub: ", string(runes2[0:4]))
}