package gee

import (
	"net/http"
)

// router 结构体保存了路由表 handlers，
// key 是请求方法和路径拼接而成，value 是对应的处理函数
type router struct {
	handlers map[string]HandlerFunc
}

// 创建一个新的 router 实例，并初始化 handlers 映射表
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute 向路由表中添加一条新的路由记录
// method 表示请求方法（如 GET、POST）
// pattern 表示请求路径
// handler 是处理该路由的函数
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // 构造唯一的键
	r.handlers[key] = handler     // 注册处理函数
}

// handle 根据请求的路径和方法查找并执行对应的处理函数
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c) // 找到对应处理函数，执行处理逻辑
	} else {
		// 没有匹配的路由，返回 404 错误
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
