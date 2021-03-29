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
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
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
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
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
	fmt.Println(c.area(), "-Area Circle")
	fmt.Println(c.perimeter(), "-Perimeter Circle")

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area(), "-Area Rectangle")
	fmt.Println(r.perimeter(), "-Perimeter Rectangle")
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is, ", p.Name)
}

type Android struct {
	//Person Person
	Person
	Model string
}
type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, vol := range shapes {
		area += vol.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, vol := range m.shapes {
		area = area + vol.area()
	}
	return area
}
