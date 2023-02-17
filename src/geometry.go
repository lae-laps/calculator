// circle UI

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func (self *line) display() {
	print(" [ ", 0)
	fmt.Printf("y = %.1fx + %.1f", self.gradient, self.intersect)
	print(" ]\n", 0)
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
				printInfo("substitute this by the real help")
				break
			case "new":
				if len(expr) < 2 {
					printUserError("unknown type for new")
					continue
				}
				switch expr[1] {
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
						s := line{
							gradient,
							intersect,
						}
						name := expr[2]
						_, ok := lines[name]
						if ok {
							printUserError("object already exists")
							break
						}	
						lines[name] = s
						printInfo("created new object")
						//fmt.Printf("y = %.2fx + %.2f\n", s.gradient, s.intersect)
						break
					default:
						printUserError("unknown object")
				}
			case "show":
				if len(expr) < 3 {
					printUserError("values needed for viewing\nusage: show <type> <object>")
					break
				}
				switch expr[1] {
					case "line":
						value, ok := lines[expr[2]]
						if !ok {
							printUserError("object does not exist")
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

