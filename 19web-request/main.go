package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	fmt.Println("LCO  web request")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type :  %T\n", response)

	defer response.Body.Close() //closing connection

	dataByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)

	fmt.Println(content)
}
