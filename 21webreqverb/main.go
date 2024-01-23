package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const MYURL = "http://localhost:8000/get"

	response, err := http.Get(MYURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is : ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const MYURL = "http://localhost:8000/"

	requestBody := strings.NewReader(`
	{
		"courseName":"Go-lang",
		"price":999,
		"platform":"youtube"
	}`)

	response, err := http.Post(MYURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const MYURL = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "sujan")
	data.Add("lastname", "rijal")
	data.Add("email", "rijal@test.io")

	response, err := http.PostForm(MYURL, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
