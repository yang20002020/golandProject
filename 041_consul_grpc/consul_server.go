package main

import (
	"041_consul_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	// go  get github.com/hashicorp/consul/api 下载
	"github.com/hashicorp/consul/api"
)

//定义类
type Children struct {
}

//绑定类方法，实现接口
func (this *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "hello" + p.Name
	return p, nil
}

/*
把grpc服务端注册到consul

*/
func main() {
	//把grpc 服务，注册到consul上
	//1.初始化cosul配置
	consulConfig := api.DefaultConfig()
	//2.创建 consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api,NewClient err:", err)
		return
	}
	//3.告诉consul，即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		ID:      "bj38",
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc And  Consul",
		Address: "127.0.0.1",
		Port:    8800,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8800",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//4.注册 grpc服务到 consul上 //consulClient是consul对象
	consulClient.Agent().ServiceRegister(&reg)

	//////////// 以下是grpc 服务远程调用//////////////////
	// 1.初始化grpc 对象
	grpcServer := grpc.NewServer()
	//2.注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))
	//3. 设置监听，指定ip port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("***********服务启动")

	//4. 启动服务
	grpcServer.Serve(listener)
}
