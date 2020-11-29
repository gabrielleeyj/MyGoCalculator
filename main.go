package main

import "fmt"

// Add Function - takes in 2 variables of value int. returns an int.
func Add(a, b int) int {
	return a + b
}

// Subtract Function - takes in 2 variables of value int. returns an int.
func Sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(Add(3,4))
	fmt.Println(Sub(4,3))
}
