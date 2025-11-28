package main	
import "fmt"



type A struct {
    a string
}

type B struct {
    A
    b string
}

type C struct {
    A
    B
    a string
    b string
    c string
}

type people struct {
	name string
	parent *people  // 指向下一级别的指针
}

func main() {		
	grandpa := &people{name: "爷爷"}
	father := &people{name: "爸爸", parent: grandpa}
	son := &people{name: "儿子", parent: father}
fmt.Println("A struct:", A{"hello"})
fmt.Println("B struct:", B{A{"hello"}, "world"})
fmt.Println("C struct:", C{A{"a_hello"}, B{A{"b_hello"}, "b_world"}, "c_a", "c_b", "c_world"})
//fmt.Println("people struct:", people{"grandpa", &people{"father", &people{"son", nil}}})	
//fmt.Println("people struct:", people{"grandpa", &people{"father", &people{"son", nil}}})	
fmt.Println("son's father is:", son.parent.name)
fmt.Println("father's grandpa is:", son.parent.parent.name)	
}