// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，
// 并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

package main
import (
	"fmt"
	"math"
)
// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}	

// Rectangle 结构体
type Rectangle struct {
	Width, Height float64	
}	
// Circle 结构体
type Circle struct {
	Radius float64
}	

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}		
// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64  {
	return 2 * (r.Width + r.Height)
}		

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}		
// Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64	 {
	return 2 * math.Pi * c.Radius
}	
	
func main() {
	// 创建 Rectangle 和 Circle 实例
	rect := Rectangle{Width: 5, Height: 10}	
	circ := Circle{Radius: 7}		
	// 调用它们的 Area 和 Perimeter 方法
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", rect.Area(), rect.Perimeter())
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", circ.Area(), circ.Perimeter())		
}	