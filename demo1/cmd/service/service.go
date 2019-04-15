package service

import (
	nethttp "net/http"
	"log"
)

// 服务运行
func Run() {
	httpHandler := createService()
	err := nethttp.ListenAndServe(":8088", httpHandler)

	if nil != err {
		log.Println(err)
	}
}
