package service

import (
	nethttp "net/http"
	"log"
)

// 服务运行
func Run() {
	httpHandler := createService()
	log.Println("demo1服务启动，服务地址：127.0.0.1:8088")
	err := nethttp.ListenAndServe(":8088", httpHandler)

	if nil != err {
		log.Println(err)
	}
}
