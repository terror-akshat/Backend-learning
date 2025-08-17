package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the Slice")
	var fruitList = []string{"apple", "peach", "banana", "tomato"}
	fmt.Printf("Type of the fruitsList is %T\n", fruitList)

	// In go-lang the array is dynamic we can add as many as value using slice method(append)
	fruitList = append(fruitList, "Mango")
	fmt.Println(fruitList)

	// It slice from very first element from the list. and return the rest of the string.
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// it give the value from the start index to last index which in my case is 2->(4-1)=>3
	fruitList = append(fruitList[2:3])
	fmt.Println(fruitList)

	//memory management next() and make() method

	//make() we have to pass 2 parameter, 1st is type, 2nd is value of len/
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 421
	highScores[2] = 221
	highScores[3] = 400

	// only this method make this array dyanamic and realocate the memory
	highScores = append(highScores, 333, 3434, 433)
	fmt.Println(highScores)

	sort.Ints(highScores) // default accessding method
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove from slices.
	var index int = 2
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
}
