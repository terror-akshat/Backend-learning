package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate us 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	//conversion string to int, and so many things.

	// This will lead an unidentified space because we are hitting the enter
	// after entering the number so we use strings.TrimSpace method for that
	// numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}
}
