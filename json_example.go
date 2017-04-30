package main

import (
	"encoding/json"
	"fmt"
	_ "os"
)

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("Jay Modi")
	fmt.Println(string(strB))

	slcD := []string{"jay", "h", "modi"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "peach": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
}
