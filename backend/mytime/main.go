package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timing")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //exact formate we have ot provide no change.:
}
