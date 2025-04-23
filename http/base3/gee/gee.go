package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

func (engine *Engine) addRouter(method string, pattern string, fun HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = fun
}

func (engine *Engine) Get(pattern string, fun HandleFunc) {
	engine.addRouter("Get", pattern, fun)
}

func (engine *Engine) Post(pattern string, fun HandleFunc) {
	engine.addRouter("Post", pattern, fun)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
