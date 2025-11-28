package main	
import "fmt"


func main() {

a := 2
var p *int
fmt.Println(&a)
p = &a
fmt.Println(p, &a)

var pp **int
pp = &p
fmt.Println(pp, p)
**pp = 3
fmt.Println(pp, *pp, p)
fmt.Println(**pp, *p)
fmt.Println(a, &a)
}