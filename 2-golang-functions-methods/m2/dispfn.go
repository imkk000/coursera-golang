package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v * t) + s
	}
}

func main() {
	var a, v, s, t float64
	fmt.Print("Enter acceleration value: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity value: ")
	fmt.Scan(&v)
	fmt.Print("Enter initial displacement value: ")
	fmt.Scan(&s)
	fmt.Print("Enter time value: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v, s)
	fmt.Printf("\nThe displacement after %.0f seconds: %f", t, fn(t))
}
