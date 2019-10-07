package main

import "fmt"
type Shape interface {
	draw()
}
type Square struct {
}
type Rectangle struct {
}
type Circle struct {
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
func getShape(str string) Shape{
	if(str=="square"){
		return new(Square)
	}
	if(str=="rectangle"){
		return new (Rectangle)
	}
	if(str=="circle"){
		return new(Circle);
	}
	return nil;
}
func main(){
	getShape("circle").draw()
}

