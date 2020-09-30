package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"micro/protos"
)

func main() {
	//注册到consul
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	//创建服务
	service := micro.NewService(
		micro.Name("greeter.service"),
		micro.Registry(reg),
	)
	//初始化
	service.Init()

	//注册服务到handler
	err := protos.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		fmt.Println(err)
	}

	//运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
