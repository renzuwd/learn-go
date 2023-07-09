package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("go http server"))
	})
	http.ListenAndServe(":8999", nil)
}
