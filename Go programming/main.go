package main

import (
	"fmt"
	// "os"
	// "io"
)
type Shape interface{
	Area() float64
}
type Rectangle struct{
	width, height float64
}
func (r Rectangle) Area() float64{
	return r.height * r.width
}
func CalculateArea(shape Shape) float64{
	return shape.Area()
}
func main() {
	rect := Rectangle{width: 5, height: 7}
	fmt.Println(CalculateArea(rect))
}
