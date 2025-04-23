package base1

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexHandleFunc")
}

func HelloWorldHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloWorldHandleFunc")
}

func main() {
	http.HandleFunc("/index", IndexHandleFunc)
	http.HandleFunc("/hello", HelloWorldHandleFunc)
	log.Fatalln("ListenAndServe: ", http.ListenAndServe(":8080", nil))
}
