package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"micro/protos"
)

var consulReg registry.Registry

func main() {

	//注册到consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"192.168.5.196:8500",
		}
	})

	//获取服务列表
	ser, _ := reg.GetService("greeter.service")
	for _, value := range ser {
		for _, item := range value.Nodes {
			fmt.Println(value.Name, item.Id, item.Address)
		}
	}

	//创建服务
	Client := micro.NewService(
		micro.Name("greeter.client"),
		micro.Registry(reg),
	)

	Client.Init()

	//调用
	greeter := protos.NewGreeterService("greeter.service", Client.Client())
	rsp, err := greeter.Hello(context.TODO(), &protos.Request{Name: "阿西吧"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Greeting)

	////运行服务
	//if err := service.Run(); err != nil {
	//	fmt.Println(err)
	//}
}
