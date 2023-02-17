// circle UI

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

var HELP_MSG = `Help for Geometry CLI

new - creates a new object
new <type> <name> <required data .. > :
	point <name> <x> <y>
	line <name> <gradient> <y intersect>
	circle <name> <center x> <center y> <radius>

posr - analizes the relative position of 2 objects
posr <obj type init> <object 1> <object 2>
	obj type init must be one of the following (and objects must be in the same order):
		- circle_line
		- point_line
`

type point struct {
	x float64
	y float64
}

type circle struct {
	center point
	radius float64
}

type line struct {
	gradient float64
	intersect float64
}

func (self *point) display() {
	print(" [ ", 0)
	fmt.Printf("(%.1f, %.1f)", self.x, self.y)
	print(" ]\n", 0)
}

func (self *line) display() {
	print(" [ ", 0)
	fmt.Printf("y = %.1fx + %.1f", self.gradient, self.intersect)
	print(" ]\n", 0)
}

func (self *circle) display() {
	line_1 := fmt.Sprintf("center: (%.1f, %.1f) ; radius: %.1f", self.center.x, self.center.y, self.radius)
	var x_term, y_term string
	if self.center.x == 0.0 {
		x_term = "x"
	} else {
		x_term = fmt.Sprintf("(x - %.1f)", self.center.x)
	}
	if self.center.y == 0.0 {
		y_term = "y"
	} else {
		y_term = fmt.Sprintf("(y - %.1f)", self.center.y)
	}
	line_2 := fmt.Sprintf("%s² + %s² = %.1f²", x_term, y_term, self.radius)

	if utf8.RuneCountInString(line_1) > utf8.RuneCountInString(line_2) {
		for utf8.RuneCountInString(line_1) > utf8.RuneCountInString(line_2) {
			line_2 = fmt.Sprintf("%s ", line_2)
		}
	} else if utf8.RuneCountInString(line_2) > utf8.RuneCountInString(line_1) {
		for utf8.RuneCountInString(line_2) > utf8.RuneCountInString(line_1) {
			line_1 = fmt.Sprintf("%s ", line_1)
		}
	}

	print(" / ", 0)
	fmt.Print(line_1)
	print(" \\ \n \\ ", 0)
	fmt.Print(line_2)
	print(" /\n", 0)
}

func (self *line) get_general() [3]float64 {
	// y = mx + n
	// mx - y + n + 0
	// return [m, -1, n]

	coeficients := [3]float64{self.gradient, -1, self.intersect}
	return coeficients
}

func distance_point_to_line(p point, s line) float64 {
	s_coefficients := s.get_general()
	return math.Abs((s_coefficients[0] * p.x) + (s_coefficients[1] * p.y) + s_coefficients[2]) / math.Sqrt(math.Pow(2, s_coefficients[0]) + math.Pow(2, s_coefficients[1]))
}

func check_relative_position_circle_line(c circle, s line) int {
	dist := distance_point_to_line(c.center, s)
	if dist > c.radius {
		return 0
	} else if dist == c.radius {
		return 1
	} else {
		return 2
	}
}

