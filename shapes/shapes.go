package main

import "math"

// interface
type Shape interface {
	Area() float64
}

// 结构体的实现
type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

type Circle struct {
	Radius float64
}

func (circles Circle) Perimeter() float64 {
	return 2 * math.Pi * circles.Radius
}

func (circles Circle) Area() float64 {
	return math.Pi * circles.Radius * circles.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height / 2
}
