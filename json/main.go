package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	coursera := []course{
		{
			"React Js", 299, "Coursera", "12345", []string{"js", "node"},
		},
		{
			"Java", 299, "Coursera", "1234", []string{"java", "springboot"},
		},
		{
			"Python", 299, "Coursera", "123", nil,
		},
	}
	finalJson, err := json.MarshalIndent(coursera,"","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The valu of json %s \n ", finalJson)
}
