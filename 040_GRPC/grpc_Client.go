package main

import (
	"040_GRPC/pb"
	"context"
	"fmt"
	"google.golang.org/grpc" //在src文件夹里存放google.golang.org/grpc文件
)

func main() {
	//1链接grpc 服务
	grpcConn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err:", err)
		return

	}
	defer grpcConn.Close()
	//2初始化grpc 客户端

	grpcClient := pb.NewSayClient(grpcConn)
	//创建并初始化techer 对象
	var teacher pb.Teacher
	teacher.Name = "itcast"
	teacher.Age = 18
	//3 调用远程服务
	t, err := grpcClient.SayHello(context.TODO(), &teacher)
	fmt.Println(t, err)

}
