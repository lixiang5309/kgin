package server

import (
	"fmt"
	"net"

	"kgin/kapi/service"
	"kgin/kgrpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":8098")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}
	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	ss := new(service.UserService)
	kgrpc.RegisterGreeterServer(s, ss)
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}
