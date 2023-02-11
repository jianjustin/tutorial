package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	proto "go.guide/go-micro/helloworld/proto"
)

func main() {
	//创建一个新的服务 命名
	service := micro.NewService(micro.Name("greeter.client"))
	//服务初始化
	service.Init()

	//创建服务 绑定客户端 这个方法是在proto生成的文件中定义的
	greeter := proto.NewGreeterService("greeter", service.Client())

	//调用Hello方法 Hello方法同样是在proto生成的文件中定义的
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "World"})
	if err != nil {
		fmt.Println(err)
	}

	//打印结果
	fmt.Println(rsp.Greeting)
}
