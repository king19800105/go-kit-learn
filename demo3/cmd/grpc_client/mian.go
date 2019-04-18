package main

import (
	"google.golang.org/grpc"
	"log"
	"github.com/king19800105/go-kit-learn/demo3/pkg/grpc/pb"
	"context"
)

// grpc客户端(微服务之间的调用，应当封装到服务对象的成员属性上)
func main()  {
	// 创建链接对象
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
		return
	}
	// 调用
	client := pb.NewOrderClient(conn)
	reply, err := client.Create(context.Background(), &pb.CreateRequest{
		OrderId: "demo3no",
	})

	if nil == reply || nil != err {
		log.Println("无效失败，原因：", err.Error())
		return
	}

	log.Println(reply)
	return

}
