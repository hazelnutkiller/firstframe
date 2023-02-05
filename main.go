//腳本中創立main.go 啟動腳本
package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()

	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello World!")
	})

	s.Run() // 默认 80 端口
}
