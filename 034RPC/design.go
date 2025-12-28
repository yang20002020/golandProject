package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MyInterface interface {
	//客户端需要调用服务端的函数 （绑定函数）
	HelloWorld(string, *string) error
}

/**************服务端****************************************/
//服务端在注册rpc对象时，能让编译期检测出，注册对象是否合法
//创建接口，在接口中定义方法原型

// 服务端调用该方法时，需要给i传参；实参数应该是 实现了helloWorld 方法的类对象！
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

/***************客户端********************************************/
//err=conn.Call("hello.HelloWorld","李白",&reply)
//目的，在编译阶段，就防止字符串"hello.HelloWorld"写错
//向调用本地函数一样，调用远程函数
//定义类
type Myclient struct {
	//conn类型为*rpc.Client
	c *rpc.Client
}

/********Client*******/
//由于使用了C调用Call,因此需要初始化C
//1.链接服务器
func initClient(addr string) Myclient {
	//func Dial(network, address string) (*rpc.Client, error)
	//conn类型为*rpc.Client
	// conn,err:= rpc.Dial("tcp","127.0.0.1:8080")
	conn, _ := jsonrpc.Dial("tcp", addr)
	return Myclient{c: conn}
}

/***************************/
//Myclient绑定HelloWorld, 实现接口MyInterface的函数声明

//client 是一个结构体 包含函数call
// func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
// 对比代码差别err=conn.Call("hello.HelloWorld","李白",&reply)
// 2.客户端远程调用

// err=conn.Call("hello.HelloWorld","李白",&reply)
// 目的，在编译阶段，就防止字符串"hello.HelloWorld"写错
func (this *Myclient) HelloWorld(a string, b *string) error {
	//参数1，参照上面的interface，registerName而来！a，传入参数；b，传出参数
	//call 方法返回值也是error
	return this.c.Call("hello.HelloWorld", a, b)
}
