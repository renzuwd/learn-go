package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	d := struct {
		Id   int64
		Name string
	}{
		Id:   100,
		Name: "xw",
	}
	jsonstr, _ := json.Marshal(d)
	fmt.Println(string(jsonstr))

	// var data struct {
	// 	Id   int64
	// 	Name string
	// }

	var data interface{} // map[Id:100 Name:xw]
	json.Unmarshal(jsonstr, &data)
	fmt.Println(data)
}
