package main

import "fmt"
import "math"


type Shape interface {
	area() float64

}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main(){
	r := Rectangle{20,10}
	c := Circle{4}
	fmt.Println("Rectangle Area =",getArea(r))
	fmt.Println("Circle Area =",getArea(c))

}

//执行结果：
//Rectangle Area = 200
//Circle Area = 50.26548245743669