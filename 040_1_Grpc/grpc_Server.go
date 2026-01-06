package main

import (
	"040_1_Grpc/pb"
	"context" //上下文  gorutine (go程) 之间用来数据传递，API包
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// 定义类
type Children struct {
}

/*
grpc_server.go 和grpc_client.go分别是独立的进程，两个进程 之间的桥梁是person.pb.go
*/

/*
//定义 服务
service SayName {
 rpc SayHello(Teacher) returns(Teacher);
*/
//按接口绑定类方法
func (this *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += "is Sleeping"
	return t, nil
}
func main() {
	//1.初始化一个grpc对象
	grpcServer := grpc.NewServer()
	//2. 注册服务 SayName 是proto文件中的服务函数SayName（）
	pb.RegisterSayNameServer(grpcServer, new(Children))
	//3.设置监听 指定ip port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()
	//4. 启动服务 server（）
	grpcServer.Serve(listener)

}
