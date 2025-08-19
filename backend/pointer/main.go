package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers")
	// var one *int
	// fmt.Println("Value of the onepointer", one) // it out is nil

	myNumber := 23
	var ptr = &myNumber // referncing the memory

	fmt.Println("Value of the actual pointer is ", ptr)  // this will give the memory address
	fmt.Println("Value of the actual pointer is ", *ptr) //  this will give the value of that adderss

	*ptr = *ptr * 2 //it is modified the actual value.
	fmt.Println("new value is", myNumber)
}
