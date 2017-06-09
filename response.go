package main

import (
	"encoding/json"
	"fmt"
)

func response(data interface{}) string {
	fmt.Println(data)
	j, _ := json.Marshal(data)
	fmt.Println(j)
	return string(j)
	// return "test"
}
