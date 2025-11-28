package main		


import "fmt"	

func main() {
    // 方式1，类型推导，用得最多
    a := 1
    // 方式2，完整的变量声明写法
    var b int = 2
    // 方式3，仅声明变量，但是不赋值，
    var c int
    fmt.Println(a, b, c)
	fmt.Println(method2(a, b))
}

// 方式4，直接在返回值中声明
func method2(a,b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
    // 这种方式必须声明return关键字
    // 并且同样不需要使用，并且也不用必须给这种变量赋值
    return 
}

func method3() (a int, b string) {
    a = 1
    b = "test"
    return
}

func method4() (a int, b string) {
    return
}


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