package main

import "fmt"

func main() { //main entry point
	fmt.Println("function")
	greeter()

	result := adder(2, 5)
	fmt.Println("Result is ", result)
	//function inside a function in not allowed in go.

	result2, val := propAdder(2, 3, 4, 4, 5, 6, 7)
	fmt.Println("Result2 is ", result2, val)
}

func adder(a int, b int) int { // int is the function signature to return back
	return a + b
}

func propAdder(values ...int) (int, string) { // if we want to return 2 value
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hii proAdder"
}

// func propAdder(values ...int) int {
// 	total := 0
// 	for _, val := range values {
// 		total += val
// 	}
// 	return total
// }

func greeter() {
	fmt.Println("Namstey")
}
