package main

import (
	"encoding/base64"
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
	djson, _ := json.Marshal(d)
	base64Str := base64.RawURLEncoding.EncodeToString([]byte(djson))
	fmt.Println(base64Str)

	s, _ := base64.RawURLEncoding.DecodeString(base64Str)
	fmt.Println(string(s))
}
