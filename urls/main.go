package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://google.com/?course=react&paymentId=1"

func main() {
	fmt.Println("Handling the urls")
	res,_ := url.Parse(myurl)
	fmt.Println(res.Scheme)
	fmt.Println(res.Path)
	qparam := res.Query()
	fmt.Printf("The type of this qparam is %T \n", qparam)
	fmt.Println(qparam)
	for i,val := range qparam{
		fmt.Printf("the value of param is %v this is key and the value is %v \n", i,val)
	}
}
