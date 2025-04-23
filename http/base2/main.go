package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/index":
		fmt.Println("indexHandleFunc")
	case "/hello":
		fmt.Println("HelloWorldHandleFunc")
	default:
		fmt.Println("404 Not Found")
	}
}

func main() {
	engine := new(Engine)
	err := http.ListenAndServe(":8080", engine)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}
