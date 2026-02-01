package main

import (
	"go-micro.dev/v4"
	handler "golandProject/bj38/handler"
)

func main() {
	// 创建服务（v4中推荐使用NewService）
	service := micro.NewService(
		micro.Name("Say"),
		micro.Address(":8080"),
	)

	// 注意：这里不需要再重复调用 micro.New("helloworld")
	// 因为上面已经通过 NewService 创建了服务

	// 初始化服务
	service.Init()

	// 注册处理程序（v4的API）
	micro.RegisterHandler(service.Server(), new(handler.SayHandler))
	// 或者等价写法：service.Server().Handle(service.Server().NewHandler(new(Say)))

	// 运行服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
