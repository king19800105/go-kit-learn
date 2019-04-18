package main

import (
	"google.golang.org/grpc"
	"log"
	"github.com/king19800105/go-kit-learn/demo5/pkg/grpc/pb"
	"context"
	"time"
)

// grpc客户端(微服务之间的调用，应当封装到服务对象的成员属性上)
func main() {
	// 创建链接对象
	conn, err := grpc.Dial("127.0.0.1:9092", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
		return
	}
	// 调用
	client := pb.NewOrderClient(conn)
	// 设置一秒超时
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 1)
	reply, err := client.Create(ctx, &pb.CreateRequest{
		OrderId: "111",
	})

	if nil == reply || nil != err {
		log.Println("无效失败，原因：", err.Error())
		return
	}

	log.Println(reply)
	return

}
