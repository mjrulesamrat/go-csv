package main

import (
	"encoding/json"
	"fmt"
)
type Response1 struct {
	Page int
	Fruits []string
}

// using tags at struct field names to customize the encoded JSON key names
type Response2 struct {
	Page int `json:"Page_No"`
	Fruits []string `json:"fruits_name"`
}

func main() {

	res1 := &Response1{
		Page: 1,
		Fruits : []string{"Banana", "Apple", "lichi"},
	}
	res1B, _ := json.Marshal(res1)
	fmt.Println(string(res1B))

	res2 := &Response2{
		Page: 1,
		Fruits : []string{"Banana", "Apple", "lichi"},
	}
	res2B, _ := json.Marshal(res2)
	fmt.Println(string(res2B))	
}