package main

import (
    "fmt"
    "math"
)

// TODO: edit struct to have <rad> and  <deg> / <degree>

type vector struct {
    x      float64
    y      float64
    rad    float64
    degree float64
    module float64
}

func rad_from_degrees(degrees float64) float64 {
    return degrees * (2 * math.Pi / 360)
}

func degrees_from_rad(rad float64) float64 {
    return rad * (180 / math.Pi)
}

func (self *vector) polar_from_cartesian() {
    self.module = math.Sqrt(self.x * self.x + self.y * self.y)
    self.rad = math.Atan(self.y / self.x)
    self.degree = degrees_from_rad(self.rad)
}

func (self *vector) print_disposition() {
    print("Vₓ", 196); fmt.Printf(" = %.3f\n", self.x)
    print("Vᵧ", 46); fmt.Printf(" = %.3f\n", self.y)
    print("θ", 5); fmt.Printf("  = %.3f° (%.3f rad)\n", self.degree, self.rad)
    print("v", 5); fmt.Printf("  = %.3f\n", self.module)
}

func (self *vector) set_cart(x, y float64) {
    self.x = x
    self.y = y
    self.polar_from_cartesian()
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

func (self *vector) rot(diff float64) {
    diff = rad_from_degrees(diff)
    self.rad = self.rad + diff
    self.x = self.module * math.Cos(self.rad)
    self.y = self.module * math.Sin(self.rad)
    self.polar_from_cartesian()
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

func (self *vector) elong(amount float64) {
    self.x = (self.module + amount) * math.Cos(self.rad)
    self.y = (self.module + amount) * math.Sin(self.rad)
    self.polar_from_cartesian()
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

func (self *vector) invert() {
    self.x = -self.x 
    self.y = -self.y
    self.polar_from_cartesian()
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

func (self *vector) flat(axis byte) {
    switch axis {
        case 'x':
            self.y = 0
            break
        case 'y':
            self.x = 0
            break
        case 'z':
            break
        default:
            printRuntimeError("Invalid character\n")
    }
    self.polar_from_cartesian()
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

func (self *vector) arc() float64 {
    arc := self.module * self.rad
    if show_changes_in_vectors {
       print("arc ⌀", 5); fmt.Printf(" = %.3f\n", arc)
    }
    return arc
}

func (self *vector) sector() float64 {
    sector := (self.rad * (self.module * self.module) / 2)
    if show_changes_in_vectors {
        print("sector ⌔", 5); fmt.Printf(" = %.3f\n", sector)
    }
    return sector
}

