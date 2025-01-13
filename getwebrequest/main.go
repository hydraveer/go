package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
	PerformPostRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("The status code is ", res.StatusCode)
	fmt.Println("Content lenght is ", res.ContentLength)

	var retrivedata strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := retrivedata.Write(content)

	fmt.Println(byteCount)
	fmt.Println("Content of this file is -> ", retrivedata.String())

	fmt.Println(string(content))
}
func PerformPostRequest() {
	const url = "http://localhost:3000/post"
	reqbody := strings.NewReader(`{
		"hi":"hi",
		"age":3
	}`)
	response, err := http.Post(url, "application/json", reqbody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	constent, _ := io.ReadAll(response.Body)

	fmt.Println(string(constent))
}
