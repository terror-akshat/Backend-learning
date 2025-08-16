package testpackage

import "fmt"

// to exporting this first word of function should be capital.
func MyFuntion(step int) { //taking parameter step as int type
	if step == 1 {
		fmt.Println("step==", step)
		return
	}
	fmt.Println("step==2")
	fmt.Println("step==3")

	myPrivateFunction()
}

func myPrivateFunction() {
	fmt.Println("This is a private function and cannot be accessed outside this package.")
}
