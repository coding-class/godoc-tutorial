package main

import (
	"encoding/json"
	"fmt"
)

// Rectangle a nice struct
type Rectangle struct {
	H int `json:"height"`
	W int `json:"width"`
}

// Area buat dapetin luas dari Rectangle
func (s Rectangle) Area() int {
	return s.H * s.W
}

func main() {
	var MyRect Rectangle

	MyRect.H = 2
	MyRect.W = 3

	resp, err := json.Marshal(MyRect)
	if err != nil {
		panic(err)
	}
	// json.Unmarshal() // String to struct

	fmt.Printf("%+v \n", string(resp))
}
