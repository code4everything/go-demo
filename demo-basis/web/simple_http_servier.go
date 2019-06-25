package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe("localhost:5656", nil); err != nil {
		fmt.Println("listen error:", err.Error())
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside server handler")
	_, _ = fmt.Fprintf(w, "hello,"+req.URL.Path[1:])
}
