package main

import (
	"fmt"
	//"net/rpc/jsonrpc"
)

// 结合design.go测试
func main() {
	//func initClient(addr string)Myclient,返回的是客户端结构体Myclient
	//1.链接服务器  //myClient 是类 Myclient的实例
	myClient := initClient("127.0.0.1:8080") //等价于rpc.Dial("tcp","127.0.0.1:8080")
	var resp string
	// 2.客户端远程调用
	//err=conn.Call("hello.HelloWorld","李白",&reply)
	//目的，在编译阶段，就防止字符串"hello.HelloWorld"写错。该部分封装在design中，client中不需要再写
	err := myClient.HelloWorld("杜甫", &resp) //HelloWorld 是服务端绑定的函数
	fmt.Println("########11111#########")
	if err != nil {
		fmt.Println("HelloWorld err:", err)
		return
	}
	fmt.Println("##########222222222222#######")
	fmt.Println(resp, err)
	fmt.Println("##########333333333333333#######")
}
