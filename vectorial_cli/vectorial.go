package main

import (
    "fmt"
    "math"
)

type vector struct {
    x      float64
    y      float64
    theta  float64
    module float64
}

func rad_from_degrees(degrees float64) float64 {
    return degrees * (2 * math.Pi / 360)
}

func (self *vector) polar_from_cart() {
    self.module = math.Sqrt(self.x * self.x + self.y * self.y)
    self.theta = math.Atan(self.y / self.x)
}

func (self *vector) print_disposition() {
    degree_theta := self.theta * (180 / math.Pi)
    print("Vₓ", 1); fmt.Printf("  = %.2f\n", self.x)
    print("Vᵧ", 2); fmt.Printf("  = %.2f\n", self.y)
    print("θ", 3); fmt.Printf("   = %.2f rad (%.2f°)\n", self.theta, degree_theta)
    print("|v|", 3); fmt.Printf(" = %.2f\n", self.module)
}

func (self *vector) set_cart(x, y float64) {
    self.x = x
    self.y = y
    self.polar_from_cart()
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

func (self *vector) rot(diff float64) {
    diff = rad_from_degrees(diff)
    alpha := self.theta + diff
    self.theta = alpha
    self.x = self.module * math.Cos(alpha)
    self.y = self.module + math.Sin(alpha)
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

