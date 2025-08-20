package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/terror-akshat"

func main() {
	fmt.Println("web")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	fmt.Printf("Response is of %T\n", response)

	defer response.Body.Close() // connection closed

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
