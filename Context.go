package peiqi

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{Writer: w, Request: r}
}

// 类似 gin 的 ctx.JSON() 方法
func (ctx *Context) Send(statuscode int, data interface{}) {
	// 将响应头设置成application/json
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(statuscode)
	// 将数据序列化成 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(ctx.Writer, "序列化JSON数据失败", http.StatusInternalServerError)
		return
	}
	ctx.Writer.Write(jsonData)
}
