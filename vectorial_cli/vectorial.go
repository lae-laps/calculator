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

func (self vector) polar_from_cart() (float64, float64) {
    module := math.Sqrt(self.x * self.x + self.y * self.y)
    theta := math.Atan(self.y / self.x)
    return module, theta
}

func (self vector) print_disposition() {
    degree_theta := self.theta * (180 / math.Pi)
    print("Vₓ", 1); fmt.Printf("  = %.2f\n", self.x)
    print("Vᵧ", 2); fmt.Printf("  = %.2f\n", self.y)
    print("θ", 3); fmt.Printf("   = %.2f rad (%.2f°)\n", self.theta, degree_theta)
    print("|v|", 3); fmt.Printf(" = %.2f\n", self.module)
}

func (self vector) set_cart(x, y float64) {
    self.x = x
    self.y = y
    module, theta := self.polar_from_cart()
    self.module = module
    self.theta = theta
    if show_changes_in_vectors {
        self.print_disposition()
    }
}

