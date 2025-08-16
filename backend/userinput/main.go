package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our go lang learning")

	// comma ok systax. || err ok systax;
	// this is like try catch in other language
	input, err := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
	fmt.Printf("type of this rarting is %T\n", input)
	if err != nil {
		fmt.Println("Error reading input:", err)
	} else {
		fmt.Println("Input read successfully!")
	}
}
