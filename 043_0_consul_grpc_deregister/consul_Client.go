package main

import (
	"043_0_consul_grpc_deregister/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func main() {
	//初始化cosul配置
	consulConfig := api.DefaultConfig()
	//创建consul对象 ---(可以重新指定consul 属性 ：ip/port,也可以使用默认)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("err:", err)
	}

	//服务发现，从consul上，获取健康的服务
	// 服务名：grpc And  Consul"
	//tag:别名 外号  如果有多个，任选一个
	//passonly：是否通过健康检查
	//q: 查询的额外参数
	// 返回值 ([]*ServiceEntry, *QueryMeta, error)
	//返回值：1.[]*ServiceEntry 储存服务的切片；服务的集群名字是一样的，但是每一个服务的ip和port 是不一样的
	//返回值 ：2。QueryMeta ，额外查询返回的值，一般为nil
	services, _, err := consulClient.Health().Service("grpc And  Consul", "grpc", true, nil)
	if err != nil {
		fmt.Println("************err:", err)
		return
	}
	if 0 == len(services) {
		fmt.Println("len(services)==0")
		return
	}
	fmt.Println("services[0].Service.Address")
	fmt.Println(services[0].Service.Address)
	fmt.Println("services[0].Service.Port")
	fmt.Println(services[0].Service.Port)
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	fmt.Println("addr")
	fmt.Println(addr)
	////////////////// 以下是grpc 服务远程调用//////////////////
	var person pb.Person
	person.Name = "Andy"
	person.Age = 18
	//1.链接  服务
	//grpcConn,_:=grpc.Dial("127.0.0.1:8800",grpc.WithTransportCredentials(insecure.NewCredentials()))
	//使用服务发现 consul 上的ip/port 来与服务建立链接
	grpcConn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	//2. 初始化 grpc 客户端
	grpcClient := pb.NewHelloClient(grpcConn)
	//3.调用远程函数
	p, err := grpcClient.SayHello(context.TODO(), &person)
	fmt.Println(p, err)
}
