package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle1 struct {
	Center Point
	Radius int
}

type Wheel1 struct {
	Circle Circle1
	Spokes int
}

type Circle2 struct {
	Point
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
	//X      int
}

func main() {
	var w Wheel1
	w.Circle.Center.X = 1
	w.Circle.Center.Y = 1
	w.Circle.Radius = 10
	w.Spokes = 10

	fmt.Printf("%#v\n", w)

	var w1 Wheel2
	w1.X = 1 //w1.Circle2.Point.X
	fmt.Printf("%#v\n", w1)
}
