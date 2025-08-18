package main

import "fmt"

func main() {
	fmt.Println("Welcome to the map")

	language := make(map[string]string) // [String]<- key
	language["JS"] = "javascripts"
	language["PY"] = "python"
	language["RB"] = "Ruby"

	fmt.Println(language)
	fmt.Println(language["JS"])

	//deleting

	delete(language, "RB")
	fmt.Println(language)
}
