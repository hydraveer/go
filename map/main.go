package main

import "fmt"

func main() {
	lang := make(map[string]string)
	lang["1"] = "JS"
	lang["2"] = "Ruby"
	lang["3"] = "Ruby"
	lang["4"] = "Ruby"
	fmt.Println(lang)
	fmt.Println(lang["1"])
	for i , val := range lang{
		fmt.Printf("The value for key %v is %v \n", i,val)
	}
}
