// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// func appendInt(x []int, y int) []int {
//     var z []int
//     zlen := len(x) + 1
//     if zlen <= cap(x) {
//         // There is room to grow.  Extend the slice.
//         z = x[:zlen]
//     } else {
//         // There is insufficient space.  Allocate a new array.
//         // Grow by doubling, for amortized linear complexity.
//         zcap := zlen
//         if zcap < 2*len(x) {
//             zcap = 2 * len(x)
//         }
//         z = make([]int, zlen, zcap)
//         copy(z, x) // a built-in function; see text
//     }
//     z[len(x)] = y
//     return z
// }

// func main() {
// 	var x []int		
// 	for i := 0; i < 10; i++ {
// 		x = appendInt(x, i)
// 		fmt.Printf("%d cap=%d\t%v\n", i, cap(x), x)
// 	}
// }



// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
// func nonempty(strings []string) []string {
//     i := 0
//     for _, s := range strings {
//         if s != "" {
//             strings[i] = s
//             i++
//         } 
//     }
//     return strings[:i]
// }

// func main() {
// 	data := []string{"one", "", "three", "", "five"}
// 	fmt.Printf("%q\n", nonempty(data)) // '["one" "three" "five"]'
// }



//go 结构体与json示例
// type Movie struct {
//     Title  string
//     Year   int  `json:"released"`
//     Color  bool `json:"color,omitempty"`
//     Actors []string
// }

// var movies = []Movie{
//     {Title: "Casablanca", Year: 1942, Color: false,
//         Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
//     {Title: "Cool Hand Luke", Year: 1967, Color: true,
//         Actors: []string{"Paul Newman"}},
//     {Title: "Bullitt", Year: 1968, Color: true,
//         Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
//     // ...
// }

// func main() {
// data, err := json.MarshalIndent(movies, "", "    ")
// if err != nil {
//     log.Fatalf("JSON marshaling failed: %s", err)
// }
// fmt.Printf("%s\n", data)
// }



// //闭包  --创建函数工厂
// package main

// import "fmt"

// func multiplier(factor int) func(int) int {
//     // factor 被闭包捕获
//     fmt.Printf("创建 multiplier(factor=%d) 的闭包\n", factor)
//     return func(x int) int {
//         result := x * factor
//         fmt.Printf("闭包执行: %d × %d = %d\n", x, factor, result)
//         return result
//     }
// }

// func main() {
//     // 步骤1: 创建第一个闭包，factor=2
//     fmt.Println("=== 创建 double 闭包 ===")
//     double := multiplier(2)
    
//     // 步骤2: 创建第二个闭包，factor=3
//     fmt.Println("\n=== 创建 triple 闭包 ===")
//     triple := multiplier(3)
    
//     // 步骤3: 使用闭包
//     fmt.Println("\n=== 使用 double 闭包 ===")
//     fmt.Println(double(5))  // 10
    
//     fmt.Println("\n=== 使用 triple 闭包 ===")
//     fmt.Println(triple(5))  // 15
// }



// package main

// import (
//     "fmt"
//     "math"
// )

// type Circle struct {
//     radius float64
// }

// // 计算面积
// func (c Circle) Area() float64 {
//     return math.Pi * c.radius * c.radius
// }

// // 计算周长
// func (c Circle) Circumference() float64 {
//     return 2 * math.Pi * c.radius
// }

// // 工厂方法（非方法，是函数）
// func NewCircle(radius float64) Circle {
//     return Circle{radius: radius}
// }

// func main() {
//     c := Circle{radius: 5}
    
//     fmt.Printf("半径: %.2f\n", c.radius)
//     fmt.Printf("面积: %.2f\n", c.Area())
//     fmt.Printf("周长: %.2f\n", c.Circumference())
// }



///////////////////////////////////////////////////////////////////////
// package main

// import (
//     "fmt"
//     "math"
// )

// // 几何形状接口
// type Shape interface {
//     Area() float64
//     Perimeter() float64
// }

// // 圆形实现
// type Circle struct {
//     Radius float64
// }

// func (c Circle) Area() float64 {
//     return math.Pi * c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
//     return 2 * math.Pi * c.Radius
// }

// // 矩形实现
// type Rectangle struct {
//     Width, Height float64
// }

// func (r Rectangle) Area() float64 {
//     return r.Width * r.Height
// }

// func (r Rectangle) Perimeter() float64 {
//     return 2 * (r.Width + r.Height)
// }

// // 使用接口的函数
// func PrintShapeInfo(s Shape) {
//     fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
// }

// func main() {
//     circle := Circle{Radius: 5}
//     rect := Rectangle{Width: 3, Height: 4}
    
//     PrintShapeInfo(circle) // Area: 78.54, Perimeter: 31.42
//     PrintShapeInfo(rect)   // Area: 12.00, Perimeter: 14.00
    
//     // 接口切片
//     shapes := []Shape{circle, rect}
//     for _, shape := range shapes {
//         PrintShapeInfo(shape)
//     }
// }



// package main


// import (
//     "fmt"
//     "time"
// )

// func main() {
//     ch1 := make(chan string)
//     ch2 := make(chan string)

//     go func() {
//         time.Sleep(1 * time.Second)
//         ch1 <- "one"
//     }()

//     go func() {
//         time.Sleep(2 * time.Second)
//         ch2 <- "two"
//     }()

//     for i := 0; i < 2; i++ {
//         select {
//         case msg1 := <-ch1:
//             fmt.Println("received", msg1)
//         case msg2 := <-ch2:
//             fmt.Println("received", msg2)
//         }
//     }
// }



package main

import "fmt"

type PayMethod interface {
    Pay(int)
}

type CreditCard struct {
    balance int
    limit   int
}

func (c *CreditCard) Pay(amout int) {
    if c.balance < amout {
        fmt.Println("余额不足")
        return
    }
    c.balance -= amout
}

func anyParam(param interface{}) {
    fmt.Println("param: ", param)
}

func main() {
    c := CreditCard{balance: 100, limit: 1000}
    c.Pay(200)
    var a PayMethod = &c
    fmt.Println("a.Pay: ", a)

    var b interface{} = &c
    fmt.Println("b: ", b)

    anyParam(c)
    anyParam(1)
    anyParam("123")
    anyParam(a)
}
