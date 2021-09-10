package main

import (
	"fmt"
	"math"
)

type Shape interface {
	//计算面积
	area() float64
	//计算周长
	perimeter() float64
}

func GetArea(shape Shape) float64 {
	return shape.area()
}

func GetPeri(shape Shape) float64 {
	return shape.perimeter()
}

//Rect 矩形实现了Shape接口
type Rect struct {
	width, height float64
}

func (r *Rect) area() float64 {
	return r.width * r.height
}

func (r *Rect) perimeter() float64 {
	return (r.width + r.height) * 2
}

type Circle struct {
	radius float64
}

func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rectangle := Rect{
		width:  1,
		height: 2,
	}
	fmt.Println("矩形的周长为：", GetPeri(&rectangle))
	fmt.Println("矩形的面积为：", GetArea(&rectangle))

	circle := Circle{
		radius: 5,
	}
	fmt.Println("圆形的周长为：", GetPeri(&circle))
	fmt.Println("圆形的面积为：", GetArea(&circle))
}
