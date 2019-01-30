package main

import "fmt"

func main() {
	//defining add3 and add10 to be functions in their own right
	add3 := adder(3)
	add10 := adder(10)

	fmt.Println(add3(10))
	fmt.Println(add10(10))
}

func adder(val int) func(int) int {
	return func(x int) int {
		return x + val
	}
}