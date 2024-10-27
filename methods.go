package main

import "fmt"

type rect struct {
	x, y int
}

func (r rect) area() int {
	return r.x * r.y
}

func (r *rect) perim() int {
	return 2 * (r.x + r.y)
}
func methods() {
	fmt.Println("Let's print something")

	r := rect{
		x: 3,
		y: 9,
	}
	fmt.Println("Let's calculate the area and perimeter of rectangle")
	fmt.Printf("the width is: %v and height is : %v\n", r.x, r.y)
	fmt.Println(r.area())
	fmt.Println(r.perim())
}
