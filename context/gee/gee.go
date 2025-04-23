package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义请求处理函数的类型，接收一个 Context 指针作为参数
type HandlerFunc func(*Context)

// Engine 实现了 ServeHTTP 接口，是框架的核心结构体
type Engine struct {
	router *router // 路由器，用于注册和查找路由
}

// New 是 Engine 的构造函数，初始化一个新的 Engine 实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 用于注册一个路由规则
// method 是请求方法（GET、POST）
// pattern 是路由路径
// handler 是处理函数
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern) // 打印路由注册日志
	engine.router.addRoute(method, pattern, handler)
}

// GET 是注册 GET 请求路由的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 是注册 POST 请求路由的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动 HTTP 服务，监听指定地址
func (engine *Engine) Run(addr string) (err error) {
	// 将 engine 作为 Handler 传入 http.ListenAndServe，会自动调用 ServeHTTP 方法
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 是 Engine 实现 http.Handler 接口的核心方法
// 每次收到请求都会调用该方法进行路由分发
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 创建 Context 封装请求和响应信息
	c := newContext(w, req)
	// 由 router 根据请求路径和方法调用对应的处理函数
	engine.router.handle(c)
}
