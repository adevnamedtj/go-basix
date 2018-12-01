package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	getArea() float64
}

type Triangle struct {
	Length float64
	Base   float64
}

type Square struct {
	LengthOfASide float64
}

func (s Square) getArea() float64 {
	return s.LengthOfASide * s.LengthOfASide
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.Length * t.Base
}

func printArea(s Shape) {
	fmt.Println(reflect.ValueOf(s), s.getArea())
}

func main() {
	square := Square{LengthOfASide: 63.4}
	triangle := Triangle{Length: 32.3, Base: 7.3}

	printArea(square)
	printArea(triangle)
}
