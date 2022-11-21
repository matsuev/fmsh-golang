package main

import (
	"fmsh-golang/lesson-02/primitive"
	"fmt"
)

func main() {
	var p1, p2, p3 primitive.Primitive

	p1 = primitive.CreateNewPoint(1, 2, "red")
	p2 = primitive.CreateNewLine(2, 3, 4, 5, "black")
	p3 = primitive.CreateNewPoint(1, 2, "white")

	fmt.Println(p1, p2, p3)

	user := &Person{
		Fname: "Alex",
		Lname: "Matsuev",
		Age:   45,
	}

	fmt.Println(user)
}

type Person struct {
	Fname, Lname string
	Age          int
}

func (p *Person) String() string {
	return fmt.Sprintf("My name is %s %s. I'm %d age old.", p.Fname, p.Lname, p.Age)
}

// func Summ(a, b int64, result *int64) {
// 	*result = a + b
// }

// func Summ(a, b int) int {
// 	c := a + b
// 	return c
// }

// func Divide(a, b float64) (c float64, err error) {
// 	if b == 0 {
// 		err = errors.New("division by zero")
// 		return
// 	}
// 	c = a / b
// 	return
// }
