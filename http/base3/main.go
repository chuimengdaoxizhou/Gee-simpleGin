package base3

import (
	"Gee/http/base3/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("indexHandleFunc")
	})

	r.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HelloWorldHandleFunc")
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}
