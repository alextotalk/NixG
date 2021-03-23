package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func Structs() {
	c := Circle{0, 0, 6}
	//c:=new(Circle)
	//c.x=5
	//fmt.Println(c)
	//c.x = 10
	//c.y = 5
	//fmt.Println(c.x, c.y, c.r)
	//fmt.Println(circleArea(&c))
	fmt.Println(c.area())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
