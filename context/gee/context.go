package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 是一个便捷的 map 类型，用于构造 JSON 数据
type H map[string]interface{}

// Context 封装了 HTTP 请求和响应的上下文信息
type Context struct {
	// 原始的 http 请求和响应对象
	Writer http.ResponseWriter
	Req    *http.Request

	// 请求相关的信息
	Path   string // 请求路径
	Method string // 请求方法（GET、POST 等）

	// 响应相关的信息
	StatusCode int // HTTP 响应状态码
}

// 创建一个新的 Context 实例，封装请求和响应
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 获取 POST 表单参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 获取 URL 查询参数 (?key=value)
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置 HTTP 响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置 HTTP 响应头部
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 返回纯文本响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 返回 JSON 格式的数据
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 返回原始字节数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 返回 HTML 格式的响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
