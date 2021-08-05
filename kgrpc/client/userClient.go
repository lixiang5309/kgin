package client

import (
	"context"
	"fmt"

	"kgin/kgrpc/pb"

	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8098", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()
	// 新建一个客户端
	c := kgrpc.NewGreeterClient(conn)
	// 调用服务端函数
	r, err := c.GetAllUsers(context.Background(), &kgrpc.All{})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Println(r.List)
}
