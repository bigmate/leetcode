package main

import "fmt"

type Point struct {
	x, y int
}

func CanFly(a, b Point) bool {
	if a.x > b.x || a.y > b.y {
		return false
	}

	if a.x == b.x && a.y == b.y {
		return true
	}

	return CanFly(Point{x: a.x + a.y, y: a.y}, b) || CanFly(Point{x: a.x, y: a.y + a.x}, b)
}

func main() {
	a := Point{
		x: 2,
		y: 5,
	}
	b := Point{
		x: 100,
		y: 100,
	}

	fmt.Println(CanFly(a, b))
}
