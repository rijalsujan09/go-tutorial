package main

import (
	"fmt"
	"net/url"
)

const MYURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("welcome to handelling url in golang")

	fmt.Println(MYURL)

	result, _ := url.Parse(MYURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of the query params are : %T\n", qparams)
	// fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println(value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
