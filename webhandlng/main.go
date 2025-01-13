package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	response, err := http.Get(url)
	nilErr(err)
	fmt.Println("The response from web is", response)
	fmt.Printf("The response type is %T \n", response)
	defer response.Body.Close()
	res, err := io.ReadAll(response.Body)
	nilErr(err)
	fmt.Println("The response from web is", string(res))
	fmt.Printf("The response type is %T \n", res)
}
func nilErr(err error) {
	if err != nil {
		panic(err)
	}
}
