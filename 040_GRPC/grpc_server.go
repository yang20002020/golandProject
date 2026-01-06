package main

import (
	"040_GRPC/pb"
	"context" // 上下文  ---goroutine （go程） 之间用来进行数据传递API包
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// 定义类
type Children struct {
}

// 按接口绑定类方法
func (this *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += "is Sleeping"
	return t, nil
}

func main() {
	//1.初始化一个grpc对象
	grpc.NewServer()
	grpcServer := grpc.NewServer()
	//2. 注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))
	//3.设置监听 指定ip port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()
	//4. 启动服务 server（）
	grpcServer.Server(listener)
}
