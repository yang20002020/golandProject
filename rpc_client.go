package main

import (
	"fmt"
	"net/rpc"
)

func main(){
	//1.用rpc链接服务器
   conn,err:= rpc.Dial("tcp","127.0.0.1:8080")
   if err!=nil{
   	fmt.Println("Dial err:",err)
   }
   defer conn.Close()
	//2.调用远程函数
	var reply string //接受返回值
	//hello是server 服务端 服务名，helloworld是server服务端中被调用的方法
	err=conn.Call("hello.HelloWorld","李白",&reply)
   if err!=nil{
   	fmt.Println("Call:",err)
   	return
   }
    fmt.Println(reply)
}