package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array")
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	// fruitList[2] = "tomato"  ///  it give the blank space if we are skipping that postion.
	fruitList[3] = "banana"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var veg = [4]string{"ptato", "onion", "carrot", "cabbage"}
	fmt.Println(veg)
	fmt.Println(len(veg))
}
