package main

import "fmt"
type Shape interface {
	draw()
}
type Color interface {
	colorUp()
}
type Square struct {
}
type Rectangle struct {
}
type Circle struct {
}
type Red struct {
}
type Yellow struct {
}

func (Square) draw() {
	fmt.Print("画一个正方形")
}
func (Rectangle) draw() {
	fmt.Print("画一个矩阵")
}
func (Circle) draw() {
	fmt.Print("画一个圆圈")
}
func (Red) colorUp() {
	fmt.Print("红色")
}
func (Yellow) colorUp() {
	fmt.Print("黄色")
}

type abstractFactory interface {
	 getShape() Shape
	 getColor() Color

}
type RedCircleFactory struct {

}

func (RedCircleFactory) getShape() Shape {
	return new(Circle)
}
func (RedCircleFactory) getColor() Color {
	return new(Red)
}
func main(){
	new(RedCircleFactory).getShape().draw()
	new(RedCircleFactory).getColor().colorUp()
}

