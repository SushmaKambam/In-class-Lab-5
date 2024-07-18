package main

import "fmt"

func TestSqrt() {
	number := 16.0
	squareRoot := Sqrt(number)
	fmt.Printf("The square root of %.2f is %.2f\n", number, squareRoot)

	// You can call the Sqrt function with other numbers too
	number2 := 25.0
	squareRoot2 := Sqrt(number2)
	fmt.Printf("The square root of %.2f is %.2f\n", number2, squareRoot2)
}
