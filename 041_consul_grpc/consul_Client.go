package main

import (
	"041_consul_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	////////////////// 以下是grpc 服务远程调用//////////////////
	var person pb.Person
	person.Name = "Andy"
	person.Age = 18
	//1.链接  服务
	grpcConn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//2. 初始化 grpc 客户端
	grpcClient := pb.NewHelloClient(grpcConn)
	//3.调用远程函数
	p, err := grpcClient.SayHello(context.TODO(), &person)
	fmt.Println(p, err)
}
