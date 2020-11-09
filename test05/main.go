package main

import (
	"fmt"
)

func main() {
	test_map := make(map[string][]string)
	test_map["test"] = make([]string, 0)
	test_map["test"] = append(test_map["test"], "Hello World")

	// verify a certain key exists
	if _, ok := test_map["map"]; !ok {
		test_map["map"] = []string {"Hello Map"}
	}

	// iterate through all keys/values
	for k, v := range test_map {
		fmt.Printf("%v, %v \n", k, v)
	}
}