package main

import (
	"encoding/json"
)

func response(data interface{}) string {
	j, _ := json.Marshal(data)
	return string(j)
}
