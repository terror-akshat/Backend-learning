package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj32eks"

func main() {
	fmt.Println(myurl)
	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of qparams %T\n", qparams) // key value pair
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	// makeing an string to url.
	partOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/learn",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)
}