func geometry_ui() {

	lines := make(map[string]line)
	points := make(map[string]point)
	circles := make(map[string]circle)

	for ;; {
		expr := strings.Split(input("geometry> ", 248), " ")
		if len(expr) == 0 {
			printUserError("invalid expression")
			continue
		}
		switch expr[0] {
			case "exit", "quit", "q":
				return
			case "clear":
				clearScreen()
			case "help", "h":
				print(HELP_MSG, 8)
				break
			case "posr":
				if len(expr) < 4 {
					printUserError("not enough arguments")
					continue
				}
				switch expr[1] {
					case "circle_line":
						c, ok := circles[expr[2]]
						if !ok {
							printUserError("circle object does not exist")
							continue
						}
						s, ok := lines[expr[3]]
						if !ok {
							printUserError("line object does not exist")
							continue
						}
						var position string
						distance := distance_point_to_line(c.center, s)
						if distance > c.radius {
							position = "exterior"
						} else if distance == c.radius {
							position = "tangent"
						} else {
							position = "sec"
						}
						print("circle:\n", 8)
						c.display()
						print("line: ", 8)
						s.display()
						print("distance (center to s): ", 8)
						fmt.Println(distance)
						print("position: ", 8)
						fmt.Println(position)
						break
					case "point_line":
						p, ok := points[expr[2]]
						if !ok {
							printUserError("point object does not exist")
							continue
						}
						s, ok := lines[expr[3]]
						if !ok {
							printUserError("line object does not exist")
							continue
						}
						distance := distance_point_to_line(p, s)
						print("point: ", 8)
						p.display()
						print("line: ", 8)
						s.display()
						print("distance: ", 8)
						fmt.Println(distance)
				}
			case "new":
				if len(expr) < 2 {
					printUserError("unknown type for new")
					continue
				}
				switch expr[1] {
					case "point":
						if len(expr) < 5 {
							printUserError("values needed for initializing point\nusage: new point <name> <x> <y>")
							continue
						}
						var err error
						var x, y float64
						x, err = strconv.ParseFloat(expr[3], 32)
						if err != nil {
							printUserError("x must be valid float64")
							continue
						}
						y, err = strconv.ParseFloat(expr[4], 32)
						if err != nil {
							printUserError("y must be valid float64")
							continue
						}
						name := expr[2]
						_, ok := points[name]
						if ok {
							printUserError("object already exists")
							break
						}
						p := point{x, y}
						points[name] = p
						printInfo("created new object")
						break
					case "line":
						if len(expr) < 5 {
							printUserError("values needed for initializing line\nusage: new line <name> <gradient> <intercept>")
							continue
						}
						var err error
						var gradient, intersect float64
						gradient, err = strconv.ParseFloat(expr[3], 32)
						if err != nil {
							printUserError("gradient must be valid float64")
							continue
						}
						intersect, err = strconv.ParseFloat(expr[4], 32)
						if err != nil {
							printUserError("intercept must be valid float64")
							continue
						}
						name := expr[2]
						_, ok := lines[name]
						if ok {
							printUserError("object already exists")
							break
						}
						s := line{
							gradient,
							intersect,
						}
						lines[name] = s
						printInfo("created new object")
						//fmt.Printf("y = %.2fx + %.2f\n", s.gradient, s.intersect)
						break
					case "circle":
						if len(expr) < 6 {
							printUserError("values needed for initializing circle\nusage: new circle <name> <center x> <center y> <radius>")
							continue
						}
						var err error
						var center_x, center_y, radius float64
						center_x, err = strconv.ParseFloat(expr[3], 32)
						if err != nil {
							printUserError("x coordinate must be float64")
							continue
						}
						center_y, err = strconv.ParseFloat(expr[4], 32)
						if err != nil {
							printUserError("y coordinate must be float64")
							continue
						}
						radius, err = strconv.ParseFloat(expr[5], 32)
						if err != nil {
							printUserError("radius must be float64")
							continue
						}
						name := expr[2]
						_, ok := circles[name]
						if ok {
							printUserError("object already exists")
							break
						}
						c := circle{
							point{center_x, center_y},
							radius,
						}
						circles[name] = c
						printInfo("created new object")
					default:
						printUserError("unknown object")
				}
			case "show":
				if len(expr) < 3 {
					printUserError("values needed for viewing\nusage: show <type> <object>")
					break
				}
				switch expr[1] {
					case "point":
						value, ok := points[expr[2]]
						if !ok {
							printUserError("point does not exist")
							break
						}
						value.display()
					case "line":
						value, ok := lines[expr[2]]
						if !ok {
							printUserError("line does not exist")
							break
						}
						value.display()
					case "circle":
						value, ok := circles[expr[2]]
						if !ok {
							printUserError("circle does not exist")
							break
						}
						value.display()
					default:
						printUserError("invalid type")
				}
			default:
				printUserError("unknown command")
		}
	}

	s := line{34, 3}
	coefficients := s.get_general()
	fmt.Printf("%fx + %fy + %f = 0\n", coefficients[0], coefficients[1], coefficients[2])

}

