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
    print("θ", 5); fmt.Printf("  = %.2f° (%.3f rad)\n", self.degree, self.rad)
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
    alpha := self.rad + diff
    self.rad = alpha
    self.x = self.module * math.Cos(alpha)
    self.y = self.module + math.Sin(alpha)
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
    self.x = self.x * -1
    self.y = self.y * -1
    self.polar_from_cartesian()
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

