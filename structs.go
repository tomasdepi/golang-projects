package main

import (
	"fmt"
)

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

type Student2 struct {
	name   string
	grades []int
}

func (s *Student2) displayName() {
	fmt.Println(s.name)
}

func (s *Student2) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

type Circle struct {
	x      int
	y      int
	radius int
}

type Circle2 struct {
	radius float64
	area   float64
}

func (c *Circle2) calcArea() {
	c.area = 3.14 * c.radius * c.radius
}

// a struct needs to implement all methods to implement the interface
type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

// reacts implement the shape interface
type rect struct {
	length, breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	var depi Student
	fmt.Printf("%+v \n", depi) // full of zero values

	st := new(Student)
	fmt.Printf("%+v \n", st)

	st2 := Student{"Tomas", 23, []int{1, 4, 6}, map[string]int{"math": 10}}
	fmt.Printf("%+v \n", st2)

	var c Circle
	c.x = 5
	c.y = 5
	c.radius = 6
	fmt.Printf("%+v \n", c)

	c2 := Circle2{radius: 5}
	fmt.Println(c2.area)
	c2.calcArea()
	fmt.Println(c2.area)

	s := Student2{name: "Joe", grades: []int{90, 75, 80}}
	s.displayName()
	fmt.Printf("%.2f%%", s.calculatePercentage())

	r := rect{length: 3, breadth: 4}
	c3 := square{side: 5}
	printData(r)
	printData(c3)
}
