package main

import "fmt"

type Geometry interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func ShowArea(g Geometry) {
	fmt.Println(g.Area())
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 5, height: 10}
	ShowArea(c)
	ShowArea(r)
}

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.name, p.age)
// }

// func main() {
// 	p := Person{name: "John", age: 42}
// 	fmt.Println(p.String())
// }
