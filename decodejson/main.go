package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	jsonDataform := []byte(`
		{
                "coursename": "Java",
                "Price": 299,
                "website": "Coursera",
                "tags": [
                        "java",
                        "springboot"
                ]
        }
	`)
	var lco course
	checkValid := json.Valid(jsonDataform)
	if checkValid {
		fmt.Println("Valid json")
		json.Unmarshal(jsonDataform, &lco)
		fmt.Printf("%#v\n", lco)
	}else{
		fmt.Println("json was not vald")
	}
}
